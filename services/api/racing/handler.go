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

	handlerScheduleTiming  = "racing.api.handler.schedule.timing"
	handlerScheduleSuccess = "racing.api.handler.schedule.success"
	handlerScheduleFailed  = "racing.api.handler.schedule.failed"
)

// Racing is the handler for the racing api
type Racing struct{}

// Racecard is called by the API as /racing/racecard
func (r *Racing) Racecard(ctx context.Context, req *api.Request, rsp *api.Response) error {

	// log := logWithContext(ctx, "handler.Racecard")

	id, ok := req.Get["race_id"]
	if !ok || len(id.Values) == 0 {
		return errors.BadRequest("go.micro.api.racing", "no race id")
	}

	// call racing.GetRace

	// if valid:
	// call racing.GetMeeting
	// call racing.ListSelections

	b, _ := json.Marshal(map[string]string{
		"message": "request received",
	})

	rsp.Header = make(map[string]*api.Pair)
	corr := &api.Pair{
		Key:    headerCorrelationID,
		Values: []string{ctx.Value(correlationID).(string)},
	}
	rsp.Header[headerCorrelationID] = corr
	rsp.StatusCode = 200
	rsp.Body = string(b)

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

	resp, err := createSchedule(start, meetingResp, raceResp)

	b, _ := json.Marshal(resp)

	rsp.Header = make(map[string]*api.Pair)
	corr := &api.Pair{
		Key:    headerCorrelationID,
		Values: []string{ctx.Value(correlationID).(string)},
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
		d, err := time.Parse("00", date.Values[0])
		if err != nil {
			return nil, nil, errors.BadRequest("go.micro.api.racing", "invalid date format")
		}

		start = d
		end = d.Add(time.Hour * 24).Add(time.Second * -1)
	}
	return &start, &end, nil
}

func createSchedule(date *time.Time, m *racing.ListMeetingsByDateResponse, r *racing.ListRacesByMeetingDateResponse) (*response.RaceSchedule, error) {
	schedule := &response.RaceSchedule{
		HasRaces: len(r.Races) > 0,
		RaceDate: date.Format("2006-01-02"),
		Meetings: make([]response.RaceScheduleMeeting, len(m.Meetings)),
		Races:    make([]response.RaceScheduleRace, len(r.Races)),
	}

	for i, r := range r.Races {
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

	for i, m := range m.Meetings {
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
