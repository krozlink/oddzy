package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-log"

	// racing "github.com/krozlink/oddzy/services/srv/racing/proto"
	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
)

type handler struct{}

// RaceCard  is called by the API as /racing/racecard
func (r *handler) Racecard(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received handler.RaceCard request")

	id, ok := req.Get["race_id"]
	if !ok || len(id.Values) == 0 {
		return errors.BadRequest("go.micro.api.racing", "no race id")
	}

	// call racing.GetRace
	// call racing.GetMeeting
	// call racing.ListSelections

	b, _ := json.Marshal(map[string]string{
		"message": "racecard request received",
	})

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}

// Schedule is called by the API as /racing/schedule
func (r *handler) Schedule(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received handler.Schedule request")

	date, ok := req.Get["date"]
	if !ok || len(date.Values) == 0 {
		return errors.BadRequest("go.micro.api.racing", "no date")
	}

	// check date is valid

	// call racing.ListMeetingsByDate
	// call racing.ListRacesByMeetingDate
	b, _ := json.Marshal(map[string]string{
		"message": "schedule request received",
	})

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
