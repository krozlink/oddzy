package main

import (
	"context"
	"encoding/json"
	"fmt"
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	"github.com/micro/go-micro/broker"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

type MockRepo struct {
	meetings   []*proto.Meeting
	races      []*proto.Race
	selections []*proto.Selection
}

func getTestRacingService(repo *MockRepo) *RacingService {
	b := &MockBroker{}
	return NewRacingService(repo, b)
}

func TestListRacesByMeetingDate(t *testing.T) {
	stats = getMockStats()

	repo := &MockRepo{
		races: []*proto.Race{
			&proto.Race{
				RaceId:       "111",
				MeetingStart: 1000,
				Name:         "Race 1",
			},
			&proto.Race{
				RaceId:       "222",
				MeetingStart: 2000,
				Name:         "Race 2",
			},
		},
	}

	ctx := context.Background()
	srv := getTestRacingService(repo)

	var tests = []struct {
		start    int64
		end      int64
		expected int
	}{
		{999, 1001, 1},
		{1000, 1000, 1},
		{999, 999, 0},
		{1000, 2000, 2},
		{2000, 3000, 1},
	}

	for _, v := range tests {
		req := &proto.ListRacesByMeetingDateRequest{
			StartDate: v.start,
			EndDate:   v.end,
		}

		resp := &proto.ListRacesByMeetingDateResponse{}
		err := srv.ListRacesByMeetingDate(ctx, req, resp)

		assert.NoError(t, err)
		assert.Equal(t, v.expected, len(resp.Races))
	}
}

func TestListRacesByMeetingDateValidation(t *testing.T) {
	stats = getMockStats()
	repo := &MockRepo{
		races: []*proto.Race{
			&proto.Race{
				RaceId:       "111",
				MeetingStart: 1000,
				Name:         "Race 1",
			},
			&proto.Race{
				RaceId:       "222",
				MeetingStart: 2000,
				Name:         "Race 2",
			},
		},
	}

	ctx := context.Background()
	srv := getTestRacingService(repo)

	var tests = []struct {
		start int64
		end   int64
		error bool
	}{
		{0, 0, true},
		{1, 1, false},
		{0, 1000, true},
		{1000, 0, true},
	}

	for _, v := range tests {
		req := &proto.ListRacesByMeetingDateRequest{
			StartDate: v.start,
			EndDate:   v.end,
		}

		resp := &proto.ListRacesByMeetingDateResponse{}
		err := srv.ListRacesByMeetingDate(ctx, req, resp)

		assert.Falsef(t, err == nil && v.error, "Expected error when calling ListRacesByMeetingDate for range %v to %v", v.start, v.end)
		assert.Falsef(t, err != nil && !v.error, "Expected no error when calling ListRacesByMeetingDate for range %v to %v", v.start, v.end)
	}
}

func TestListMeetingsByDate(t *testing.T) {
	stats = getMockStats()
	repo := &MockRepo{
		meetings: []*proto.Meeting{
			&proto.Meeting{
				MeetingId:      "Albion",
				Country:        "AUS",
				ScheduledStart: 1000,
			},
			&proto.Meeting{
				MeetingId:      "Ascot",
				Country:        "AUS",
				ScheduledStart: 2000,
			},
		},
	}

	ctx := context.Background()
	srv := getTestRacingService(repo)

	var tests = []struct {
		start    int64
		end      int64
		expected int
	}{
		{999, 1001, 1},
		{1000, 1000, 1},
		{999, 999, 0},
		{1000, 2000, 2},
		{2000, 3000, 1},
	}

	for _, v := range tests {
		req := &proto.ListMeetingsByDateRequest{
			StartDate: v.start,
			EndDate:   v.end,
		}

		resp := &proto.ListMeetingsByDateResponse{}
		err := srv.ListMeetingsByDate(ctx, req, resp)

		assert.NoError(t, err)
		assert.Equal(t, v.expected, len(resp.Meetings))
	}
}

func TestListMeetingsByDateValidation(t *testing.T) {
	stats = getMockStats()
	repo := &MockRepo{
		meetings: []*proto.Meeting{
			&proto.Meeting{
				MeetingId:      "Albion",
				Country:        "AUS",
				ScheduledStart: 1000,
			},
			&proto.Meeting{
				MeetingId:      "Ascot",
				Country:        "AUS",
				ScheduledStart: 2000,
			},
		},
	}

	ctx := context.Background()
	srv := getTestRacingService(repo)

	var tests = []struct {
		start int64
		end   int64
		error bool
	}{
		{0, 0, true},
		{1, 1, false},
		{0, 1000, true},
		{1000, 0, true},
	}

	for _, v := range tests {
		req := &proto.ListMeetingsByDateRequest{
			StartDate: v.start,
			EndDate:   v.end,
		}

		resp := &proto.ListMeetingsByDateResponse{}
		err := srv.ListMeetingsByDate(ctx, req, resp)

		assert.Falsef(t, err == nil && v.error, "Expected error when calling ListMeetingsByDate for range %v to %v", v.start, v.end)
		assert.Falsef(t, err != nil && !v.error, "Expected no error when calling ListMeetingsByDate for range %v to %v", v.start, v.end)
	}
}

func TestAddMeetingsValidation(t *testing.T) {

	stats = getMockStats()
	repo := &MockRepo{}
	ctx := context.Background()
	srv := getTestRacingService(repo)

	var testValues = []struct {
		meetings    []*proto.Meeting
		expectedErr string
	}{
		{
			meetings: []*proto.Meeting{&proto.Meeting{
				MeetingId:      "meeting-1",
				SourceId:       "source-123",
				RaceIds:        []string{"race1", "race2"},
				Country:        "AUS",
				Name:           "Meeting 1",
				ScheduledStart: 1000,
				RaceType:       "harness",
			}},
			expectedErr: "",
		},
		{
			meetings: []*proto.Meeting{&proto.Meeting{
				SourceId:       "source-123",
				RaceIds:        []string{"race1", "race2"},
				Country:        "AUS",
				Name:           "Meeting 1",
				ScheduledStart: 1000,
				RaceType:       "harness",
			}},
			expectedErr: "Meeting id not provided",
		},
		{
			meetings: []*proto.Meeting{&proto.Meeting{
				MeetingId:      "meeting-1",
				RaceIds:        []string{"race1", "race2"},
				Country:        "AUS",
				Name:           "Meeting 1",
				ScheduledStart: 1000,
				RaceType:       "harness",
			}},
			expectedErr: "Source id not provided",
		},
		{
			meetings: []*proto.Meeting{&proto.Meeting{
				MeetingId:      "meeting-1",
				SourceId:       "source-123",
				Country:        "AUS",
				Name:           "Meeting 1",
				ScheduledStart: 1000,
				RaceType:       "harness",
			}},
			expectedErr: "No race ids provided",
		},
		{
			meetings: []*proto.Meeting{&proto.Meeting{
				MeetingId: "meeting-1",
				SourceId:  "source-123",
				RaceIds:   []string{"race1", "race2"},
				Country:   "AUS",
				Name:      "Meeting 1",
				RaceType:  "harness",
			}},
			expectedErr: "No scheduled start date provided",
		},
		{
			meetings: []*proto.Meeting{&proto.Meeting{
				MeetingId:      "meeting-1",
				SourceId:       "source-123",
				RaceIds:        []string{"race1", "race2"},
				Country:        "AUS",
				Name:           "Meeting 1",
				ScheduledStart: 1000,
			}},
			expectedErr: "No race type provided",
		},
	}

	for _, v := range testValues {
		req := &proto.AddMeetingsRequest{
			Meetings: v.meetings,
		}

		resp := &proto.AddMeetingsResponse{}

		err := srv.AddMeetings(ctx, req, resp)
		if err != nil && v.expectedErr == "" {
			t.Errorf("Expected no error but got %v", err)
		} else if err != nil {
			if !strings.Contains(err.Error(), v.expectedErr) {
				t.Errorf("Expected '%v' but got '%v'", v.expectedErr, err)
			}
		} else if err == nil && v.expectedErr != "" {
			t.Errorf("Did not get expected error '%v'", v.expectedErr)
		}
	}
}

func TestAddMeetings(t *testing.T) {
	stats = getMockStats()
	repo := &MockRepo{}
	ctx := context.Background()
	srv := getTestRacingService(repo)

	var meetings = []*proto.Meeting{
		&proto.Meeting{
			MeetingId:      "meeting-1",
			SourceId:       "source-123",
			RaceIds:        []string{"race1", "race2"},
			Country:        "AUS",
			Name:           "Meeting 1",
			ScheduledStart: 1000,
			RaceType:       "harness",
		},
		&proto.Meeting{
			MeetingId:      "meeting-2",
			SourceId:       "source-abc",
			RaceIds:        []string{"race3", "race4"},
			Country:        "AUS",
			Name:           "Meeting 2",
			ScheduledStart: 2000,
			RaceType:       "harness",
		},
	}

	req := &proto.AddMeetingsRequest{
		Meetings: meetings,
	}

	resp := &proto.AddMeetingsResponse{}
	err := srv.AddMeetings(ctx, req, resp)

	if err != nil {
		t.Error(err)
	}

	if len(repo.meetings) != 2 {
		t.Errorf("Expected 2 meetings but only got %v", len(repo.meetings))
	}

	if repo.meetings[0].MeetingId != "meeting-1" {
		t.Errorf("Expected meeting id %v but got %v", "meeting-1", repo.meetings[0].MeetingId)
	}

	if repo.meetings[1].MeetingId != "meeting-2" {
		t.Errorf("Expected meeting id %v but got %v", "meeting-2", repo.meetings[1].MeetingId)
	}
}

func TestAddRacesValidation(t *testing.T) {
	stats = getMockStats()
	repo := &MockRepo{}
	ctx := context.Background()
	srv := getTestRacingService(repo)

	var testValues = []struct {
		races       []*proto.Race
		expectedErr string
	}{
		{
			races: []*proto.Race{&proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         1,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			}},
			expectedErr: "",
		},
		{
			races: []*proto.Race{&proto.Race{
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         1,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			}},
			expectedErr: "Race id not provided",
		},
		{
			races: []*proto.Race{&proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         1,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				Status:         "AWESOME",
			}},
			expectedErr: "Source id not provided",
		},
		{
			races: []*proto.Race{&proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Number:         1,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			}},
			expectedErr: "No name provided",
		},
		{
			races: []*proto.Race{&proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			}},
			expectedErr: "No number provided",
		},
		{
			races: []*proto.Race{&proto.Race{
				RaceId:       "race-1",
				ActualStart:  1000,
				DateCreated:  2000,
				LastUpdated:  3000,
				MeetingId:    "meeting-1",
				MeetingStart: 4000,
				Name:         "Race 1",
				Number:       1,
				Results:      "2,3,4",
				SourceId:     "source-1",
				Status:       "AWESOME",
			}},
			expectedErr: "No scheduled start time provided",
		},
		{
			races: []*proto.Race{&proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         1,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
			}},
			expectedErr: "No status provided",
		},
		{
			races: []*proto.Race{&proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         1,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			}},
			expectedErr: "No last updated time provided",
		},
		{
			races: []*proto.Race{&proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				Name:           "Race 1",
				Number:         1,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			}},
			expectedErr: "No meeting start time provided",
		},
	}

	for _, v := range testValues {
		req := &proto.AddRacesRequest{
			Races: v.races,
		}

		resp := &proto.AddRacesResponse{}

		err := srv.AddRaces(ctx, req, resp)
		if err != nil && v.expectedErr == "" {
			t.Errorf("Expected no error but got %v", err)
		} else if err != nil {
			if !strings.Contains(err.Error(), v.expectedErr) {
				t.Errorf("Expected '%v' but got '%v'", v.expectedErr, err)
			}
		} else if err == nil && v.expectedErr != "" {
			t.Errorf("Did not get expected error '%v'", v.expectedErr)
		}
	}
}

func TestAddRaces(t *testing.T) {
	stats = getMockStats()
	repo := &MockRepo{}
	ctx := context.Background()
	srv := getTestRacingService(repo)

	var races = []*proto.Race{
		&proto.Race{
			RaceId:         "race-1",
			ActualStart:    1000,
			DateCreated:    2000,
			LastUpdated:    3000,
			MeetingId:      "meeting-1",
			MeetingStart:   4000,
			Name:           "Race 1",
			Number:         1,
			Results:        "2,3,4",
			ScheduledStart: 5000,
			SourceId:       "source-1",
			Status:         "AWESOME",
		},
		&proto.Race{
			RaceId:         "race-2",
			ActualStart:    1000,
			DateCreated:    2000,
			LastUpdated:    3000,
			MeetingId:      "meeting-2",
			MeetingStart:   4000,
			Name:           "Race 2",
			Number:         2,
			Results:        "2,3,4",
			ScheduledStart: 5000,
			SourceId:       "source-2",
			Status:         "AWESOME",
		},
	}

	req := &proto.AddRacesRequest{
		Races: races,
	}

	resp := &proto.AddRacesResponse{}
	err := srv.AddRaces(ctx, req, resp)

	if err != nil {
		t.Error(err)
	}

	if len(repo.races) != len(races) {
		t.Errorf("Expected %v races but only got %v", len(races), len(repo.meetings))
	}

	if repo.races[0].RaceId != "race-1" {
		t.Errorf("Expected race id %v but got %v", "race-1", repo.races[0].RaceId)
	}

	if repo.races[1].RaceId != "race-2" {
		t.Errorf("Expected race id %v but got %v", "race-2", repo.races[1].RaceId)
	}
}

func TestUpdateRaceValidatesRace(t *testing.T) {
	stats = getMockStats()
	repo := &MockRepo{
		races: []*proto.Race{
			&proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			},
		},
	}
	ctx := context.Background()
	srv := getTestRacingService(repo)

	var testValues = []struct {
		race        *proto.Race
		expectedErr string
	}{
		{
			race: &proto.Race{
				RaceId:         "race-1",
				ActualStart:    1001,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			},
			expectedErr: "",
		},
		{
			race: &proto.Race{
				ActualStart:    1001,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			},
			expectedErr: "Race id not provided",
		},
		{
			race: &proto.Race{
				RaceId:         "race-1",
				ActualStart:    1001,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				Status:         "AWESOME",
			},
			expectedErr: "Source id not provided for the race",
		},
		{
			race: &proto.Race{
				RaceId:       "race-1",
				ActualStart:  1001,
				DateCreated:  2000,
				LastUpdated:  3000,
				MeetingId:    "meeting-1",
				MeetingStart: 4000,
				Name:         "Race 1",
				Number:       2,
				Results:      "2,3,4",
				SourceId:     "source-1",
				Status:       "AWESOME",
			},
			expectedErr: "Scheduled start time not provided for the race",
		},
		{
			race: &proto.Race{
				RaceId:         "race-1",
				ActualStart:    1001,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			},
			expectedErr: "Number not provided for the race",
		},
		{
			race: &proto.Race{
				RaceId:         "race-1",
				ActualStart:    1001,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			},
			expectedErr: "Name not provided for the race",
		},
		{
			race: &proto.Race{
				RaceId:         "race-1",
				ActualStart:    1001,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
			},
			expectedErr: "Status not provided for the race",
		},
	}

	for _, v := range testValues {
		req := &proto.UpdateRaceRequest{
			Race: v.race,
			Selections: []*proto.Selection{
				&proto.Selection{
					SelectionId:        "selection-a",
					BarrierNumber:      2,
					CompetitorId:       "competitor-1",
					Jockey:             "Daniel",
					Name:               "Winx",
					Number:             3,
					RaceId:             "race-1",
					SourceCompetitorId: "source-comp-1",
					SourceId:           "source-a",
				},
			},
		}

		resp := &proto.UpdateRaceResponse{}
		err := srv.UpdateRace(ctx, req, resp)
		if err != nil && v.expectedErr == "" {
			t.Errorf("Expected no error but got %v", err)
		} else if err != nil {
			if !strings.Contains(err.Error(), v.expectedErr) {
				t.Errorf("Expected '%v' but got '%v'", v.expectedErr, err)
			}
		} else if err == nil && v.expectedErr != "" {
			t.Errorf("Did not get expected error '%v'", v.expectedErr)
		}
	}
}

func TestUpdateRaceValidatesSelections(t *testing.T) {

	stats = getMockStats()
	repo := &MockRepo{
		races: []*proto.Race{
			&proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			},
		},
	}
	ctx := context.Background()
	srv := getTestRacingService(repo)

	var testValues = []struct {
		selection   *proto.Selection
		expectedErr string
	}{
		{
			selection: &proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-1",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			expectedErr: "",
		},
		{
			selection: &proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-1",
				SourceCompetitorId: "source-comp-1",
			},
			expectedErr: "Source id not provided",
		},
		{
			selection: &proto.Selection{
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-1",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			expectedErr: "Selection id not provided",
		},
		{
			selection: &proto.Selection{
				SelectionId:   "selection-a",
				BarrierNumber: 2,
				CompetitorId:  "competitor-1",
				Jockey:        "Daniel",
				Name:          "Winx",
				Number:        3,
				RaceId:        "race-1",
				SourceId:      "source-a",
			},
			expectedErr: "Source competitor id not provided",
		},
		{
			selection: &proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			expectedErr: "Race id not provided for selection",
		},
		{
			selection: &proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			expectedErr: "does not match the race",
		},
		{
			selection: &proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Number:             3,
				RaceId:             "race-1",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			expectedErr: "Name not provided for selection",
		},
		{
			selection: &proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-1",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			expectedErr: "Jockey not provided for selection",
		},
		{
			selection: &proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				RaceId:             "race-1",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			expectedErr: "Number not provided for selection",
		},
		{
			selection: &proto.Selection{
				SelectionId:        "selection-a",
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-1",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			expectedErr: "Barrier number not provided for selection",
		},
	}

	for _, v := range testValues {

		req := &proto.UpdateRaceRequest{
			Race: &proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         1,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			},
			Selections: []*proto.Selection{v.selection},
		}

		resp := &proto.UpdateRaceResponse{}
		err := srv.UpdateRace(ctx, req, resp)
		if err != nil && v.expectedErr == "" {
			t.Errorf("Expected no error but got %v", err)
		} else if err != nil {
			if !strings.Contains(err.Error(), v.expectedErr) {
				t.Errorf("Expected '%v' but got '%v'", v.expectedErr, err)
			}
		} else if err == nil && v.expectedErr != "" {
			t.Errorf("Did not get expected error '%v'", v.expectedErr)
		}
	}
}

func TestUpdateRaceCreatesSelectionsOnInitalCall(t *testing.T) {

	stats = getMockStats()
	// Create two races
	// Race 1 has no selections
	// Race 2 has 2 selections
	repo := &MockRepo{
		races: []*proto.Race{
			&proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         1,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			},
			&proto.Race{
				RaceId:         "race-2",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-2",
				Status:         "AWESOME",
			},
		},
		selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			&proto.Selection{
				SelectionId:        "selection-b",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-b",
			},
		},
	}
	ctx := context.Background()
	srv := getTestRacingService(repo)

	// Update Race 1 with 2 selections
	req := &proto.UpdateRaceRequest{
		Race: &proto.Race{
			RaceId:         "race-1",
			ActualStart:    1000,
			DateCreated:    2000,
			LastUpdated:    3000,
			MeetingId:      "meeting-1",
			MeetingStart:   4000,
			Name:           "Race 1",
			Number:         1,
			Results:        "2,3,4",
			ScheduledStart: 5000,
			SourceId:       "source-1",
			Status:         "AWESOME",
		},
		Selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-1",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-1",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-1",
			},
			&proto.Selection{
				SelectionId:        "selection-2",
				BarrierNumber:      3,
				CompetitorId:       "competitor-2",
				Jockey:             "Dan",
				Name:               "Hartnell",
				Number:             4,
				RaceId:             "race-1",
				SourceCompetitorId: "source-comp-2",
				SourceId:           "source-2",
			},
			&proto.Selection{
				SelectionId:        "selection-3",
				BarrierNumber:      4,
				CompetitorId:       "competitor-3",
				Jockey:             "DT",
				Name:               "Humidor",
				Number:             5,
				RaceId:             "race-1",
				SourceCompetitorId: "source-comp-3",
				SourceId:           "source-3",
			},
		},
	}

	resp := &proto.UpdateRaceResponse{}

	err := srv.UpdateRace(ctx, req, resp)

	if err != nil {
		t.Error(err)
	}

	//expect repository to now contain the 2 original selections and the 2 new ones - 4 in total
	expected := 2 + len(req.Selections)
	if expected != len(repo.selections) {
		t.Errorf("Expected %v selections but only %v exist", expected, len(repo.selections))
	}
}

func TestUpdateRaceFailsWhenNumberOfSelectionsIncreases(t *testing.T) {
	stats = getMockStats()
	// Create two races
	// Race 1 has no selections
	// Race 2 has 2 selections
	repo := &MockRepo{
		races: []*proto.Race{
			&proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         1,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			},
			&proto.Race{
				RaceId:         "race-2",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-2",
				Status:         "AWESOME",
			},
		},
		selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			&proto.Selection{
				SelectionId:        "selection-b",
				BarrierNumber:      4,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Hartnell",
				Number:             5,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-b",
			},
		},
	}
	ctx := context.Background()
	srv := getTestRacingService(repo)

	// Update 3 selections in Race 2
	req := &proto.UpdateRaceRequest{
		Race: &proto.Race{
			RaceId:         "race-2",
			ActualStart:    1000,
			DateCreated:    2000,
			LastUpdated:    3000,
			MeetingId:      "meeting-1",
			MeetingStart:   4000,
			Name:           "Race 2",
			Number:         1,
			Results:        "2,3,4",
			ScheduledStart: 5000,
			SourceId:       "source-1",
			Status:         "AWESOME",
		},
		Selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      22,
				CompetitorId:       "competitor-winx",
				Jockey:             "Hugh",
				Name:               "Winx Updated",
				Number:             33,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-winx",
				SourceId:           "source-a",
			},
			&proto.Selection{
				SelectionId:        "selection-b",
				BarrierNumber:      44,
				CompetitorId:       "competitor-hartnell",
				Jockey:             "Daniel",
				Name:               "Hartnell Updated",
				Number:             55,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-hartnell",
				SourceId:           "source-b",
			},
			&proto.Selection{
				SelectionId:        "selection-c",
				BarrierNumber:      66,
				CompetitorId:       "competitor-humidor",
				Jockey:             "Daniel",
				Name:               "Humidor Updated",
				Number:             77,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-humidor",
				SourceId:           "source-c",
			},
		},
	}

	resp := &proto.UpdateRaceResponse{}

	err := srv.UpdateRace(ctx, req, resp)

	if err == nil {
		t.Error("Expected error for number of selections changing from 2 to 3")
	} else if !strings.Contains(err.Error(), "Number of selections unexpectedly changed") {
		t.Error("Expected error for number of selections changing from 2 to 3")
	}
}

func TestUpdateRaceFlagsScratchedWhenSelectionIsRemoved(t *testing.T) {

	var originalLastUpdate int64 = 2000
	stats = getMockStats()
	// Create two races
	// Race 1 has no selections
	// Race 2 has 2 selections
	repo := &MockRepo{
		races: []*proto.Race{
			&proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    originalLastUpdate,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         1,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			},
			&proto.Race{
				RaceId:         "race-2",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    originalLastUpdate,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-2",
				Status:         "AWESOME",
			},
		},
		selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
				LastUpdated:        originalLastUpdate,
			},
			&proto.Selection{
				SelectionId:        "selection-b",
				BarrierNumber:      4,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Hartnell",
				Number:             5,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-b",
				LastUpdated:        originalLastUpdate,
			},
		},
	}
	ctx := context.Background()
	srv := getTestRacingService(repo)

	// Update 1 selection in Race 2 (no change, but 2nd selection is no longer included)
	req := &proto.UpdateRaceRequest{
		Race: &proto.Race{
			RaceId:         "race-2",
			ActualStart:    1000,
			DateCreated:    2000,
			LastUpdated:    3000,
			MeetingId:      "meeting-1",
			MeetingStart:   4000,
			Name:           "Race 2",
			Number:         1,
			Results:        "2,3,4",
			ScheduledStart: 5000,
			SourceId:       "source-1",
			Status:         "AWESOME",
		},
		Selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      3,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
		},
	}

	resp := &proto.UpdateRaceResponse{}

	err := srv.UpdateRace(ctx, req, resp)
	assert.NoError(t, err)

	// Selection that was not updated is now flagged as scratched
	assert.True(t, repo.selections[1].Scratched)

	// Selection that was updated is not flagged as scratched
	assert.False(t, repo.selections[0].Scratched)

	// Confirm both selections were modified
	assert.NotEqual(t, originalLastUpdate, repo.selections[0].LastUpdated)
	assert.NotEqual(t, originalLastUpdate, repo.selections[1].LastUpdated)
}

func TestUpdateRaceModifiesExistingSelections(t *testing.T) {

	stats = getMockStats()
	// Create two races
	// Race 1 has no selections
	// Race 2 has 2 selections
	repo := &MockRepo{
		races: []*proto.Race{
			&proto.Race{
				RaceId:         "race-1",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         1,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-1",
				Status:         "AWESOME",
			},
			&proto.Race{
				RaceId:         "race-2",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-2",
				Status:         "AWESOME",
			},
		},
		selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			&proto.Selection{
				SelectionId:        "selection-b",
				BarrierNumber:      4,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Hartnell",
				Number:             5,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-b",
			},
		},
	}
	ctx := context.Background()
	srv := getTestRacingService(repo)

	// Update the two selections in Race 2
	req := &proto.UpdateRaceRequest{
		Race: &proto.Race{
			RaceId:         "race-2",
			ActualStart:    1000,
			DateCreated:    2000,
			LastUpdated:    3000,
			MeetingId:      "meeting-1",
			MeetingStart:   4000,
			Name:           "Race 2",
			Number:         1,
			Results:        "2,3,4",
			ScheduledStart: 5000,
			SourceId:       "source-1",
			Status:         "AWESOME",
		},
		Selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      22,
				CompetitorId:       "competitor-winx",
				Jockey:             "Hugh",
				Name:               "Winx Updated",
				Number:             33,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-winx",
				SourceId:           "source-a",
			},
			&proto.Selection{
				SelectionId:        "selection-b",
				BarrierNumber:      44,
				CompetitorId:       "competitor-hartnell",
				Jockey:             "Daniel",
				Name:               "Hartnell Updated",
				Number:             55,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-hartnell",
				SourceId:           "source-b",
			},
		},
	}

	resp := &proto.UpdateRaceResponse{}

	err := srv.UpdateRace(ctx, req, resp)

	if err != nil {
		t.Error(err)
	}

	//expect repository to still contain the 2 original selections but now with updated details
	if len(req.Selections) != len(repo.selections) {
		t.Errorf("Expected %v selections but only %v exist", len(req.Selections), len(repo.selections))
	}

	if repo.selections[0].Name != req.Selections[0].Name {
		t.Errorf("Expected selection name to be %v but got %v", req.Selections[0].Name, repo.selections[0].Name)
	}

	if repo.selections[0].BarrierNumber != req.Selections[0].BarrierNumber {
		t.Errorf("Expected barrier number to be %v but got %v", req.Selections[0].BarrierNumber, repo.selections[0].BarrierNumber)
	}

	if repo.selections[0].Number != req.Selections[0].Number {
		t.Errorf("Expected number to be %v but got %v", req.Selections[0].Number, repo.selections[0].Number)
	}

}

func TestUpdateRaceNotifiesOnRaceChange(t *testing.T) {
	stats = getMockStats()
	// Create 1 race with 2 selections
	repo := &MockRepo{
		races: []*proto.Race{
			&proto.Race{
				RaceId:         "race-2",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-2",
				Status:         "AWESOME",
			},
		},
		selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			&proto.Selection{
				SelectionId:        "selection-b",
				BarrierNumber:      4,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Hartnell",
				Number:             5,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-b",
			},
		},
	}
	ctx := context.Background()
	broker := &MockBroker{}
	srv := NewRacingService(repo, broker)

	// Update the race but not the selections
	req := &proto.UpdateRaceRequest{
		Race: &proto.Race{
			RaceId:         "race-2",
			ActualStart:    1000,
			DateCreated:    2000,
			LastUpdated:    3000,
			MeetingId:      "meeting-1",
			MeetingStart:   4000,
			Name:           "Race 1",
			Number:         2,
			Results:        "2,3,4",
			ScheduledStart: 6000, // Scheduled start changed from 5000 to 6000
			SourceId:       "source-2",
			Status:         "AWESOME",
		},
		Selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			&proto.Selection{
				SelectionId:        "selection-b",
				BarrierNumber:      4,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Hartnell",
				Number:             5,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-b",
			},
		},
	}

	resp := &proto.UpdateRaceResponse{}

	if err := srv.UpdateRace(ctx, req, resp); err != nil {
		t.Error(err)
	}

	msg := &proto.RaceUpdatedMessage{}
	err := json.Unmarshal(broker.messages[0].Body, msg)
	if err != nil {
		t.Error(err)
	}

	if len(broker.messages) != 1 {
		t.Errorf("Expected 1 message but got %v", len(broker.messages))
	}

	if broker.messages[0].Header["race_id"] != "race-2" {
		t.Errorf("Expected message with race id %v but got %v", "race-2", broker.messages[0].Header["race_id"])
	}

	if broker.messages[0].Header["meeting_id"] != "meeting-1" {
		t.Errorf("Expected message with meeting id %v but got %v", "meeting-1", broker.messages[0].Header["meeting_id"])
	}

	if msg.Race.SourceId != repo.races[0].SourceId {
		t.Errorf("Expected message with source id %v but got %v", repo.races[0].SourceId, msg.Race.SourceId)
	}

	if msg.Selections[0].SelectionId != repo.selections[0].SelectionId {
		t.Errorf("Expected message with selection id %v but got %v", repo.selections[0].SelectionId, msg.Selections[0].SelectionId)
	}
}

func TestUpdateRaceNotifiesOnSelectionChange(t *testing.T) {
	stats = getMockStats()
	// Create 1 race with 2 selections
	repo := &MockRepo{
		races: []*proto.Race{
			&proto.Race{
				RaceId:         "race-2",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-2",
				Status:         "AWESOME",
			},
		},
		selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			&proto.Selection{
				SelectionId:        "selection-b",
				BarrierNumber:      4,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Hartnell",
				Number:             5,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-b",
			},
		},
	}
	ctx := context.Background()
	broker := &MockBroker{}
	srv := NewRacingService(repo, broker)

	// Update the selections but not the race
	req := &proto.UpdateRaceRequest{
		Race: &proto.Race{
			RaceId:         "race-2",
			ActualStart:    1000,
			DateCreated:    2000,
			LastUpdated:    3000,
			MeetingId:      "meeting-1",
			MeetingStart:   4000,
			Name:           "Race 1",
			Number:         2,
			Results:        "2,3,4",
			ScheduledStart: 5000,
			SourceId:       "source-2",
			Status:         "AWESOME",
		},
		Selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      22,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			&proto.Selection{
				SelectionId:        "selection-b",
				BarrierNumber:      44,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Hartnell",
				Number:             5,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-b",
			},
		},
	}

	resp := &proto.UpdateRaceResponse{}

	if err := srv.UpdateRace(ctx, req, resp); err != nil {
		t.Error(err)
	}

	if len(broker.messages) != 1 {
		t.Fatalf("Expected 1 message but got %v", len(broker.messages))
	}

	msg := &proto.RaceUpdatedMessage{}
	err := json.Unmarshal(broker.messages[0].Body, msg)
	if err != nil {
		t.Error(err)
	}

	if msg.Selections[0].BarrierNumber != 22 {
		t.Errorf("Expected message with selection name %v but got %v", 22, msg.Selections[0].Name)
	}
}

func TestUpdateRaceNoNotificationIfNoChange(t *testing.T) {
	stats = getMockStats()
	// Create 1 race with 2 selections
	repo := &MockRepo{
		races: []*proto.Race{
			&proto.Race{
				RaceId:         "race-2",
				ActualStart:    1000,
				DateCreated:    2000,
				LastUpdated:    3000,
				MeetingId:      "meeting-1",
				MeetingStart:   4000,
				Name:           "Race 1",
				Number:         2,
				Results:        "2,3,4",
				ScheduledStart: 5000,
				SourceId:       "source-2",
				Status:         "AWESOME",
			},
		},
		selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			&proto.Selection{
				SelectionId:        "selection-b",
				BarrierNumber:      4,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Hartnell",
				Number:             5,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-b",
			},
		},
	}
	ctx := context.Background()
	broker := &MockBroker{}
	srv := NewRacingService(repo, broker)

	// Update the selections but not the race
	req := &proto.UpdateRaceRequest{
		Race: &proto.Race{
			RaceId:         "race-2",
			ActualStart:    1000,
			DateCreated:    2000,
			LastUpdated:    3000,
			MeetingId:      "meeting-1",
			MeetingStart:   4000,
			Name:           "Race 1",
			Number:         2,
			Results:        "2,3,4",
			ScheduledStart: 5000,
			SourceId:       "source-2",
			Status:         "AWESOME",
		},
		Selections: []*proto.Selection{
			&proto.Selection{
				SelectionId:        "selection-a",
				BarrierNumber:      2,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Winx",
				Number:             3,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-a",
			},
			&proto.Selection{
				SelectionId:        "selection-b",
				BarrierNumber:      4,
				CompetitorId:       "competitor-1",
				Jockey:             "Daniel",
				Name:               "Hartnell",
				Number:             5,
				RaceId:             "race-2",
				SourceCompetitorId: "source-comp-1",
				SourceId:           "source-b",
			},
		},
	}

	resp := &proto.UpdateRaceResponse{}

	if err := srv.UpdateRace(ctx, req, resp); err != nil {
		t.Error(err)
	}

	if len(broker.messages) != 0 {
		t.Fatalf("Expected no message but got %v", len(broker.messages))
	}
}

func TestGetNextRace(t *testing.T) {

	stats = getMockStats()
	repo := &MockRepo{}

	ctx := context.Background()
	srv := getTestRacingService(repo)

	var testValues = []struct {
		races    []*proto.Race
		expected int32
	}{
		{ // 2 races that already started
			races: []*proto.Race{
				getTestRace("meeting-1", 1, 1000),
				getTestRace("meeting-1", 2, 2000),
			},
			expected: 0,
		},
		{ // 1 race that has started
			races: []*proto.Race{
				getTestRace("meeting-1", 1, 1000),
				getTestRace("meeting-1", 2, 0),
			},
			expected: 2,
		},
		{ // 2 races that are yet to start
			races: []*proto.Race{
				getTestRace("meeting-1", 1, 0),
				getTestRace("meeting-1", 2, 0),
			},
			expected: 1,
		},
		{ // 2 races that already started - reverse order
			races: []*proto.Race{
				getTestRace("meeting-1", 2, 2000),
				getTestRace("meeting-1", 1, 1000),
			},
			expected: 0,
		},
		{ // 1 race that has started - reverse order
			races: []*proto.Race{
				getTestRace("meeting-1", 2, 0),
				getTestRace("meeting-1", 1, 1000),
			},
			expected: 2,
		},
		{ // 2 races that are yet to start - reverse order
			races: []*proto.Race{
				getTestRace("meeting-1", 2, 0),
				getTestRace("meeting-1", 1, 0),
			},
			expected: 1,
		},
	}

	for _, v := range testValues {

		repo.races = v.races

		req := &proto.GetNextRaceRequest{
			MeetingId: "meeting-1",
		}

		resp := &proto.GetNextRaceResponse{}
		err := srv.GetNextRace(ctx, req, resp)
		if err != nil {
			t.Error(err)
		}

		result := int32(0)
		if resp.Race != nil {
			result = resp.Race.Number
		}

		if v.expected != result {
			t.Errorf("Expected race %v but got race %v", v.expected, result)
		}
	}
}

func getTestRace(meetingID string, number int32, actualStart int64) *proto.Race {
	return &proto.Race{
		MeetingId:   meetingID,
		Number:      number,
		RaceId:      fmt.Sprintf("race-%v", number),
		ActualStart: actualStart,
		Name:        fmt.Sprintf("%v", number),
	}
}

func (repo *MockRepo) ListMeetingsByDate(start, end int64) ([]*proto.Meeting, error) {

	var result []*proto.Meeting
	for _, v := range repo.meetings {
		if v.ScheduledStart >= start && v.ScheduledStart <= end {
			result = append(result, v)
		}
	}

	return result, nil
}

func (repo *MockRepo) ListRacesByMeetingDate(start, end int64) ([]*proto.Race, error) {

	var result []*proto.Race
	for _, v := range repo.races {
		if v.MeetingStart >= start && v.MeetingStart <= end {
			result = append(result, v)
		}
	}

	return result, nil
}

func (repo *MockRepo) AddMeetings(meetings []*proto.Meeting) error {

	repo.meetings = append(repo.meetings, meetings...)
	return nil
}

func (repo *MockRepo) AddRaces(races []*proto.Race) error {
	repo.races = append(repo.races, races...)
	return nil
}

func (repo *MockRepo) AddSelections(selections []*proto.Selection) error {
	repo.selections = append(repo.selections, selections...)
	return nil
}

func (repo *MockRepo) UpdateRace(race *proto.Race) error {

	var existing *proto.Race
	for _, v := range repo.races {
		if v.RaceId == race.RaceId {
			existing = v
		}
	}

	if existing == nil {
		return fmt.Errorf("Expected race id %v does not exist", race.RaceId)
	}

	existing.ScheduledStart = race.ScheduledStart
	existing.ActualStart = race.ActualStart
	existing.Status = race.Status
	existing.Results = race.Results
	existing.LastUpdated = time.Now().Unix()

	return nil
}

func (repo *MockRepo) GetRace(raceID string) (*proto.Race, error) {

	for _, v := range repo.races {
		if v.RaceId == raceID {
			return v, nil
		}
	}

	return nil, fmt.Errorf("No race found with id %v", raceID)
}

func (repo *MockRepo) GetMeeting(meetingID string) (*proto.Meeting, error) {

	for _, v := range repo.meetings {
		if v.MeetingId == meetingID {
			return v, nil
		}
	}

	return nil, fmt.Errorf("No meeting found with id %v", meetingID)
}

func (repo *MockRepo) ListRacesByMeetingID(meetingID string) ([]*proto.Race, error) {
	var result []*proto.Race
	for _, v := range repo.races {
		if v.MeetingId == meetingID {
			result = append(result, v)
		}
	}

	return result, nil
}

func (repo *MockRepo) ListSelectionsByRaceID(raceID string) ([]*proto.Selection, error) {

	var result []*proto.Selection
	for _, v := range repo.selections {
		if v.RaceId == raceID {
			result = append(result, v)
		}
	}

	return result, nil
}

func (repo *MockRepo) UpdateSelection(selection *proto.Selection) error {

	var existing *proto.Selection
	for _, v := range repo.selections {
		if v.SelectionId == selection.SelectionId {
			existing = v
		}
	}

	if existing == nil {
		return fmt.Errorf("No selection found with id %v", selection.SelectionId)
	}

	if selection.Scratched {
		existing.Scratched = true
	} else {
		existing.BarrierNumber = selection.BarrierNumber
		existing.Jockey = selection.Jockey
		existing.Number = selection.Number
		existing.SourceCompetitorId = selection.SourceCompetitorId
		existing.Name = selection.Name
	}
	existing.LastUpdated = time.Now().Unix()

	return nil
}

func (repo *MockRepo) Close() {}

func (repo *MockRepo) NewSession() Repository {
	return repo
}

type MockBroker struct {
	messages []*broker.Message
}

type MockSubscriber struct {
	id    string
	topic string
}

func (b *MockBroker) Options() broker.Options {
	return b.Options()
}

func (b *MockBroker) Address() string {
	return ""
}

func (b *MockBroker) String() string {
	return "mock"
}

func (b *MockBroker) Connect() error {
	return nil
}

func (b *MockBroker) Disconnect() error {
	return nil
}

func (b *MockBroker) Init(...broker.Option) error {
	return nil
}

func (b *MockBroker) Publish(topic string, message *broker.Message, opts ...broker.PublishOption) error {
	b.messages = append(b.messages, message)
	return nil
}

func (b *MockBroker) Subscribe(topic string, handler broker.Handler, opts ...broker.SubscribeOption) (broker.Subscriber, error) {
	sub := &MockSubscriber{}
	return sub, nil
}

func (m *MockSubscriber) Options() broker.SubscribeOptions {
	return broker.SubscribeOptions{}
}

func (m *MockSubscriber) Topic() string {
	return m.topic
}

func (m *MockSubscriber) Unsubscribe() error {
	return nil
}

func getMockStats() *mockStats {
	return &mockStats{
		counters: make(map[string]int),
	}
}

type mockStats struct {
	counters map[string]int
}

type mockTiming struct {
	start time.Time
	end   time.Time
}

func (m *mockStats) Increment(bucket string) {
	m.counters[bucket]++
}

func (m *mockStats) NewTiming() statsTiming {
	return &mockTiming{
		start: time.Now(),
	}
}

func (t *mockTiming) Send(bucket string) {
	t.end = time.Now()
}
