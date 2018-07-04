package main

import (
	"context"
	racing "github.com/krozlink/oddzy/services/srv/racing/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

type serviceKey struct{}

// RacingFromContext retrieves the racing client from the Context
func RacingFromContext(ctx context.Context) (racing.RacingService, bool) {
	c, ok := ctx.Value(serviceKey{}).(racing.RacingService)
	return c, ok
}

// RacingWrapper returns a wrapper for the RacingClient
func RacingWrapper(service micro.Service) server.HandlerWrapper {
	client := racing.NewRacingService("go.micro.srv.racing", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, serviceKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
