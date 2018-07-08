package main

import (
	"context"
	racing "github.com/krozlink/oddzy/services/srv/racing/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/pborman/uuid"
)

const (
	racingService metaContext = "racing_service"
)

// RacingFromContext retrieves the racing client from the Context
func RacingFromContext(ctx context.Context) (racing.RacingService, bool) {
	c, ok := ctx.Value(racingService).(racing.RacingService)
	return c, ok
}

// CorrelationWrapper adds a correlation id to each request
func CorrelationWrapper(service micro.Service) server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = metadata.NewContext(ctx, map[string]string{
				string(correlationID): uuid.NewUUID().String(),
			})
			return fn(ctx, req, rsp)
		}
	}
}

// RacingWrapper returns a wrapper for the RacingClient
func RacingWrapper(service micro.Service) server.HandlerWrapper {
	client := racing.NewRacingService(racingServiceName, service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, racingService, client)
			return fn(ctx, req, rsp)
		}
	}
}
