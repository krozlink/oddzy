package main

import (
	"context"
	// proto "github.com/krozlink/oddzy/services/srv/race-scraper/proto"
	"fmt"
	racing "github.com/krozlink/oddzy/services/srv/racing/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/consul"
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
}

var raceTypes = []string{"horses", "harness", "greyhounds"}

func (p *scrapeProcess) run() {

	setup(p)

	// STATUS - RACE_CREATION
	// For each item on the missingEvents queue
	//	Scrape the race
	// 	Update the race - UpdateRace

	// STATUS - RACE MONITORING
	// Process 1
	//		Monitor time until race and race status and push required races to queue if they are not already on there
	//		Only needs to run at most every 30 seconds. It can determine its own sleep duration depending on upcoming races
	// Process 2 - Periodically take items off the queue and scrape them. Ensure minimum interval to avoid overloading server

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
	p.status = "SETUP"
	p.startDate, p.endDate = getDateRange()
	m, r, err := readInternal(p)
	cal, err := readExternal(p)

	//TODO: what is unique meeting id on racecard response?? is meetingName unique??
	// should readExternal also convert race calendar to []meeting and []races?

	// STATUS - SETUP
	// read all internal meeting data for scraping period (yesterday to 2 days from now)
	//		ListMeetingsByDate
	//		ListRacesByMeetingDate

	// for each external meeting
	// 		if the meeting doesnt exist internally (look up source id), create meeting id and add to new meetings list

	// for each race
	// 		if the race doesnt exist internally (look up source id) create a race id and add to new races list

	// update new meeting list with their race ids
	// call AddMeetings with new meetings list
	// call AddRaces with  new races list (leave last_updated as null)

	// Add all existing races with a null last_updated as well as all new races to a missingEvents queue

}

func readExternal(p *scrapeProcess) ([]*RaceCalendar, error) {
	calendars := make([]*RaceCalendar, len(raceTypes)*4)
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

			calendars = append(calendars, cal)
		}
	}

	return calendars, nil
}

func readInternal(p *scrapeProcess) ([]*racing.Meeting, []*racing.Race, error) {
	// read all internal meeting data for scraping period (yesterday to 2 days from now)
	//		ListMeetingsByDate
	//		ListRacesByMeetingDate
	ctx := context.Background()
	mReq := &racing.ListMeetingsByDateRequest{
		StartDate: p.startDate.Format("2006-01-02"),
		EndDate:   p.endDate.Format("2006-01-02"),
	}
	mResp, err := p.racing.ListMeetingsByDate(ctx, mReq)
	if err != nil {
		return nil, nil, err
	}

	rReq := &racing.ListRacesByMeetingDateRequest{
		StartDate: p.startDate.Format("2006-01-02"),
		EndDate:   p.endDate.Format("2006-01-02"),
	}
	rResp, err := p.racing.ListRacesByMeetingDate(ctx, rReq)
	if err != nil {
		return nil, nil, err
	}

	return mResp.Meetings, rResp.Races, nil
}

func parseRaceCalendar(c *RaceCalendar) ([]*racing.Meeting, []*racing.Race, error) {
	//TODO parse race calendar
}

func parseRaceCard(c *RaceCard) ([]*racing.Meeting, []*racing.Race, error) {
	//TODO parse race card
}
