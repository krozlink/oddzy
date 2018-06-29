package main

import (
	"context"
	"fmt"
	racing "github.com/krozlink/oddzy/services/srv/racing/proto"
	microclient "github.com/micro/go-micro/client"
	_ "github.com/micro/go-plugins/registry/consul"
	"github.com/satori/go.uuid"
	"strconv"
	"strings"
	"sync"
	"time"
)

type scrapeProcess struct {
	status string
	done   chan bool
	// http      handler
	racing    racing.RacingService
	scraper   Scraper
	dateRange [2]int
	today     time.Time

	upcoming []*racing.Race

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

const (
	scrapeRaceTiming  = "race-scraper.service.scraperace.timing"
	scrapeRaceSuccess = "race-scraper.service.scraperace.success"
	scrapeRaceFailed  = "race-scraper.service.scraperace.failed"
)

// TODO put harness and greyhounds back
var raceTypes = []string{"horse-racing"} //, "harness", "greyhounds"}

func (p *scrapeProcess) run() {
	log := logWithField("function", "run")
loop:
	for { // this should only loop once a day, with the loop ending when the day ends
		// calculate date range
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		tomorrow := today.Add(time.Hour * 24).Sub(time.Now())
		start, end := getDateRange(p.dateRange)

		log.Infof("Reading races between %v and %v", start.Format("02 Jan 2006"), end.Format("02 Jan 2006"))
		open, missing := readRaces(p, start, end)

		log.Infof("There are %v races requiring scraping", len(missing))
		scrapeRaces(p, missing)

		stop, done := monitorOpenRaces(p, open)
		select {
		case <-done: // something else has killed off the monitor
			log.Info("Scraper flagged as done")
			break loop
		case <-time.After(tomorrow): // channel starting following day
			log.Info("Day complete. Exit current scraping loop")
			stop <- true // kill the current scrape session
			<-done       // wait for it to exit
			reset(p)     // reset the scraper back to its inital state

			log.Info("Scraping loop ended. Start scraping for next day.")
			// No point in rescraping dates that were scraped the previous round
			// Just need to scrape the end day of the previous round
			p.dateRange = [2]int{p.dateRange[1], p.dateRange[1]}
		}
	}
}

func reset(p *scrapeProcess) {
	p.status = "DAY_RESET"
	p.meetingsByID = make(map[string]*racing.Meeting)
	p.meetingsBySource = make(map[string]*racing.Meeting)
	p.racesByID = make(map[string]*racing.Race)
	p.racesBySource = make(map[string]*racing.Race)
}

func scrapeRaces(p *scrapeProcess, missing []*racing.Race) error {
	p.status = "RACE_CREATION"

	// queue all races that require scraping
	scrape := make(chan *racing.Race, 1000)
	go func() {
		for _, r := range missing {
			scrape <- r
		}
		close(scrape)
	}()

	update := make(chan *racing.UpdateRaceRequest, 1000)
	var lock = &sync.Mutex{}
	var errors []error

	// scrape each race and queue for updates
	go func() {
		for r := range scrape {
			t := stats.NewTiming()
			card, err := p.scraper.ScrapeRaceCard(r.SourceId)
			if err != nil {
				lock.Lock()
				errors = append(errors, err)
				lock.Unlock()
			}
			selections := parseRaceCard(card)
			for _, s := range selections {
				s.SelectionId = uuid.NewV4().String()
				s.RaceId = r.RaceId
			}

			req := &racing.UpdateRaceRequest{
				ActualStart:    r.ActualStart,
				RaceId:         r.RaceId,
				Results:        r.Results,
				ScheduledStart: r.ScheduledStart,
				Status:         r.Status,
				Selections:     selections,
			}

			if err == nil {
				stats.Increment(scrapeRaceSuccess)
			} else {
				stats.Increment(scrapeRaceFailed)
			}
			t.Send(scrapeRaceTiming)
			update <- req
		}
		close(update)
	}()

	// Updated scraped races
	ctx := context.Background()
	for r := range update {
		_, err := p.racing.UpdateRace(ctx, r)
		if err != nil {
			lock.Lock()
			errors = append(errors, err)
			lock.Unlock()
		}
	}

	err := merge(errors)
	return err
}

func getDateRange(r [2]int) (time.Time, time.Time) {
	//r[0] - number of days ago to start scraping from. e.g -1 is yesterday
	//r[1] - number of days ago to scrape to. e.g 2 is day after tomorrow
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Add(time.Hour * time.Duration(-24*r[0]))
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location()).Add(time.Hour * time.Duration(24*r[1]))
	return start, end
}

func newScrapeProcess() scrapeProcess {

	r := racing.NewRacingService(racingServiceName, microclient.DefaultClient)

	return scrapeProcess{
		status: "INITIALISING",
		done:   make(chan bool),
		// http:      handler{},
		racing:  r,
		scraper: NewOddsScraper(&handler{}),
		//TODO - change back to -1, 2
		dateRange:        [2]int{0, 0}, // scrape from 1 day ago to 2 days in the future (4 days total)
		meetingsByID:     make(map[string]*racing.Meeting),
		meetingsBySource: make(map[string]*racing.Meeting),
		racesByID:        make(map[string]*racing.Race),
		racesBySource:    make(map[string]*racing.Race),
	}
}

func readRaces(p *scrapeProcess, start, end time.Time) ([]*racing.Race, []*racing.Race) {

	log := logWithField("function", "readRaces")

	// STATUS - SETUP
	p.status = "SETUP"

	// read all internal meeting data for scraping period (yesterday to 2 days from now)
	meetings, races, err := readInternal(p, start, end)
	if err != nil {
		log.Fatalf("An error occurred reading internal racing data - %v", err)
	}

	for _, m := range meetings {
		log.Debugf("Internal meeting with source id %v and meeting id %v", m.SourceId, m.MeetingId)
		p.meetingsByID[m.MeetingId] = m
		p.meetingsBySource[m.SourceId] = m
	}

	for _, r := range races {
		log.Debugf("Internal race with source id %v and race id %v", r.SourceId, r.RaceId)
		p.racesByID[r.RaceId] = r
		p.racesBySource[r.SourceId] = r
	}

	log.Infof("Internal - %v meetings and %v races found", len(meetings), len(races))

	// read all external meeting data for scraping period (yesterday to 2 days from now)
	data, err := readExternal(p, start, end)
	if err != nil {
		log.Errorf("An error occurred reading external racing data - %v", err)
	}
	log.Infof("External (existing) - %v meetings and %v races found", len(data.existingMeetings), len(data.existingRaces))
	log.Infof("External (new) - %v meetings and %v races found", len(data.newMeetings), len(data.newRaces))

	// update all existing meetings that have changed
	uMeetings := make([]*racing.Meeting, 0)
	for _, m := range data.existingMeetings {
		current := p.meetingsByID[m.MeetingId]
		if meetingChanged(current, m) {
			uMeetings = append(uMeetings, m)
		}
	}

	unscraped := make([]*racing.Race, 0)
	open := make([]*racing.Race, 0)

	// update all existing races that have changed
	uRaces := make([]*racing.Race, 0)
	for _, r := range data.existingRaces {
		current := p.racesByID[r.RaceId]
		if raceChanged(current, r) {
			uRaces = append(uRaces, r)
		}

		// flag existing races that have not been individually scraped yet
		if current.IsScraped == false {
			unscraped = append(unscraped, r)
		}

		if r.Status == "OPEN" {
			open = append(open, r)
		}
	}

	for _, r := range data.newRaces {
		if r.Status == "OPEN" {
			open = append(open, r)
		}
	}

	updateExistingMeetings(p.racing, uMeetings)
	updateExistingRaces(p.racing, uRaces)

	createNewMeetings(p.racing, data.newMeetings)
	createNewRaces(p.racing, data.newRaces)

	// flag new races that have not been individually scraped yet
	for _, r := range data.newRaces {
		unscraped = append(unscraped, r)
	}

	return open, unscraped
}

// readExternal reads race and meeting data from odds.com.au
func readExternal(p *scrapeProcess, start, end time.Time) (*externalRaceData, error) {
	data := &externalRaceData{}

	// read race calendars from scraping period
	for _, t := range raceTypes {
		for d := start; d.Before(end) || d.Equal(end); d = d.Add(time.Hour * 24) {
			cal, err := p.scraper.ScrapeRaceCalendar(t, d.Format("2006-01-02"))
			if err != nil {
				return nil, fmt.Errorf("Unable to read race calendar for %v on %v - %v", t, d, err)
			}

			c, err := processRaceCalendar(p, t, cal)
			if err != nil {
				return nil, fmt.Errorf("Unable to process race calendar for %v on %v - %v", t, d, err)
			}

			data.existingMeetings = append(data.existingMeetings, c.existingMeetings...)
			data.existingRaces = append(data.existingRaces, c.existingRaces...)
			data.newMeetings = append(data.newMeetings, c.newMeetings...)
			data.newRaces = append(data.newRaces, c.newRaces...)
		}
	}

	return data, nil
}

// readInternal reads race and meeting data stored in oddzy
func readInternal(p *scrapeProcess, start, end time.Time) ([]*racing.Meeting, []*racing.Race, error) {
	log := logWithField("function", "readInternal")
	// read all internal meeting data for scraping period (yesterday to 2 days from now)
	//		ListMeetingsByDate
	//		ListRacesByMeetingDate
	ctx := context.Background()
	mReq := &racing.ListMeetingsByDateRequest{
		StartDate: start.Unix(),
		EndDate:   end.Unix(),
	}
	mResp, err := p.racing.ListMeetingsByDate(ctx, mReq)
	if err != nil {
		log.Errorf("Error calling ListMeetingsByDate - %v", err)
		return nil, nil, err
	}

	rReq := &racing.ListRacesByMeetingDateRequest{
		StartDate: start.Unix(),
		EndDate:   end.Unix(),
	}
	rResp, err := p.racing.ListRacesByMeetingDate(ctx, rReq)
	if err != nil {
		log.Errorf("Error calling ListRacesByMeetingDate - %v", err)
		return nil, nil, err
	}

	return mResp.Meetings, rResp.Races, nil
}

func processRaceCalendar(p *scrapeProcess, eventType string, c *RaceCalendar) (*externalRaceData, error) {
	log := logWithField("function", "processRaceCalendar")

	newMeetings := make([]*racing.Meeting, 0)
	newRaces := make([]*racing.Race, 0)

	existingMeetings := make([]*racing.Meeting, 0)
	existingRaces := make([]*racing.Race, 0)

	if !c.HasResults {
		data := &externalRaceData{
			existingMeetings,
			existingRaces,
			newMeetings,
			newRaces,
		}

		return data, nil
	}

	for _, rg := range c.RegionGroups {
		for _, m := range rg.Meetings {
			firstEvent := m.Events[0]

			mSource, err := getMeetingSourceID(firstEvent.DateWithYear, firstEvent.EventURL)
			if err != nil {
				log.Error(err)
				return nil, fmt.Errorf("unable to create source id for date %v and url %v", firstEvent.DateWithYear, firstEvent.EventURL)
			}

			meeting := &racing.Meeting{
				Name:           m.MeetingName,
				ScheduledStart: firstEvent.StartTime,
				RaceType:       eventType,
				Country:        m.RegionDescription,
				SourceId:       mSource,
				RaceIds:        make([]string, 0),
			}

			if val, ok := p.meetingsBySource[mSource]; ok {
				log.Debugf("Meeting with source id %v already exists with id %v", mSource, val.MeetingId)
				meeting.MeetingId = val.MeetingId
				meeting.DateCreated = val.DateCreated
				meeting.LastUpdated = val.LastUpdated
				existingMeetings = append(existingMeetings, meeting)
			} else {
				log.Debugf("Meeting with source id %v is new", mSource)
				meeting.MeetingId = uuid.NewV4().String()
				newMeetings = append(newMeetings, meeting)
			}

			for _, e := range m.Events {
				rSource := strconv.Itoa(int(e.EventID))
				race := &racing.Race{
					MeetingId:      meeting.MeetingId,
					Name:           e.EventName,
					Number:         e.EventNumber,
					Results:        e.Results,
					ScheduledStart: e.StartTime,
					Status:         getRaceStatusFromCalendar(e.IsAbandoned, e.Resulted, e.Results),
					SourceId:       rSource,
					MeetingStart:   meeting.ScheduledStart,
				}

				if val, ok := p.racesBySource[rSource]; ok {
					race.RaceId = val.RaceId
					race.DateCreated = val.DateCreated
					race.LastUpdated = val.LastUpdated
					existingRaces = append(existingRaces, race)
				} else {
					race.RaceId = uuid.NewV4().String()
					newRaces = append(newRaces, race)
				}

				meeting.RaceIds = append(meeting.RaceIds, race.RaceId)
			}
		}
	}

	log.Debugf("Race calendar processed.   Existing Meetings: %v    Existing Races: %v    New Meetings: %v    New Races: %v", len(existingMeetings), len(existingRaces), len(newMeetings), len(newRaces))

	data := &externalRaceData{
		existingMeetings,
		existingRaces,
		newMeetings,
		newRaces,
	}

	return data, nil
}

// parseRaceCard parses the selections (horses/greyhounds) from a race card
func parseRaceCard(c *RaceCard) []*racing.Selection {
	selections := make([]*racing.Selection, len(c.Selections))
	for i, v := range c.Selections {

		number, err := strconv.Atoi(v.CompetitorNumber)
		if err != nil {
			number = 0
		}

		barrier, err := strconv.Atoi(v.BarrierNumber)
		if err != nil {
			barrier = 0
		}

		s := &racing.Selection{
			RaceId:             "",
			Name:               v.Name,
			Scratched:          false,
			SourceId:           v.SelectionID,
			BarrierNumber:      int32(barrier),
			Jockey:             v.JockeyName,
			Number:             int32(number),
			SourceCompetitorId: v.CompetitorID,
		}

		selections[i] = s
	}

	return selections
}

// getResults takes a string result in the format "3,1,12" and returns the equivalent slice of ints
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

// getMeetingSourceID takes the date and url of an event on odds.com.au and generates a source id
// this source id is used as the link between the scraped meeting data and the internally stored meeting data
func getMeetingSourceID(date, url string) (string, error) {
	// url is in the format '/horse-racing/bendigo/race-1/'
	// date is in the format '20 May 2018'
	// This should return a source id of '20180520-bendigo-horse-racing'

	u := strings.Split(url, "/")
	d, err := time.Parse("02 Jan 2006", date)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v-%v-%v", d.Format("20060102"), u[2], u[1]), nil
}

func getRaceStatusFromCalendar(isAbandoned int32, resulted int32, results string) string {
	if isAbandoned == 1 {
		return "ABANDONED"
	} else if resulted == 1 && results == "" {
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

func createNewMeetings(client racing.RacingService, meetings []*racing.Meeting) {
	log := logWithField("function", "createNewMeetings")

	if len(meetings) > 0 {
		log.Infof("%v new meetings have been added - updating repository", len(meetings))
	} else {
		log.Infof("No meetings have been added")
		return
	}

	ctx := context.Background()
	req := &racing.AddMeetingsRequest{
		Meetings: meetings,
	}
	if _, err := client.AddMeetings(ctx, req); err != nil {
		log.Fatal(err)
	}
}

func createNewRaces(client racing.RacingService, races []*racing.Race) {
	log := logWithField("function", "createNewRaces")

	if len(races) > 0 {
		log.Infof("%v new races have been added - updating repository", len(races))
	} else {
		log.Infof("No races have been added")
		return
	}

	ctx := context.Background()
	req := &racing.AddRacesRequest{
		Races: races,
	}

	if _, err := client.AddRaces(ctx, req); err != nil {
		log.Fatal(err)
	}
}

func updateExistingMeetings(client racing.RacingService, meetings []*racing.Meeting) {
	log := logWithField("function", "updateExistingMeetings")

	if len(meetings) > 0 {
		log.Infof("%v meetings have changed - updating repository", len(meetings))
		log.Warn("updating existing races not currently required or supported")
	} else {
		log.Infof("No meetings have changed")
		return
	}

	return // TODO: do meetings get updated??
}

func updateExistingRaces(client racing.RacingService, races []*racing.Race) {
	log := logWithField("function", "updateExistingRaces")

	if len(races) > 0 {
		log.Infof("%v races have changed - updating repository", len(races))
	} else {
		log.Infof("No races have changed")
		return
	}

	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(len(races))
	var errLock = &sync.Mutex{}

	var errors []error

	for _, v := range races {
		v.LastUpdated = time.Now().Unix()

		go func(r *racing.Race) {
			defer wg.Done()
			req := &racing.UpdateRaceRequest{
				ActualStart:    r.ActualStart,
				RaceId:         r.RaceId,
				Results:        r.Results,
				ScheduledStart: r.ScheduledStart,
				Status:         r.Status,
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
	if err := merge(errors); err != nil {
		log.Fatal(err)
	}
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
