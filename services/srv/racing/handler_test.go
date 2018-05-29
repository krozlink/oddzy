package main

import (
	"context"
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	"testing"
)

type MockRepo struct {
	meetings   []*proto.Meeting
	races      []*proto.Race
	selections []*proto.Selection
}

func (repo *MockRepo) ListMeetingsByDate(start, end int64) ([]*proto.Meeting, error) {

	var result []*proto.Meeting
	for _, v := range repo.meetings {
		if v.ScheduledStart >= start && v.ScheduledStart <= end {
			result = append(result, v)
		}
	}

	return nil, nil
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

func (repo *MockRepo) UpdateRace(race *proto.Race, selections []*proto.Selection) error {
	return nil
}

func (repo *MockRepo) Close() {}

func (repo *MockRepo) NewSession() Repository {
	return repo
}

func TestCanRetrieveRacesByMeetingDateRange(t *testing.T) {
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
	srv := NewRacingService(repo)

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

		if err != nil {
			t.Error("Expected no error when calling ListRacesByMeetingDate")
		}

		if len(resp.Races) != v.expected {
			t.Errorf("Expected %v races, got %v", v.expected, len(resp.Races))
		}
	}
}

func TestDatesMandatoryForListRacesByMeetingDate(t *testing.T) {
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
	srv := NewRacingService(repo)

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

		if err == nil && v.error {
			t.Errorf("Expected error when calling ListRacesByMeetingDate for range %v to %v", v.start, v.end)
		}

		if err != nil && !v.error {
			t.Errorf("Expected no error when calling ListRacesByMeetingDate for range %v to %v", v.start, v.end)
		}
	}
}
