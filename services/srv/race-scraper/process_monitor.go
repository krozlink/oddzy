package main

import (
	"context"
	"fmt"
	racing "github.com/krozlink/oddzy/services/srv/racing/proto"
	_ "github.com/micro/go-plugins/registry/consul"
	"sort"
	"strconv"
	"time"
)

type byNextScrapeTime []*scheduledScrape

type scheduledScrape struct {
	race *racing.Race
	next time.Time
}

func monitorOpenRaces(p *scrapeProcess, open []*racing.Race) (chan<- bool, <-chan bool) {
	log := logWithContext(p.ctx, "monitorOpenRaces")
	p.status = "RACE_MONITORING"

	log.Info("Starting race monitoring")

	overdue, upcoming := categorise(open)

	sort.Sort(byNextScrapeTime(overdue))  // races that are past their scheduled start time
	sort.Sort(byNextScrapeTime(upcoming)) // races before their scheduled start time

	log.Infof("Overdue races: %v", len(overdue))
	log.Infof("Upcoming races: %v", len(upcoming))

	scrapeQueue := make(chan *scheduledScrape)
	updateQueue := make(chan *scheduledScrape, 10)

	stopScrape := make(chan bool)
	stopUpdate := make(chan bool)

	go func() {
		for { // loop through all open races
			race := popNext(&overdue, &upcoming) // get the next race to scrape (will remove it from the list)
		processmessages:
			for {
				if race == nil { // no race found so wait for a previous race to be updated or wait for a while and check again for a new race
					log.Debugf("No races found - waiting")
					select {
					case u := <-updateQueue: // a race has been scraped and updated - put it back on either upcoming or overdue depending on start time
						requeueRace(p.ctx, u, &overdue, &upcoming)
						break processmessages
					case <-time.After(time.Minute * 5):
						break processmessages
					case <-stopScrape:
						return
					}
				} else { // race found so wait until time to scrape and then queue it
					wait := time.Until(race.next).Seconds()
					if wait > 0 {
						log.Debugf("Race %v (%v) found. Last scrape: %v  Next scrape: %v  - waiting %.2f seconds", race.race.RaceId, race.race.Name, formatTime(time.Unix(race.race.LastUpdated, 0)), formatTime(race.next), wait)
					} else {
						wait = 0
						log.Debugf("Race %v (%v) found. Last scrape: %v  Next scrape: %v  - scraping immediately", race.race.RaceId, race.race.Name, formatTime(time.Unix(race.race.LastUpdated, 0)), formatTime(race.next))
					}

					select {
					// a race has been scraped and updated - put it back on either upcoming or overdue depending on start time
					// once added re-retrieve a race off the queue as the newly updated race may still be the next one to be scraped again
					case u := <-updateQueue:
						requeueRace(p.ctx, u, &overdue, &upcoming)
						requeueRace(p.ctx, race, &overdue, &upcoming)
						break processmessages
					case <-time.After(time.Second * time.Duration(wait)): // wait until it is time then queue it for scraping
						log.Debugf("Race %v about to be scraped", race.race.RaceId)
						scrapeQueue <- race
						break processmessages
					case <-stopScrape:
						return
					}
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case r := <-scrapeQueue:
				log.Debugf("Scraping race %v", r.race.RaceId)
				mDate := time.Unix(r.race.MeetingStart, 0).Format("2006-01-02")
				m, ok := p.meetingsByID[r.race.MeetingId]
				if !ok {
					log.Fatalf("race %v (%v) does not have an existing meeting (%v)", r.race.RaceId, r.race.Name, r.race.MeetingId)
				}
				cal, err := p.scraper.ScrapeRaceSchedule(p.ctx, m.RaceType, mDate)
				if err != nil {
					log.Errorf("Unable to scrape schedule for event type '%v' on %v' - %v\nSkipping race", m.RaceType, mDate, err)
					continue
				}
				updated, err := getRaceFromSchedule(cal, r.race)
				if err != nil {
					log.Errorf("Unable to scrape race %v (source id: %v) as is it could not be found in the race schedule\nSkipping race", r.race.RaceId, r.race.SourceId)
					continue
				}
				updated.LastUpdated = time.Now().Unix()
				p.racesByID[r.race.RaceId] = updated
				p.racesBySource[r.race.SourceId] = updated

				req := &racing.UpdateRaceRequest{
					ActualStart:    updated.ActualStart,
					RaceId:         updated.RaceId,
					Results:        updated.Results,
					ScheduledStart: updated.ScheduledStart,
					Status:         updated.Status,
				}
				_, err = p.racing.UpdateRace(p.ctx, req)
				if err != nil {
					log.Errorf("Unable to update race id '%v' - %v", r.race.RaceId, err)
					log.Errorf("Skipping race %v", r.race.RaceId)
					continue
				}

				// if status is still open put back on either overdue or upcoming depending on start time
				if updated.Status == "OPEN" || updated.Status == "INTERIM" {
					log.Debugf("Race %v has has been scraped but has not ended (status %v) - pushing back on the update queue", r.race.RaceId, updated.Status)
					s := &scheduledScrape{
						race: updated,
						next: nextScrapeTime(updated),
					}
					updateQueue <- s
				} else {
					log.Debugf("Race %v has has been scraped and is now %v - removing from update queue", r.race.RaceId, updated.Status)
				}
			case <-stopUpdate:
				return
			}
		}
	}()

	stop := make(chan bool) // stop the monitoring by sending this channel a message
	done := make(chan bool) // notify when monitoring completed or stopped

	go func() {
		<-stop
		log.Info("Race monitoring has been stopped")
		stopScrape <- true
		stopUpdate <- true
		done <- true
	}()

	return stop, done
}

// categorise takes a list of open races and splits them into collections of upcoming and overdue races
// A race is overdue if it is past its scheduled start time
func categorise(open []*racing.Race) ([]*scheduledScrape, []*scheduledScrape) {

	overdue := make([]*scheduledScrape, 0)  // open races past their start time that are unresulted
	upcoming := make([]*scheduledScrape, 0) // open races that are not past their start time

	now := time.Now()
	for _, r := range open {
		if now.Unix() > r.ScheduledStart { // race is overdue
			overdue = append(overdue, &scheduledScrape{
				race: r,
				next: nextScrapeTime(r),
			})
		} else { // race is upcoming
			upcoming = append(upcoming, &scheduledScrape{
				race: r,
				next: nextScrapeTime(r),
			})
		}
	}

	return overdue, upcoming
}

func nextScrapeTime(r *racing.Race) time.Time {
	scheduled := time.Unix(r.ScheduledStart, 0)
	now := time.Now()
	lastUpdate := time.Unix(r.LastUpdated, 0)

	var next time.Time
	if scheduled.Before(now) { // if overdue race
		next = lastUpdate.Add(time.Second * 60)
		return max(next, now) // scrape max of (last update + overdue interval) and now
	}

	// upcoming races are scraped:
	//		every 6 hours > 12 hours before
	//		every 1 hour <= 12 hours before
	//		every 10 minutes < 1 hour before (or on scheduled start time, whichever comes first)
	if time.Until(scheduled).Hours() > 12 {
		next = lastUpdate.Add(time.Hour * 6)
	} else if time.Until(scheduled).Hours() > 1 {
		next = lastUpdate.Add(time.Hour * 1)
	} else {
		next = lastUpdate.Add(time.Minute * 10)
		if scheduled.Before(next) {
			next = scheduled
		}
	}

	return max(next, now)
}

func max(t1, t2 time.Time) time.Time {
	if t1.After(t2) {
		return t1
	}
	return t2
}

func requeueRace(ctx context.Context, r *scheduledScrape, overdue, upcoming *[]*scheduledScrape) {
	log := logWithContext(ctx, "requeueRace")

	if time.Unix(r.race.ScheduledStart, 0).After(time.Now()) {
		log.Debugf("Race %v added back on upcoming queue for scraping on %v", r.race.RaceId, formatTime(r.next))
		*upcoming = append(*upcoming, r)
		sort.Sort(byNextScrapeTime(*upcoming))
	} else {
		log.Debugf("Race %v added back on overdue queue for scraping on %v", r.race.RaceId, formatTime(r.next))
		*overdue = append(*overdue, r)
		sort.Sort(byNextScrapeTime(*overdue))
	}
}

func popNext(overdue, upcoming *[]*scheduledScrape) *scheduledScrape {
	if len(*overdue) == 0 && len(*upcoming) == 0 {
		return nil
	} else if len(*overdue) == 0 {
		first := (*upcoming)[0]
		*upcoming = (*upcoming)[1:]
		return first
	} else if len(*upcoming) == 0 {
		first := (*overdue)[0]
		*overdue = (*overdue)[1:]
		return first
	}

	if (*upcoming)[0].next.Before((*overdue)[0].next) {
		first := (*upcoming)[0]
		*upcoming = (*upcoming)[1:]
		return first
	}
	first := (*overdue)[0]
	*overdue = (*overdue)[1:]
	return first
}

func pushRace(overdue, upcoming []*scheduledScrape, r *scheduledScrape) ([]*scheduledScrape, []*scheduledScrape) {
	start := time.Unix(r.race.ScheduledStart, 0)
	if start.After(time.Now()) {
		upcoming = append(upcoming, r)
		sort.Sort(byNextScrapeTime(upcoming))
	} else {
		overdue = append(overdue, r)
		sort.Sort(byNextScrapeTime(overdue))
	}

	return overdue, upcoming
}

func getRaceFromSchedule(cal *RaceSchedule, original *racing.Race) (*racing.Race, error) {
	var race *racing.Race

	for _, rg := range cal.RegionGroups {
		for _, m := range rg.Meetings {
			for _, e := range m.Events {
				if original.SourceId == strconv.Itoa(int(e.EventID)) {
					race = &racing.Race{
						MeetingId:      original.MeetingId,
						RaceId:         original.RaceId,
						MeetingStart:   original.MeetingStart,
						Name:           e.EventName,
						Number:         e.EventNumber,
						Results:        e.Results,
						ScheduledStart: e.StartTime,
						Status:         getRaceStatusFromSchedule(e.IsAbandoned, e.Resulted, e.Results),
						SourceId:       original.SourceId,
					}
					return race, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("No race found in the schedule with source id %v", original.SourceId)
}

func (s byNextScrapeTime) Len() int {
	return len(s)
}

func (s byNextScrapeTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byNextScrapeTime) Less(i, j int) bool {
	return s[i].next.Before(s[j].next)
}

func formatTime(t time.Time) string {
	return t.Local().Format("02-01-2006 15:04:05")
}
