package main

import (
	"context"
	"encoding/json"
	"fmt"
	racing "github.com/krozlink/oddzy/services/srv/racing/proto"
	client "github.com/micro/go-micro/client"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestReadInternalReturnsData(t *testing.T) {

	meetings := []*racing.Meeting{
		getTestMeeting("m123", "source-123", 2),
		getTestMeeting("m456", "source-456", 1),
	}
	races := []*racing.Race{
		getTestRace(meetings[0], 1001, 1),
		getTestRace(meetings[0], 1001, 2),
		getTestRace(meetings[1], 2001, 1),
	}

	client := &mockRacingClient{
		meetings: meetings,
		races:    races,
	}
	p := getTestProcess(client, &mockScraper{})

	start, end := getDateRange([2]int{0, 0})
	intMeetings, intRaces, err := readInternal(p, start, end)
	if err != nil {
		t.Error(err)
	}

	if len(intMeetings) != len(meetings) {
		t.Errorf("Expected %v meetings but got %v", len(meetings), len(intMeetings))
	}

	if len(intRaces) != len(races) {
		t.Errorf("Expected %v races but got %v", len(races), len(intRaces))
	}

	for i, o := range meetings {
		m := intMeetings[i]
		if o.Country != m.Country || // compare meeting to original
			o.MeetingId != m.MeetingId ||
			o.Name != m.Name ||
			o.RaceType != m.RaceType ||
			o.ScheduledStart != m.ScheduledStart ||
			o.SourceId != m.SourceId {
			t.Errorf("Meetings have unexpected difference")
		}
	}

	for i, o := range races {
		r := intRaces[i]
		if o.ActualStart != r.ActualStart ||
			o.DateCreated != r.DateCreated ||
			o.LastUpdated != r.LastUpdated ||
			o.MeetingId != r.MeetingId ||
			o.MeetingStart != r.MeetingStart ||
			o.Name != r.Name ||
			o.Number != r.Number ||
			o.RaceId != r.RaceId ||
			o.Results != r.Results ||
			o.ScheduledStart != r.ScheduledStart ||
			o.SourceId != r.SourceId ||
			o.Status != r.Status {
			t.Errorf("Races have unexpected difference")
		}
	}
}

func TestReadExternalReturnsData(t *testing.T) {
	c := &mockRacingClient{}
	s := &mockScraper{}

	// Not an ideal test but working for now
	// There would be 3 calls to ScrapeRaceCalendar - 1 day x 3 race types.
	// This mocked version will return each one incrementally
	calendars := []*RaceCalendar{
		getTestRaceCalendar(1),
		getTestRaceCalendar(2),
		getTestRaceCalendar(3),
	}

	s.calendars = calendars

	p := getTestProcess(c, s)
	start, end := getDateRange([2]int{0, 0}) // only read today

	data, err := readExternal(p, start, end)

	assert.Equal(t, len(calendars), s.scrapeCalendarCount, "Unexpected calendar call count")
	assert.NoError(t, err)

	// As no internal meetings have been mocked it is expected that all meetings will be treated as new
	assert.Equal(t, 0, len(data.existingMeetings), "Unexpected existing meetings")
	assert.Equal(t, 0, len(data.existingRaces), "Unexpected existing races")

	assert.Equal(t, len(calendars), len(data.newMeetings), "Unexpected new meetings")
	assert.Equal(t, len(calendars), len(data.newRaces), "Unexpected new races")
}

func TestUpdateExistingRacesHandlesMultipleRaces(t *testing.T) {
	c := &mockRacingClient{}

	// create 3 races and update them
	m := getTestMeeting("m123", "meeting-123", 3)
	races := []*racing.Race{
		getTestRace(m, 1001, 1),
		getTestRace(m, 1002, 2),
		getTestRace(m, 1003, 3),
	}

	updateExistingRaces(c, races)

	// confirm update race was called 3 times
	assert.Equal(t, len(races), c.updateRaceCount)
}

func TestProcessRaceCalendarSplitsNewAndExisting(t *testing.T) {

	// Create a mock race calendar with a two meetings each with one event
	// 1st meeting has a url/date giving it a source id of "20180604-test-park-greyhounds"
	// 	its event has a source id of 1234
	// 2nd meeting has a url/date giving it a source id of "20180101-brisbane-greyhounds"
	// 	its event has a source id of 4567

	meeting1 := Meeting{
		MeetingName:       "Meeting 1",
		RegionDescription: "Australia",
		RegionIconURL:     "www.example.com",
		Events: []Event{
			getTestCalendarEvent(1, 1234, "/greyhounds/test-park/race-1/", "04 Jun 2018"),
		},
	}

	meeting2 := Meeting{
		MeetingName:       "Meeting 2",
		RegionDescription: "Australia",
		RegionIconURL:     "www.example.com",
		Events: []Event{
			getTestCalendarEvent(1, 4567, "/greyhounds/brisbane/race-1/", "01 Jan 2018"),
		},
	}

	cal := &RaceCalendar{
		HasResults: true,
		RegionGroups: []RegionGroup{
			{
				GroupName: "Aussie Horses",
				Meetings: []Meeting{
					meeting1,
					meeting2,
				},
			},
		},
	}

	// Mock the internal repo to have the first meeting/race but not the second
	meetings := []*racing.Meeting{
		getTestMeeting("20180604-test-park-greyhounds", "m123", 2),
	}
	races := []*racing.Race{
		getTestRace(meetings[0], 1234, 1),
	}

	client := &mockRacingClient{
		meetings: meetings,
		races:    races,
	}
	p := getTestProcess(client, &mockScraper{})

	for _, m := range meetings {
		p.meetingsByID[m.MeetingId] = m
		p.meetingsBySource[m.SourceId] = m
	}

	for _, r := range races {
		p.racesByID[r.RaceId] = r
		p.racesBySource[r.SourceId] = r
	}

	data, err := processRaceCalendar(p, "horse-racing", cal)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, len(data.newMeetings), "Unexpected number of new meetings")
	assert.Equal(t, 1, len(data.existingMeetings), "Unexpected number of existing meetings")

	assert.Equal(t, meeting2.MeetingName, data.newMeetings[0].Name, "Unexpected meeting name")
	assert.Equal(t, meeting2.RegionDescription, data.newMeetings[0].Country, "Unexpected country")
	assert.Equal(t, len(meeting2.Events), len(data.newRaces), "Unexpected number of races")

	assert.Equal(t, meeting1.MeetingName, data.existingMeetings[0].Name, "Unexpected meeting name")
	assert.Equal(t, meeting1.RegionDescription, data.existingMeetings[0].Country, "Unexpected country")
	assert.Equal(t, len(meeting1.Events), len(data.existingRaces), "Unexpected number of races")

	assert.Equal(t, 1, len(data.newRaces), "Unexpected number of new races")
	assert.Equal(t, 1, len(data.existingRaces), "Unexpected number of existing races")
}

func TestGetMeetingSourceID(t *testing.T) {

	var tests = []struct {
		url      string
		date     string
		expected string
	}{
		{"/horse-racing/sapphire-coast/race-1/", "10 Jun 2018", "20180610-sapphire-coast-horse-racing"},
		{"/greyhounds/albion-park/race-1/", "04 Jun 2018", "20180604-albion-park-greyhounds"},
		{"/harness/wentworth/race-13/", "31 Jan 2022", "20220131-wentworth-harness"},
	}

	for _, v := range tests {
		id, err := getMeetingSourceID(v.date, v.url)
		assert.NoError(t, err)
		assert.Equal(t, v.expected, id, "Unexpected source id")
	}
}

func TestGetRaceStatusFromCalendar(t *testing.T) {

	var tests = []struct {
		isAbandoned int32
		resulted    int32
		results     string
		expected    string
	}{
		{1, 0, "", "ABANDONED"},
		{0, 1, "", "INTERIM"},
		{0, 1, "2,4,3", "CLOSED"},
		{0, 0, "", "OPEN"},
	}

	for _, v := range tests {
		result := getRaceStatusFromCalendar(v.isAbandoned, v.resulted, v.results)
		assert.Equal(t, v.expected, result, "Unexpected race status")
	}
}

func TestParseRaceCardReadsSelections(t *testing.T) {

	// add test race card with 1 selection
	card := getTestRaceCard(1, 1)

	s, err := parseRaceCard(card)

	// confirm 1 selection parsed
	assert.NoError(t, err)
	assert.Equal(t, 1, len(s))
}

func TestScrapeRacesScrapesAllMissingRaces(t *testing.T) {
	c := &mockRacingClient{}
	s := &mockScraper{}

	// inject 3 responses to scraping a race card
	s.cards = []*RaceCard{
		getTestRaceCard(1, 1),
		getTestRaceCard(2, 2),
		getTestRaceCard(3, 3),
	}

	p := getTestProcess(c, s)

	// inject 3 races to scrape
	m := getTestMeeting("m123", "meeting-123", 3)
	races := []*racing.Race{
		getTestRace(m, 1001, 1),
		getTestRace(m, 1002, 2),
		getTestRace(m, 1003, 3),
	}

	// scrape 3 races and confirm that ScrapeRaceCard was called 3 times
	scrapeRaces(p, races)

	assert.Equal(t, len(races), s.scrapeRaceCardCount)
}

func TestScrapeRacesUpdatesScrapedRaces(t *testing.T) {
	c := &mockRacingClient{}
	s := &mockScraper{}

	// inject 3 responses to scraping a race card
	s.cards = []*RaceCard{
		getTestRaceCard(1, 1),
		getTestRaceCard(2, 2),
		getTestRaceCard(3, 3),
	}

	p := getTestProcess(c, s)

	m := getTestMeeting("m123", "meeting-123", 3)
	races := []*racing.Race{
		getTestRace(m, 1001, 1),
		getTestRace(m, 1002, 2),
		getTestRace(m, 1003, 3),
	}

	// scrape 3 races and confirm that ScrapeRaceCard was called 3 times
	scrapeRaces(p, races)

	assert.Equal(t, len(races), c.updateRaceCount)
}

func getTestProcess(c *mockRacingClient, s *mockScraper) *scrapeProcess {
	return &scrapeProcess{
		status:           "TEST",
		done:             make(chan bool),
		http:             handler{},
		racing:           c,
		scraper:          s,
		meetingsByID:     make(map[string]*racing.Meeting),
		meetingsBySource: make(map[string]*racing.Meeting),
		racesByID:        make(map[string]*racing.Race),
		racesBySource:    make(map[string]*racing.Race),
	}
}

func getTestMeeting(sourceID, meetingID string, numRaces int) *racing.Meeting {
	raceIds := make([]string, numRaces)
	for i := 0; i < numRaces; i++ {
		raceIds[i] = fmt.Sprintf("%v-%v", meetingID, (i + 1))
	}

	return &racing.Meeting{
		Country:        "TESTLAND",
		MeetingId:      meetingID,
		Name:           "Test Meeting 1",
		RaceIds:        raceIds,
		RaceType:       "test-type",
		ScheduledStart: 1234,
		SourceId:       sourceID,
	}
}

func getTestRace(m *racing.Meeting, eventID, number int32) *racing.Race {
	var result string
	var status string
	if number%2 == 0 {
		result = ""
		status = "open"
	} else {
		result = fmt.Sprintf("%v,%v,%v", number, number+1, number+2)
		status = "closed"
	}

	return &racing.Race{
		RaceId:         m.RaceIds[number-1],
		MeetingId:      m.MeetingId,
		Name:           fmt.Sprintf("Race number %v", number),
		Number:         number,
		Results:        result,
		ScheduledStart: int64(number * 1000),
		SourceId:       string(eventID),
		Status:         status,
	}
}

func getTestCalendarEvent(number, eventID int32, url, date string) Event {
	return Event{
		EventID:      eventID,
		DateWithYear: date,
		EventName:    "Event number " + string(number),
		EventNumber:  number,
		EventURL:     url,
		IsAbandoned:  0,
		IsImminent:   false,
		Resulted:     0,
		Results:      "",
		StartTime:    1000,
	}
}

func readTestRaceCalendar() *RaceCalendar {
	decoded, _ := ioutil.ReadFile("./testdata/race_calendar_decoded.json")
	cal := &RaceCalendar{}
	json.Unmarshal(decoded, cal)
	return cal
}

func getTestRaceCalendar(id int) *RaceCalendar {
	meeting := Meeting{
		MeetingName:       "Meeting " + string(id),
		RegionDescription: "Australia",
		RegionIconURL:     "www.example.com",
		Events: []Event{
			getTestCalendarEvent(1, int32(1000+id), "/greyhounds/meeting-"+string(id)+"/race-1/", "04 Jun 2018"),
		},
	}

	cal := &RaceCalendar{
		HasResults: true,
		RegionGroups: []RegionGroup{
			{
				GroupName: "Aussie Horses",
				Meetings: []Meeting{
					meeting,
				},
			},
		},
	}

	return cal
}

func getTestRaceCard(id, number int32) *RaceCard {
	prices := []RacePrice{}
	selections := []RaceSelection{
		RaceSelection{
			BarrierNumber:    "1",
			CompetitorID:     "2",
			CompetitorNumber: "3",
			Flucs:            "2.00,2.05,2.1",
			ImageURL:         "example.com",
			JockeyName:       "Bob",
			JockeyURL:        "bob.com",
			JockeyWeight:     "40kg",
			Name:             "Winx",
			Prices:           prices,
			ProfileURL:       "winx.com",
			Result:           "",
			ResultOrdinal:    "",
			SelectionID:      "123",
			Weight:           "10032kg",
		},
	}
	card := &RaceCard{
		EventDescription: "race number " + string(number),
		EventDistance:    "1000m",
		EventID:          id,
		EventName:        "race number " + string(number),
		EventNameFull:    "race number " + string(number),
		IsGreyhounds:     true,
		IsHarness:        false,
		IsHorseRacing:    false,
		IsRacing:         true,
		ResultState:      "open",
		Selections:       selections,
		SportID:          1,
		StartTime:        1000,
		Status:           "open",
	}

	return card
}

type mockScraper struct {
	scrapeCalendarCount int
	scrapeRaceCardCount int

	cards     []*RaceCard
	calendars []*RaceCalendar
}

func (c *mockScraper) ScrapeRaceCard(sourceID string) (*RaceCard, error) {
	card := c.cards[c.scrapeRaceCardCount]
	c.scrapeRaceCardCount++
	return card, nil
}

func (c *mockScraper) ScrapeRaceCalendar(eventType string, date string) (*RaceCalendar, error) {
	cal := c.calendars[c.scrapeCalendarCount]
	c.scrapeCalendarCount++
	return cal, nil
}

type mockRacingClient struct {
	meetings []*racing.Meeting
	races    []*racing.Race

	updateRaceCount int
}

func (c *mockRacingClient) AddMeetings(ctx context.Context, req *racing.AddMeetingsRequest, opts ...client.CallOption) (*racing.AddMeetingsResponse, error) {
	return nil, nil
}

func (c *mockRacingClient) AddRaces(ctx context.Context, req *racing.AddRacesRequest, opts ...client.CallOption) (*racing.AddRacesResponse, error) {
	return nil, nil
}

func (c *mockRacingClient) ListMeetingsByDate(ctx context.Context, req *racing.ListMeetingsByDateRequest, opts ...client.CallOption) (*racing.ListMeetingsByDateResponse, error) {
	resp := &racing.ListMeetingsByDateResponse{
		Meetings: c.meetings,
	}
	return resp, nil
}

func (c *mockRacingClient) ListRacesByMeetingDate(ctx context.Context, req *racing.ListRacesByMeetingDateRequest, opts ...client.CallOption) (*racing.ListRacesByMeetingDateResponse, error) {
	resp := &racing.ListRacesByMeetingDateResponse{
		Races: c.races,
	}
	return resp, nil
}

func (c *mockRacingClient) UpdateRace(ctx context.Context, req *racing.UpdateRaceRequest, opts ...client.CallOption) (*racing.UpdateRaceResponse, error) {
	c.updateRaceCount++
	return nil, nil
}

func (c *mockRacingClient) GetNextRace(ctx context.Context, req *racing.GetNextRaceRequest, opts ...client.CallOption) (*racing.GetNextRaceResponse, error) {
	return nil, nil
}
