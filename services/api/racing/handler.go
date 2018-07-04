package main

import (
	"context"
	// "encoding/json"
	"github.com/micro/go-log"

	// racing "github.com/krozlink/oddzy/services/srv/racing/proto"
	api "github.com/micro/go-api/proto"
	// "github.com/micro/go-micro/errors"
)

type handler struct{}

// RaceCard  is called by the API as /racing/racecard
func (r *handler) Racecard(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received handler.RaceCard request")

	// call racing.GetRace

	return nil
}

// Schedule is called by the API as /racing/schedule
func (r *handler) Schedule(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received handler.Schedule request")

	return nil
}
