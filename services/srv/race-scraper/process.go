package main

import (
	"context"
	// proto "github.com/krozlink/oddzy/services/srv/race-scraper/proto"
	"fmt"
	racing "github.com/krozlink/oddzy/services/srv/racing/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/consul"
	"github.com/satori/go.uuid"
	"strconv"
	"strings"
	"sync"
	"time"
)

type scrapeProcess struct {
	status    string
	done      chan bool
	http      handler
	racing    racing.RacingService
	scraper   Scraper
	startDate time.Time
	endDate   time.Time

	unscrapedRaces []*racing.Race

	meetingsByID     map[string]*racing.Meeting
	meetingsBySource map[string]*racing.Meeting
	racesByID        map[string]*racing.Race
	racesBySource    map[string]*racing.Race
}

type externalRaceData struct {
	existingMeetings []*racing.Meeting
	existingRaces    []*racing.Race
	newMeetings      []*racing.Meeting
	newRaces         []*racing.Race
}

var raceTypes = []string{"horse-racing", "harness", "greyhounds"}

func (p *scrapeProcess) run() {

	setup(p)

	scrapeUnscrapedRaces(p)

	// STATUS - RACE MONITORING
	// Process 1
	//		Monitor time until race and race status and push required races to queue if they are not already on there
	//		Only needs to run at most every 30 seconds. It can determine its own sleep duration depending on upcoming races
	// Process 2 - Periodically take items off the queue and scrape them. Ensure minimum interval to avoid overloading server
	monitorUpcomingRaces(p)
}

func scrapeUnscrapedRaces(p *scrapeProcess) {
	// STATUS - RACE_CREATION
	p.status = "RACE_CREATION"

	scrape := make(chan *racing.Race, 1000)
	go func() {
		for _, r := range p.unscrapedRaces {
			scrape <- r
		}
		close(scrape)
	}()

	update := make(chan *racing.UpdateRaceRequest, 1000)
	first := true
	for r := range scrape {
		if !first {
			<-time.After(1 * time.Second)
		}

		card, err := p.scraper.ScrapeRaceCard(r.SourceId)
		// TODO error handling
		selections, err := parseRaceCard(card)
		// TODO error handling

		req := &racing.UpdateRaceRequest{
			Race:       r,
			Selections: selections,
		}
		update <- req

		first = false
	}
	close(update)

	ctx := context.Background()
	for r := range update {
		_, err := p.racing.UpdateRace(ctx, r)
		// TODO error handling
	}

	// go routine to take items off the list in a loop
	// (wait 1 second if not the first)
	// scrape race
	// push to update channel

	// second goroutine to read update channel
	//call UpdateRace
	// return when channel closed

	// For each unscraped race
	//	Scrape the race
	// 	Update the race - UpdateRace
}

func monitorUpcomingRaces(p *scrapeProcess) {

}

func getDateRange() (time.Time, time.Time) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Add(time.Hour * -24)
	end := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Add(time.Hour * 48)
	return start, end
}

func newScrapeProcess() scrapeProcess {
	service := micro.NewService(micro.Name("racing.client"))
	service.Init()
	client := racing.NewRacingService("racing", service.Client())

	return scrapeProcess{
		status:  "INITIALISING",
		done:    make(chan bool),
		http:    handler{},
		racing:  client,
		scraper: &OddscomauScraper{},
	}
}

func setup(p *scrapeProcess) {
	// STATUS - SETUP
	p.status = "SETUP"

	// read all internal meeting data for scraping period (yesterday to 2 days from now)
	p.startDate, p.endDate = getDateRange()
	meetings, races, err := readInternal(p)
	if err != nil {
		fmt.Printf("An error occurred reading internal racing data - %v", err)
	}

	for _, m := range meetings {
		p.meetingsByID[m.MeetingId] = m
		p.meetingsBySource[m.SourceId] = m
	}

	for _, r := range races {
		p.racesByID[r.RaceId] = r
		p.racesBySource[r.SourceId] = r
	}

	// read all external meeting data for scraping period (yesterday to 2 days from now)
	data, err := readExternal(p)
	if err != nil {
		fmt.Printf("An error occurred reading external racing data - %v", err)
	}

	// update all existing meetings that have changed
	uMeetings := make([]*racing.Meeting, 1)
	for _, m := range data.existingMeetings {
		current := p.meetingsByID[m.MeetingId]
		if meetingChanged(current, m) {
			uMeetings = append(uMeetings, m)
		}
	}
	updateExistingMeetings(p.racing, uMeetings)

	// update all existing races that have changed
	uRaces := make([]*racing.Race, 1)
	for _, r := range data.existingRaces {
		current := p.racesByID[r.RaceId]
		if raceChanged(current, r) {
			uRaces = append(uRaces, r)
		}

		// flag existing races that have not been individually scraped yet
		if r.LastUpdated == 0 {
			p.unscrapedRaces = append(p.unscrapedRaces, r)
		}
	}
	updateExistingRaces(p.racing, uRaces)

	// Create new meetings and races
	err = createNewMeetings(p.racing, data.newMeetings)
	err = createNewRaces(p.racing, data.newRaces)

	// flag new races that have not been individually scraped yet
	for _, r := range data.newRaces {
		p.unscrapedRaces = append(p.unscrapedRaces, r)
	}
}

func readExternal(p *scrapeProcess) (*externalRaceData, error) {
	data := &externalRaceData{}

	// read race calendars from scraping period
	for _, t := range raceTypes {
		for d := p.startDate; d.Before(p.endDate) || d.Equal(p.endDate); d = d.Add(time.Hour * 24) {
			if d != p.startDate {
				<-time.After(time.Second * 1)
			}
			cal, err := p.scraper.ScrapeRaceCalendar(t, d.Format("2006-01-02"))
			if err != nil {
				return nil, fmt.Errorf("Unable to read race calendar for %v on %v - %v", t, d, err)
			}

			d, err := processRaceCalendar(p, d, t, cal)
			if err != nil {
				return nil, fmt.Errorf("Unable to process race calendar for %v on %v - %v", t, d, err)
			}

			data.existingMeetings = append(data.existingMeetings, d.existingMeetings...)
			data.existingRaces = append(data.existingRaces, d.existingRaces...)
			data.newMeetings = append(data.newMeetings, d.newMeetings...)
			data.newRaces = append(data.newRaces, d.newRaces...)
		}
	}

	return data, nil
}

func readInternal(p *scrapeProcess) ([]*racing.Meeting, []*racing.Race, error) {
	// read all internal meeting data for scraping period (yesterday to 2 days from now)
	//		ListMeetingsByDate
	//		ListRacesByMeetingDate
	ctx := context.Background()
	mReq := &racing.ListMeetingsByDateRequest{
		StartDate: p.startDate.Unix(),
		EndDate:   p.endDate.Unix(),
	}
	mResp, err := p.racing.ListMeetingsByDate(ctx, mReq)
	if err != nil {
		return nil, nil, err
	}

	rReq := &racing.ListRacesByMeetingDateRequest{
		StartDate: p.startDate.Unix(),
		EndDate:   p.endDate.Unix(),
	}
	rResp, err := p.racing.ListRacesByMeetingDate(ctx, rReq)
	if err != nil {
		return nil, nil, err
	}

	return mResp.Meetings, rResp.Races, nil
}

func processRaceCalendar(p *scrapeProcess, date time.Time, eventType string, c *RaceCalendar) (*externalRaceData, error) {
	newMeetings := make([]*racing.Meeting, 1)
	newRaces := make([]*racing.Race, 1)

	existingMeetings := make([]*racing.Meeting, 1)
	existingRaces := make([]*racing.Race, 1)

	for _, rg := range c.RegionGroups {
		for _, m := range rg.Meetings {
			firstEvent := m.Events[0]

			mSource, err := getMeetingSourceID(firstEvent.DateWithYear, firstEvent.EventURL)
			if err != nil {
				return nil, fmt.Errorf("unable to create source id for date %v and url %v", firstEvent.DateWithYear, firstEvent.EventURL)
			}

			meeting := &racing.Meeting{
				Name:           m.MeetingName,
				ScheduledStart: firstEvent.StartTime,
				RaceType:       eventType,
				Country:        m.RegionDescription,
				SourceId:       mSource,
				RaceIds:        make([]string, 1),
			}

			if val, ok := p.meetingsBySource[mSource]; ok {
				meeting.MeetingId = val.MeetingId
				existingMeetings = append(existingMeetings, meeting)
			} else {
				meeting.MeetingId = uuid.NewV4().String()
				newMeetings = append(newMeetings, meeting)
			}

			for _, e := range m.Events {
				rSource := string(e.EventID)
				race := &racing.Race{
					MeetingId:      meeting.MeetingId,
					Name:           e.EventName,
					Number:         e.EventNumber,
					Results:        e.Results,
					ScheduledStart: e.StartTime,
					Status:         getRaceStatusFromCalendar(e.IsAbandoned, e.Resulted, e.Results),
					SourceId:       rSource,
				}

				if val, ok := p.racesBySource[rSource]; ok {
					race.RaceId = val.RaceId
					existingRaces = append(existingRaces, race)
				} else {
					race.RaceId = uuid.NewV4().String()
					newRaces = append(newRaces, race)
				}

				meeting.RaceIds = append(meeting.RaceIds, race.RaceId)
			}
		}
	}

	data := &externalRaceData{
		existingMeetings,
		existingRaces,
		newMeetings,
		newRaces,
	}

	return data, nil
}

func parseRaceCard(c *RaceCard) ([]*racing.Selection, error) {
	selections := make([]*racing.Selection, len(c.Selections))
	for _, v := range c.Selections {

		number, err := strconv.Atoi(v.CompetitorNumber)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse competitor number '%v'", v.CompetitorNumber)
		}

		barrier, err := strconv.Atoi(v.BarrierNumber)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse barrier number '%v'", v.BarrierNumber)
		}

		s := &racing.Selection{
			SourceId:           v.SelectionID,
			BarrierNumber:      int32(barrier),
			Jockey:             v.JockeyName,
			Number:             int32(number),
			SourceCompetitorId: v.CompetitorID,
		}

		selections = append(selections, s)
	}

	return selections, nil
}

func getResults(rStr string) ([]int32, error) {
	if rStr == "" {
		return nil, nil
	}

	results := make([]int32, 1)
	for _, v := range strings.Split(rStr, ",") {
		r, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		results = append(results, int32(r))
	}

	return results, nil
}

func getMeetingSourceID(date, url string) (string, error) {
	// url is in the format '\/horse-racing\/bendigo\/race-1\/'
	// date is in the format '20 May 2018'
	// This should return a source id of '20180520-bendigo-horse-racing'

	u := strings.Split(url, "\\/")
	d, err := time.Parse("02 Jan 2016", date)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v-%v-%v", u[0], u[1], d), nil
}

func getRaceStatusFromCalendar(isAbandoned int32, resulted int32, results string) string {
	if isAbandoned == 1 {
		return "ABANDONED"
	} else if resulted == 1 && results != "" {
		return "INTERIM"
	} else if resulted == 1 {
		return "CLOSED"
	}

	return "OPEN"
}

func meetingChanged(from, to *racing.Meeting) bool {
	return from.Name != to.Name ||
		from.ScheduledStart != to.ScheduledStart
}

func raceChanged(from, to *racing.Race) bool {
	return from.ScheduledStart != to.ScheduledStart ||
		from.ActualStart != to.ActualStart ||
		from.Status != to.Status ||
		from.Results != to.Results
}

func createNewMeetings(client racing.RacingService, meetings []*racing.Meeting) error {
	ctx := context.Background()
	req := &racing.AddMeetingsRequest{
		Meetings: meetings,
	}
	_, err := client.AddMeetings(ctx, req)
	return err
}

func createNewRaces(client racing.RacingService, races []*racing.Race) error {
	ctx := context.Background()
	req := &racing.AddRacesRequest{
		MeetingId: races[0].MeetingId,
		Races:     races,
	}
	_, err := client.AddRaces(ctx, req)
	return err
}

func updateExistingMeetings(client racing.RacingService, meetings []*racing.Meeting) error {
	return nil // TODO: do meetings get updated??
}

func updateExistingRaces(client racing.RacingService, races []*racing.Race) error {
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(len(races))
	var errLock = &sync.Mutex{}

	errors := make([]error, 1)

	for _, v := range races {
		go func(r *racing.Race) {
			defer wg.Done()
			req := &racing.UpdateRaceRequest{
				Race: r,
			}
			_, err := client.UpdateRace(ctx, req)
			if err != nil {
				errLock.Lock()
				errors = append(errors, err)
				errLock.Unlock()
			}
		}(v)
	}

	wg.Wait()
	return merge(errors)
}

func merge(e []error) error {
	if len(e) == 0 {
		return nil
	}

	var errStr []string
	for _, v := range e {
		errStr = append(errStr, v.Error())
	}

	return fmt.Errorf(strings.Join(errStr, "\n"))
}
