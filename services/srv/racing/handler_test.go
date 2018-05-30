package main

import (
	"context"
	"fmt"
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
	//TODO: implement mock UpdateRace function
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

func TestCanRetrieveMeetingsByDateRange(t *testing.T) {
	t.Fail()
}

func TestDatesMandatoryForListMeetingsByDate(t *testing.T) {
	t.Fail()
}

func TestAddMeetingsValidation(t *testing.T) {
	t.Fail()
}

func TestAddMeetingsAddsToRepo(t *testing.T) {
	t.Fail()
}

func TestAddRacesValidation(t *testing.T) {
	t.Fail()
}

func TestAddRacesAddsToRepo(t *testing.T) {
	t.Fail()
}

func TestUpdateRaceCreatesSelectionsOnInitalCall(t *testing.T) {
	t.Fail()
}

func TestUpdateRaceFailsWhenNumberOfSelectionsChange(t *testing.T) {
	t.Fail()
}

func TestUpdateRaceModifiesExistingSelections(t *testing.T) {
	t.Fail()
}

func TestUpdateRaceNotifiesOnRaceChange(t *testing.T) {
	t.Fail()
}

func TestUpdateRaceNotifiesOnSelectionChange(t *testing.T) {
	t.Fail()
}

func TestUpdateRaceNoNotificationIfNoChange(t *testing.T) {
	t.Fail()
}
