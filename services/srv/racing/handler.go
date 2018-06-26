package main

import (
	"context"
	"encoding/json"
	"fmt"
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	"github.com/micro/go-micro/broker"
	"sort"
	"strconv"
)

const (
	listMeetingsTiming  = "racing.service.listmeetings.timing"
	listMeetingsSuccess = "racing.service.listmeetings.success"
	listMeetingsFailed  = "racing.service.listmeetings.failed"
	listRacesTiming     = "racing.service.listraces.timing"
	listRacesSuccess    = "racing.service.listraces.success"
	listRacesFailed     = "racing.service.listraces.failed"
	addMeetingsTiming   = "racing.service.addmeetings.timing"
	addMeetingsSuccess  = "racing.service.addmeetings.success"
	addMeetingsFailed   = "racing.service.addmeetings.failed"
	addRacesTiming      = "racing.service.addraces.timing"
	addRacesSuccess     = "racing.service.addraces.success"
	addRacesFailed      = "racing.service.addraces.failed"
	updateRaceTiming    = "racing.service.updaterace.timing"
	updateRaceSuccess   = "racing.service.updaterace.success"
	updateRaceFailed    = "racing.service.updaterace.failed"
	getNextRaceTiming   = "racing.service.getnextrace.timing"
	getNextRaceSuccess  = "racing.service.getnextrace.success"
	getNextRaceFailed   = "racing.service.getnextrace.failed"
)

// RacingService is for interacing with racing data such as meetings and races
type RacingService struct {
	repo   Repository
	broker broker.Broker
}

const raceUpdateTopic = "race.updated"

// NewRacingService returns a new instance of the service using the provided repository
func NewRacingService(r Repository, b broker.Broker) *RacingService {
	return &RacingService{
		repo:   r,
		broker: b,
	}
}

// GetRepo returns the repository to be used when accessing racing data
func (s *RacingService) GetRepo() Repository {
	return s.repo.NewSession()
}

// ListMeetingsByDate will return all meetings between the provided start and end dates
func (s *RacingService) ListMeetingsByDate(ctx context.Context, req *proto.ListMeetingsByDateRequest, resp *proto.ListMeetingsByDateResponse) error {
	timing := stats.NewTiming()
	defer timing.Send(listMeetingsTiming)

	log := logWithField("function", "handler.ListMeetingsByDate")

	if req.StartDate == 0 {
		err := fmt.Errorf("Start date is a mandatory field")
		log.Error(err)
		return err
	}

	if req.EndDate == 0 {
		err := fmt.Errorf("End date is a mandatory field")
		log.Error(err)
		return err
	}

	repo := s.GetRepo()
	defer repo.Close()

	meetings, err := repo.ListMeetingsByDate(req.StartDate, req.EndDate)
	if err != nil {
		stats.Increment(listMeetingsFailed)
		log.Error(err)
		return err
	}

	resp.Meetings = meetings

	stats.Increment(listMeetingsSuccess)
	return nil
}

// ListRacesByMeetingDate will return races between the provided start and end dates
func (s *RacingService) ListRacesByMeetingDate(ctx context.Context, req *proto.ListRacesByMeetingDateRequest, resp *proto.ListRacesByMeetingDateResponse) error {
	timing := stats.NewTiming()
	defer timing.Send(listRacesTiming)
	log := logWithField("function", "handler.ListRacesByMeetingDate")

	if req.StartDate == 0 {
		err := fmt.Errorf("Start date is a mandatory field")
		log.Error(err)
		return err
	}

	if req.EndDate == 0 {
		err := fmt.Errorf("End date is a mandatory field")
		log.Error(err)
		return err
	}

	repo := s.GetRepo()
	defer repo.Close()

	races, err := repo.ListRacesByMeetingDate(req.StartDate, req.EndDate)
	if err != nil {
		stats.Increment(listRacesFailed)
		log.Error(err)
		return err
	}

	resp.Races = races

	stats.Increment(listRacesSuccess)
	return nil
}

// AddMeetings will save the provided meetings
func (s *RacingService) AddMeetings(ctx context.Context, req *proto.AddMeetingsRequest, resp *proto.AddMeetingsResponse) error {
	log := logWithField("function", "AddMeetings")
	timing := stats.NewTiming()
	defer timing.Send(addMeetingsTiming)

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
		log.Errorf("Validation failed when adding %v meetings - %v", len(req.Meetings), errors)
		stats.Increment(addMeetingsFailed)
		return fmt.Errorf(errors)
	}

	repo := s.GetRepo()
	defer repo.Close()

	err := repo.AddMeetings(req.Meetings)
	if err != nil {
		log.Errorf("Failed to add %v meetings - %v", len(req.Meetings), err)
		stats.Increment(addMeetingsFailed)
		return err
	}

	resp.Created = true

	log.Debugf("%v meetings successfully added", len(req.Meetings))
	stats.Increment(addMeetingsSuccess)
	return nil
}

// AddRaces will save the provided races
func (s *RacingService) AddRaces(ctx context.Context, req *proto.AddRacesRequest, resp *proto.AddRacesResponse) error {
	log := logWithField("function", "AddRaces")
	timing := stats.NewTiming()
	defer timing.Send(addRacesTiming)

	errors := ""
	for i, v := range req.Races {
		id := v.RaceId
		if v.RaceId == "" {
			errors += fmt.Sprintf("Race id not provided on race %v\n", i)
		} else {
			id = strconv.Itoa(i)
		}

		if v.SourceId == "" {
			errors += fmt.Sprintf("Source id not provided on race %v\n", id)
		}

		if v.Name == "" {
			errors += fmt.Sprintf("No name provided for race %v\n", id)
		}

		if v.Number == 0 {
			errors += fmt.Sprintf("No number provided for race %v\n", id)
		}

		if v.ScheduledStart == 0 {
			errors += fmt.Sprintf("No scheduled start time provided for race %v\n", id)
		}

		if v.Status == "" {
			errors += fmt.Sprintf("No status provided for race %v\n", id)
		}

		if v.DateCreated != 0 {
			errors += fmt.Sprintf("Date created time (%v) should not be set when adding race %v\n", v.DateCreated, id)
		}

		if v.LastUpdated != 0 {
			errors += fmt.Sprintf("Last update time (%v) should not be set when adding race %v\n", v.LastUpdated, id)
		}

		if v.MeetingStart == 0 {
			errors += fmt.Sprintf("No meeting start time provided for race %v\n", id)
		}
	}

	if errors != "" {
		log.Errorf("Failed to validate %v races - %v", len(req.Races), errors)
		stats.Increment(addRacesFailed)
		return fmt.Errorf(errors)
	}

	repo := s.GetRepo()
	defer repo.Close()

	err := repo.AddRaces(req.Races)
	if err != nil {
		log.Errorf("Failed to add %v races - %v", len(req.Races), errors)
		stats.Increment(addRacesFailed)
		return err
	}

	resp.Created = true
	stats.Increment(addRacesSuccess)

	return nil
}

// UpdateRace will update the race and (optionally) selection data for the provided race
func (s *RacingService) UpdateRace(ctx context.Context, req *proto.UpdateRaceRequest, resp *proto.UpdateRaceResponse) error {
	log := logWithField("function", "UpdateRace")
	timing := stats.NewTiming()
	defer timing.Send(updateRaceTiming)

	err := validateRace(req)
	if err != nil {
		log.Errorf("Failed to update race. Error: %v", err)
		stats.Increment(updateRaceFailed)
		return err
	}

	repo := s.GetRepo()
	defer repo.Close()

	originalRace, err := repo.GetRace(req.RaceId)
	if err != nil {
		stats.Increment(updateRaceFailed)
		return err
	}

	race := &proto.RaceUpdatedMessage{
		RaceId:         req.RaceId,
		Results:        req.Results,
		ScheduledStart: req.ScheduledStart,
		ActualStart:    req.ActualStart,
		Status:         req.Status,
		Selections:     req.Selections,
	}

	raceUpdated := hasRaceChanged(originalRace, race)
	if raceUpdated {
		err = repo.UpdateRace(race)
	}

	selectionUpdated := false

	if len(req.Selections) > 0 { // May not include selection data in an update
		originalSelections, err := repo.ListSelectionsByRaceID(req.RaceId)
		if err != nil {
			stats.Increment(updateRaceFailed)
			return err
		}

		if len(originalSelections) == 0 { // Add selections if included and none already exist
			err = repo.AddSelections(req.Selections)
			if err != nil {
				stats.Increment(updateRaceFailed)
				return err
			}
			selectionUpdated = true
		} else if len(originalSelections) < len(req.Selections) {
			return fmt.Errorf("Number of selections unexpectedly changed from %v to %v", len(originalSelections), len(req.Selections))
		} else {

			for _, v := range req.Selections {
				o := getSelectionByID(v.SelectionId, originalSelections)
				if o == nil {
					stats.Increment(updateRaceFailed)
					return fmt.Errorf("Expected to find selection %v", v.SelectionId)
				}

				if hasSelectionChanged(o, v) {
					selectionUpdated = true
					err = repo.UpdateSelection(v)
					if err != nil {
						stats.Increment(updateRaceFailed)
						return err
					}
				}
			}

			// When the number of selections has decreased it means a horse has been scratched
			if len(originalSelections) > len(req.Selections) {

				// find the selection that no longer exists
				for _, v := range originalSelections {
					if getSelectionByID(v.SelectionId, req.Selections) == nil {

						// scratch it
						selectionUpdated = true
						u := &proto.Selection{
							SelectionId: v.SelectionId,
							Scratched:   true,
						}
						err = repo.UpdateSelection(u)
						if err != nil {
							stats.Increment(updateRaceFailed)
							return err
						}
					}
				}
			}
		}
	}

	if raceUpdated || selectionUpdated {
		if err := s.publishRaceUpdate(race, req.Selections); err != nil {
			stats.Increment(updateRaceFailed)
			return err
		}
	}

	stats.Increment(updateRaceSuccess)
	return nil
}

// GetNextRace returns the next race which has not completed, or nil if all races have completed
func (s *RacingService) GetNextRace(ctx context.Context, req *proto.GetNextRaceRequest, resp *proto.GetNextRaceResponse) error {
	timing := stats.NewTiming()
	defer timing.Send(getNextRaceTiming)

	repo := s.GetRepo()

	races, err := repo.ListRacesByMeetingID(req.MeetingId)
	if err != nil {
		stats.Increment(getNextRaceFailed)
		return err
	}

	// ensure races are in number order
	sort.Slice(races, func(i, j int) bool {
		return races[i].Number < races[j].Number
	})
	for _, v := range races {
		if v.ActualStart == 0 {
			resp.Race = v
			stats.Increment(getNextRaceSuccess)
			return nil
		}
	}

	resp.Race = nil
	stats.Increment(getNextRaceSuccess)
	return nil
}

func (s *RacingService) publishRaceUpdate(update *proto.RaceUpdatedMessage, selections []*proto.Selection) error {
	log := logWithField("function", "publishRaceUpdate")
	b, err := json.Marshal(update)
	if err != nil {
		return err
	}
	msg := &broker.Message{
		Header: map[string]string{
			"race_id": update.RaceId,
		},
		Body: b,
	}

	if err := s.broker.Publish(raceUpdateTopic, msg); err != nil {
		log.Errorf("Publish failed: %v", err)
		return fmt.Errorf("Publish failed: %v", err)
	}

	return nil
}

func validateRace(req *proto.UpdateRaceRequest) error {
	errors := ""
	if req.RaceId == "" {
		errors += fmt.Sprintln("Race id not provided")
	}

	if req.ScheduledStart == 0 {
		errors += fmt.Sprintln("Scheduled start time not provided for the race")
	}

	if req.Status == "" {
		errors += fmt.Sprintln("Status not provided for the race")
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

		if v.RaceId != req.RaceId {
			errors += fmt.Sprintf("Race id for selection %v does not match the race\n", i)
		}

		if v.Name == "" {
			errors += fmt.Sprintf("Name not provided for selection %v\n", i)
		}

		if v.Number == 0 {
			errors += fmt.Sprintf("Number not provided for selection %v\n", i)
		}
	}

	if errors != "" {
		return fmt.Errorf(errors)
	}

	return nil
}

func hasRaceChanged(from *proto.Race, to *proto.RaceUpdatedMessage) bool {
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
