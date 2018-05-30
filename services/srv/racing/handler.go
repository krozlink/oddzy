package main

import (
	"context"
	"fmt"
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	"sort"
)

// RacingService is for interacing with racing data such as meetings and races
type RacingService struct {
	repo Repository
}

// NewRacingService returns a new instance of the service using the provided repository
func NewRacingService(r Repository) *RacingService {
	return &RacingService{
		repo: r,
	}
}

// GetRepo returns the repository to be used when accessing racing data
func (s *RacingService) GetRepo() Repository {
	return s.repo.NewSession()
}

// ListMeetingsByDate will return all meetings between the provided start and end dates
func (s *RacingService) ListMeetingsByDate(ctx context.Context, req *proto.ListMeetingsByDateRequest, resp *proto.ListMeetingsByDateResponse) error {
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
func (s *RacingService) ListRacesByMeetingDate(ctx context.Context, req *proto.ListRacesByMeetingDateRequest, resp *proto.ListRacesByMeetingDateResponse) error {
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
func (s *RacingService) AddMeetings(ctx context.Context, req *proto.AddMeetingsRequest, resp *proto.AddMeetingsResponse) error {
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
func (s *RacingService) AddRaces(ctx context.Context, req *proto.AddRacesRequest, resp *proto.AddRacesResponse) error {
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
func (s *RacingService) UpdateRace(ctx context.Context, req *proto.UpdateRaceRequest, resp *proto.UpdateRaceResponse) error {
	err := validateRace(req)
	if err != nil {
		return err
	}

	repo := s.GetRepo()
	defer repo.Close()

	originalRace, err := repo.GetRace(req.Race.RaceId)
	if err != nil {
		return err
	}

	raceUpdated := hasRaceChanged(originalRace, req.Race)
	if raceUpdated {
		err = repo.UpdateRace(req.Race)
	}

	originalSelections, err := repo.ListSelectionsByRaceID(req.Race.RaceId)
	if err != nil {
		return err
	}

	selectionUpdated := false

	if len(originalSelections) == 0 {
		err = repo.AddSelections(req.Selections)
		if err != nil {
			return err
		}
		selectionUpdated = true
	} else if len(originalSelections) != len(req.Selections) {
		return fmt.Errorf("Number of selections unexpectedly changed from %v to %v", len(originalSelections), len(req.Selections))
	} else {
		for _, v := range req.Selections {
			o := getSelectionByID(v.SelectionId, originalSelections)
			if o == nil {
				return fmt.Errorf("Expected to find selection %v", v.SelectionId)
			}

			if hasSelectionChanged(o, v) {
				selectionUpdated = true
				err = repo.UpdateSelection(v)
				if err != nil {
					return err
				}
			}
		}
	}

	if raceUpdated || selectionUpdated {
		//TODO: notify of change
	}

	return nil
}

// GetNextRace returns the next race which has not completed, or nil if all races have completed
func (s *RacingService) GetNextRace(ctx context.Context, req *proto.GetNextRaceRequest, resp *proto.GetNextRaceResponse) error {
	repo := s.GetRepo()

	races, err := repo.ListRacesByMeetingID(req.MeetingId)
	if err != nil {
		return err
	}

	// ensure races are in number order
	sort.Slice(races, func(i, j int) bool {
		return races[i].Number < races[j].Number
	})
	for _, v := range races {
		if v.ActualStart == 0 {
			resp.Race = v
			return nil
		}
	}

	resp.Race = nil
	return nil
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

func hasRaceChanged(from, to *proto.Race) bool {
	return from.ScheduledStart != to.ScheduledStart ||
		from.ActualStart != to.ActualStart ||
		from.Status != to.Status ||
		from.Results != to.Results
}

func hasSelectionChanged(from, to *proto.Selection) bool {
	return from.BarrierNumber != to.BarrierNumber ||
		from.Jockey != to.Jockey ||
		from.SourceCompetitorId != to.SourceCompetitorId ||
		from.Number != to.Number
}

func getSelectionByID(selectionID string, selections []*proto.Selection) *proto.Selection {
	for _, v := range selections {
		if v.SelectionId == selectionID {
			return v
		}
	}
	return nil
}
