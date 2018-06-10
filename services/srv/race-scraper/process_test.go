package main

import (
	"context"
	"fmt"
	racing "github.com/krozlink/oddzy/services/srv/racing/proto"
	client "github.com/micro/go-micro/client"
	"testing"
)

func TestReadInternalReturnsData(t *testing.T) {

	meetings := []*racing.Meeting{
		getTestMeeting("m123", 2),
		getTestMeeting("m456", 1),
	}
	races := []*racing.Race{
		getTestRace(meetings[0], 1),
		getTestRace(meetings[0], 2),
		getTestRace(meetings[1], 1),
	}

	client := &mockRacingClient{
		meetings: meetings,
		races:    races,
	}
	p := getTestProcess(client)

	intMeetings, intRaces, err := readInternal(p)
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
	t.Fail()
}

func TestUpdateExistingRacesHandlesMultipleRaces(t *testing.T) {
	t.Fail()
}

func TestProcessRaceCalendarSplitsNewAndExisting(t *testing.T) {
	t.Fail()
}

func TestGetMeetingSourceID(t *testing.T) {
	t.Fail()
}

func TestGetRaceStatusFromCalendar(t *testing.T) {
	t.Fail()
}

func TestParseRaceCardReadsSelections(t *testing.T) {
	t.Fail()
}

func TestScrapeRacesScrapesAllMissingRaces(t *testing.T) {
	t.Fail()
}

func TestScrapeRacesUpdatesScrapedRaces(t *testing.T) {
	t.Fail()
}

func getTestProcess(c *mockRacingClient) *scrapeProcess {
	return &scrapeProcess{
		status:  "TEST",
		done:    make(chan bool),
		http:    handler{},
		racing:  c,
		scraper: &mockScraper{},
	}
}

func getTestMeeting(meetingID string, numRaces int) *racing.Meeting {
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
		SourceId:       "source-" + meetingID,
	}
}

func getTestRace(m *racing.Meeting, number int32) *racing.Race {
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
		SourceId:       m.SourceId + "-1",
		Status:         status,
	}
}

type mockScraper struct {
}

func (c *mockScraper) ScrapeRaceCard(sourceID string) (*RaceCard, error) {
	return nil, nil
}

func (c *mockScraper) ScrapeRaceCalendar(eventType string, date string) (*RaceCalendar, error) {
	return nil, nil
}

type mockRacingClient struct {
	meetings []*racing.Meeting
	races    []*racing.Race
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
	return nil, nil
}

func (c *mockRacingClient) GetNextRace(ctx context.Context, req *racing.GetNextRaceRequest, opts ...client.CallOption) (*racing.GetNextRaceResponse, error) {
	return nil, nil
}
