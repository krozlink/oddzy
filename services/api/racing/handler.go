package main

import (
	"context"
	"encoding/json"
	"github.com/krozlink/oddzy/services/api/racing/response"
	racing "github.com/krozlink/oddzy/services/srv/racing/proto"
	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
	"sync"
	"time"
)

const (
	headerCorrelationID = "X-Correlation-Id"

	handlerScheduleTiming  = "oddzy.racing.api.handler.schedule.timing"
	handlerScheduleSuccess = "oddzy.racing.api.handler.schedule.success"
	handlerScheduleFailed  = "oddzy.racing.api.handler.schedule.failed"
	handlerRacecardTiming  = "oddzy.racing.api.handler.racecard.timing"
	handlerRacecardSuccess = "oddzy.racing.api.handler.racecard.success"
	handlerRacecardFailed  = "oddzy.racing.api.handler.racecard.failed"
)

// Racing is the handler for the racing api
type Racing struct{}

// Racecard is called by the API as /racing/racecard
func (r *Racing) Racecard(ctx context.Context, req *api.Request, rsp *api.Response) error {

	timing := stats.NewTiming()
	defer timing.Send(handlerRacecardTiming)
	log := logWithContext(ctx, "handler.Racecard")

	v, ok := req.Get["race_id"]
	if !ok || len(v.Values) == 0 {
		stats.Increment(handlerScheduleFailed)
		return errors.BadRequest("go.micro.api.racing", "no race id")
	}

	raceID := v.Values[0]

	client, ok := RacingFromContext(ctx)
	if !ok {
		stats.Increment(handlerScheduleFailed)
		log.Error("No racing client exists in context")
		return errors.InternalServerError("go.micro.api.racing", "invalid context")
	}

	raceReq := &racing.GetRaceRequest{
		RaceId: raceID,
	}
	raceResp, err := client.GetRace(ctx, raceReq)
	if err != nil {
		log.Errorf("Unable to read race with id %v - %v", raceID, err)
		stats.Increment(handlerScheduleFailed)
		return errors.BadRequest("go.micro.api.racing", "invalid race id")
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	var meetingResp *racing.GetMeetingResponse
	var selectionsResp *racing.ListSelectionsResponse

	go func() {
		defer wg.Done()
		meetingReq := &racing.GetMeetingRequest{
			MeetingId: raceResp.Race.MeetingId,
		}
		meetingResp, err = client.GetMeeting(ctx, meetingReq)
		if err != nil {
			log.Fatalf("Unable to read meeting with id %v - %v", raceResp.Race.MeetingId, err)
		}
	}()

	go func() {
		defer wg.Done()
		selectionsReq := &racing.ListSelectionsRequest{
			RaceId: raceResp.Race.RaceId,
		}
		selectionsResp, err = client.ListSelections(ctx, selectionsReq)
		if err != nil {
			log.Fatalf("Unable to read selections for race %v - %v", raceID, err)
		}
	}()

	wg.Wait()

	resp, err := createRacecard(meetingResp.Meeting, raceResp.Race, selectionsResp.Selections)

	b, _ := json.Marshal(resp)

	rsp.Header = make(map[string]*api.Pair)
	corr := &api.Pair{
		Key:    headerCorrelationID,
		Values: []string{getValueFromMetadata(ctx, string(correlationID))},
	}
	rsp.Header[headerCorrelationID] = corr
	rsp.StatusCode = 200
	rsp.Body = string(b)

	stats.Increment(handlerRacecardSuccess)

	return nil
}

// Schedule is called by the API as /racing/schedule
func (r *Racing) Schedule(ctx context.Context, req *api.Request, rsp *api.Response) error {
	timing := stats.NewTiming()
	defer timing.Send(handlerScheduleTiming)

	log := logWithContext(ctx, "handler.Schedule")

	start, end, err := validateScheduleRequest(req)
	if err != nil {
		stats.Increment(handlerScheduleFailed)
		return err
	}

	client, ok := RacingFromContext(ctx)
	if !ok {
		stats.Increment(handlerScheduleFailed)
		log.Error("No racing client exists in context")
		return errors.InternalServerError("go.micro.api.racing", "invalid context")
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	var meetingResp *racing.ListMeetingsByDateResponse
	var raceResp *racing.ListRacesByMeetingDateResponse

	go func() {
		defer wg.Done()
		meetingReq := &racing.ListMeetingsByDateRequest{
			StartDate: start.Unix(),
			EndDate:   end.Unix(),
		}
		meetingResp, err = client.ListMeetingsByDate(ctx, meetingReq)
		if err != nil {
			log.Fatal("Unable to list meetings")
		}
	}()

	go func() {
		defer wg.Done()
		raceReq := &racing.ListRacesByMeetingDateRequest{
			StartDate: start.Unix(),
			EndDate:   end.Unix(),
		}
		raceResp, err = client.ListRacesByMeetingDate(ctx, raceReq)
		if err != nil {
			log.Fatal("Unable to list races")
		}
	}()

	wg.Wait()

	resp, err := createSchedule(start, meetingResp.Meetings, raceResp.Races)

	b, _ := json.Marshal(resp)

	rsp.Header = make(map[string]*api.Pair)
	corr := &api.Pair{
		Key:    headerCorrelationID,
		Values: []string{getValueFromMetadata(ctx, string(correlationID))},
	}
	rsp.Header[headerCorrelationID] = corr
	rsp.StatusCode = 200
	rsp.Body = string(b)

	stats.Increment(handlerScheduleSuccess)

	return nil
}

func validateScheduleRequest(req *api.Request) (*time.Time, *time.Time, error) {
	var start, end time.Time

	date, ok := req.Get["date"]
	if !ok || len(date.Values) == 0 {
		return nil, nil, errors.BadRequest("go.micro.api.racing", "no date")
	} else if len(date.Values) > 1 {
		return nil, nil, errors.BadRequest("go.micro.api.racing", "only one date required")
	} else {
		d, err := time.Parse("2006-01-02", date.Values[0])
		if err != nil {
			return nil, nil, errors.BadRequest("go.micro.api.racing", "invalid date format")
		}

		start = d
		end = d.Add(time.Hour * 24).Add(time.Second * -1)
	}
	return &start, &end, nil
}

func createRacecard(meeting *racing.Meeting, race *racing.Race, selections []*racing.Selection) (*response.RaceCard, error) {
	m := response.RaceCardMeeting{
		MeetingID: meeting.MeetingId,
		Name:      meeting.Name,
		Country:   meeting.Country,
		RaceType:  meeting.RaceType,
		RaceIds:   meeting.RaceIds,
	}

	s := make([]response.RaceCardSelection, len(selections))
	for i, v := range selections {
		s[i] = response.RaceCardSelection{
			Barrier:     v.BarrierNumber,
			IsScratched: v.Scratched,
			Jockey:      v.Jockey,
			Name:        v.Name,
			Number:      v.Number,
			SelectionID: v.SelectionId,
		}
	}

	racecard := &response.RaceCard{
		LastUpdated:    race.LastUpdated,
		Meeting:        m,
		Name:           race.Name,
		Number:         race.Number,
		RaceID:         race.RaceId,
		Results:        race.Results,
		ScheduledStart: race.ScheduledStart,
		Selections:     s,
		Status:         race.Status,
	}

	return racecard, nil
}

func createSchedule(date *time.Time, meetings []*racing.Meeting, races []*racing.Race) (*response.RaceSchedule, error) {
	schedule := &response.RaceSchedule{
		HasRaces: len(races) > 0,
		RaceDate: date.Format("2006-01-02"),
		Meetings: make([]response.RaceScheduleMeeting, len(meetings)),
		Races:    make([]response.RaceScheduleRace, len(races)),
	}

	for i, r := range races {
		schedule.Races[i] = response.RaceScheduleRace{
			LastUpdated:    r.LastUpdated,
			Name:           r.Name,
			Number:         r.Number,
			RaceID:         r.RaceId,
			Results:        r.Results,
			ScheduledStart: r.ScheduledStart,
			Status:         r.Status,
		}
	}

	for i, m := range meetings {
		schedule.Meetings[i] = response.RaceScheduleMeeting{
			Country:        m.Country,
			LastUpdated:    m.LastUpdated,
			MeetingID:      m.MeetingId,
			Name:           m.Name,
			RaceType:       m.RaceType,
			ScheduledStart: m.ScheduledStart,
			RaceIds:        m.RaceIds,
		}
	}

	return schedule, nil
}
