package main

import (
	"context"
	"fmt"
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	"gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
}

// GetRepo returns the repository to be used when accessing racing data
func (s *service) GetRepo() Repository {
	return &RacingRepository{s.session.Clone()}
}

// ListMeetingsByDate will return all meetings between the provided start and end dates
func (s *service) ListMeetingsByDate(ctx context.Context, req *proto.ListMeetingsByDateRequest, resp *proto.ListMeetingsByDateResponse) error {
	if req.StartDate == 0 {
		return fmt.Errorf("Start date is a mandatory field")
	}

	if req.EndDate == 0 {
		return fmt.Errorf("End date is a mandatory field")
	}

	repo := s.GetRepo()
	defer repo.Close()

	meetings, err := repo.ListMeetingsByDate(req.StartDate, req.EndDate)
	if err != nil {
		return err
	}

	resp.Meetings = meetings

	return nil
}

// ListRacesByMeetingDate will return races between the provided start and end dates
func (s *service) ListRacesByMeetingDate(ctx context.Context, req *proto.ListRacesByMeetingDateRequest, resp *proto.ListRacesByMeetingDateResponse) error {
	if req.StartDate == 0 {
		return fmt.Errorf("Start date is a mandatory field")
	}

	if req.EndDate == 0 {
		return fmt.Errorf("End date is a mandatory field")
	}

	repo := s.GetRepo()
	defer repo.Close()

	races, err := repo.ListRacesByMeetingDate(req.StartDate, req.EndDate)
	if err != nil {
		return err
	}

	resp.Races = races

	return nil
}

// AddMeetings will save the provided meetings
func (s *service) AddMeetings(ctx context.Context, req *proto.AddMeetingsRequest, resp *proto.AddMeetingsResponse) error {
	errors := ""
	for i, v := range req.Meetings {
		if v.MeetingId == "" {
			errors += fmt.Sprintf("Meeting id not provided on meeting %v\n", i)
		}

		if v.SourceId == "" {
			errors += fmt.Sprintf("Source id not provided on meeting %v\n", i)
		}

		if len(v.RaceIds) == 0 {
			errors += fmt.Sprintf("No race ids provided meeting %v\n", i)
		}

		if v.ScheduledStart == 0 {
			errors += fmt.Sprintf("No scheduled start date provided for meeting %v\n", i)
		}

		if v.RaceType == "" {
			errors += fmt.Sprintf("No race type provided for meeting %v\n", i)
		}
	}

	if errors != "" {
		return fmt.Errorf(errors)
	}

	repo := s.GetRepo()
	defer repo.Close()

	err := repo.AddMeetings(req.Meetings)
	if err != nil {
		return err
	}

	resp.Created = true

	return nil
}

// AddRaces will save the provided races
func (s *service) AddRaces(ctx context.Context, req *proto.AddRacesRequest, resp *proto.AddRacesResponse) error {
	errors := ""
	for i, v := range req.Races {
		if v.RaceId == "" {
			errors += fmt.Sprintf("Race id not provided on race %v\n", i)
		}

		if v.SourceId == "" {
			errors += fmt.Sprintf("Source id not provided on race %v\n", i)
		}

		if v.Name == "" {
			errors += fmt.Sprintf("No name provided for race %v\n", i)
		}

		if v.Number == 0 {
			errors += fmt.Sprintf("No number provided for race %v\n", i)
		}

		if v.ScheduledStart == 0 {
			errors += fmt.Sprintf("No scheduled start time provided for race %v\n", i)
		}

		if v.Status == "" {
			errors += fmt.Sprintf("No status provided for race %v\n", i)
		}

		if v.LastUpdated == 0 {
			errors += fmt.Sprintf("No last updated time provided for race %v\n", i)
		}

		if v.MeetingStart == 0 {
			errors += fmt.Sprintf("No meeting start time provided for race %v\n", i)
		}
	}

	if errors != "" {
		return fmt.Errorf(errors)
	}

	repo := s.GetRepo()
	defer repo.Close()

	err := repo.AddRaces(req.Races)
	if err != nil {
		return err
	}

	resp.Created = true

	return nil
}

// UpdateRace will update the race and selection data for the provided race
func (s *service) UpdateRace(ctx context.Context, req *proto.UpdateRaceRequest, resp *proto.UpdateRaceResponse) error {
	err := validateRace(req)
	if err != nil {
		return err
	}

	repo := s.GetRepo()
	defer repo.Close()

	err = repo.UpdateRace(req.Race, req.Selections)

	return err
}

func validateRace(req *proto.UpdateRaceRequest) error {
	errors := ""
	if req.Race.RaceId == "" {
		errors += fmt.Sprintln("Race id not provided")
	}

	if req.Race.SourceId == "" {
		errors += fmt.Sprintln("Source id not provided for the race")
	}

	if req.Race.ScheduledStart == 0 {
		errors += fmt.Sprintln("Scheduled start time not provided for the race")
	}

	if req.Race.Number == 0 {
		errors += fmt.Sprintln("Number not provided for the race")
	}

	if req.Race.Name == "" {
		errors += fmt.Sprintln("Name not provided for the race")
	}

	if req.Race.Status == "" {
		errors += fmt.Sprintln("Status not provided for the race")
	}

	if len(req.Selections) == 0 {
		errors += fmt.Sprintln("No selections provided for the race")
	}

	for i, v := range req.Selections {
		if v.SourceId == "" {
			errors += fmt.Sprintf("Source id not provided for selection %v\n", i)
		}

		if v.SelectionId == "" {
			errors += fmt.Sprintf("Selection id not provided for selection %v\n", i)
		}

		if v.SourceCompetitorId == "" {
			errors += fmt.Sprintf("Source competitor id not provided for selection %v\n", i)
		}

		if v.RaceId == "" {
			errors += fmt.Sprintf("Race id not provided for selection %v\n", i)
		}

		if v.RaceId != req.Race.RaceId {
			errors += fmt.Sprintf("Race id for selection %v does not match the race\n", i)
		}

		if v.Name == "" {
			errors += fmt.Sprintf("Name not provided for selection %v\n", i)
		}

		if v.Jockey == "" {
			errors += fmt.Sprintf("Jockey not provided for selection %v\n", i)
		}

		if v.Number == 0 {
			errors += fmt.Sprintf("Number not provided for selection %v\n", i)
		}

		if v.BarrierNumber == 0 {
			errors += fmt.Sprintf("Barrier number not provided for selection %v\n", i)
		}
	}

	if errors != "" {
		return fmt.Errorf(errors)
	}

	return nil
}
