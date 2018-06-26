// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/racing.proto

/*
Package go_micro_srv_racing is a generated protocol buffer package.

It is generated from these files:
	proto/racing.proto

It has these top-level messages:
	RaceUpdatedMessage
	ListMeetingsByDateRequest
	ListMeetingsByDateResponse
	ListRacesByMeetingDateRequest
	ListRacesByMeetingDateResponse
	AddRacesRequest
	AddRacesResponse
	AddMeetingsRequest
	AddMeetingsResponse
	UpdateRaceRequest
	UpdateRaceResponse
	GetNextRaceRequest
	GetNextRaceResponse
	Meeting
	Competitor
	Selection
	Race
*/
package go_micro_srv_racing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for RacingService service

type RacingService interface {
	ListMeetingsByDate(ctx context.Context, in *ListMeetingsByDateRequest, opts ...client.CallOption) (*ListMeetingsByDateResponse, error)
	ListRacesByMeetingDate(ctx context.Context, in *ListRacesByMeetingDateRequest, opts ...client.CallOption) (*ListRacesByMeetingDateResponse, error)
	AddMeetings(ctx context.Context, in *AddMeetingsRequest, opts ...client.CallOption) (*AddMeetingsResponse, error)
	AddRaces(ctx context.Context, in *AddRacesRequest, opts ...client.CallOption) (*AddRacesResponse, error)
	UpdateRace(ctx context.Context, in *UpdateRaceRequest, opts ...client.CallOption) (*UpdateRaceResponse, error)
	GetNextRace(ctx context.Context, in *GetNextRaceRequest, opts ...client.CallOption) (*GetNextRaceResponse, error)
}

type racingService struct {
	c    client.Client
	name string
}

func NewRacingService(name string, c client.Client) RacingService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.racing"
	}
	return &racingService{
		c:    c,
		name: name,
	}
}

func (c *racingService) ListMeetingsByDate(ctx context.Context, in *ListMeetingsByDateRequest, opts ...client.CallOption) (*ListMeetingsByDateResponse, error) {
	req := c.c.NewRequest(c.name, "RacingService.ListMeetingsByDate", in)
	out := new(ListMeetingsByDateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racingService) ListRacesByMeetingDate(ctx context.Context, in *ListRacesByMeetingDateRequest, opts ...client.CallOption) (*ListRacesByMeetingDateResponse, error) {
	req := c.c.NewRequest(c.name, "RacingService.ListRacesByMeetingDate", in)
	out := new(ListRacesByMeetingDateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racingService) AddMeetings(ctx context.Context, in *AddMeetingsRequest, opts ...client.CallOption) (*AddMeetingsResponse, error) {
	req := c.c.NewRequest(c.name, "RacingService.AddMeetings", in)
	out := new(AddMeetingsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racingService) AddRaces(ctx context.Context, in *AddRacesRequest, opts ...client.CallOption) (*AddRacesResponse, error) {
	req := c.c.NewRequest(c.name, "RacingService.AddRaces", in)
	out := new(AddRacesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racingService) UpdateRace(ctx context.Context, in *UpdateRaceRequest, opts ...client.CallOption) (*UpdateRaceResponse, error) {
	req := c.c.NewRequest(c.name, "RacingService.UpdateRace", in)
	out := new(UpdateRaceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racingService) GetNextRace(ctx context.Context, in *GetNextRaceRequest, opts ...client.CallOption) (*GetNextRaceResponse, error) {
	req := c.c.NewRequest(c.name, "RacingService.GetNextRace", in)
	out := new(GetNextRaceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RacingService service

type RacingServiceHandler interface {
	ListMeetingsByDate(context.Context, *ListMeetingsByDateRequest, *ListMeetingsByDateResponse) error
	ListRacesByMeetingDate(context.Context, *ListRacesByMeetingDateRequest, *ListRacesByMeetingDateResponse) error
	AddMeetings(context.Context, *AddMeetingsRequest, *AddMeetingsResponse) error
	AddRaces(context.Context, *AddRacesRequest, *AddRacesResponse) error
	UpdateRace(context.Context, *UpdateRaceRequest, *UpdateRaceResponse) error
	GetNextRace(context.Context, *GetNextRaceRequest, *GetNextRaceResponse) error
}

func RegisterRacingServiceHandler(s server.Server, hdlr RacingServiceHandler, opts ...server.HandlerOption) {
	type racingService interface {
		ListMeetingsByDate(ctx context.Context, in *ListMeetingsByDateRequest, out *ListMeetingsByDateResponse) error
		ListRacesByMeetingDate(ctx context.Context, in *ListRacesByMeetingDateRequest, out *ListRacesByMeetingDateResponse) error
		AddMeetings(ctx context.Context, in *AddMeetingsRequest, out *AddMeetingsResponse) error
		AddRaces(ctx context.Context, in *AddRacesRequest, out *AddRacesResponse) error
		UpdateRace(ctx context.Context, in *UpdateRaceRequest, out *UpdateRaceResponse) error
		GetNextRace(ctx context.Context, in *GetNextRaceRequest, out *GetNextRaceResponse) error
	}
	type RacingService struct {
		racingService
	}
	h := &racingServiceHandler{hdlr}
	s.Handle(s.NewHandler(&RacingService{h}, opts...))
}

type racingServiceHandler struct {
	RacingServiceHandler
}

func (h *racingServiceHandler) ListMeetingsByDate(ctx context.Context, in *ListMeetingsByDateRequest, out *ListMeetingsByDateResponse) error {
	return h.RacingServiceHandler.ListMeetingsByDate(ctx, in, out)
}

func (h *racingServiceHandler) ListRacesByMeetingDate(ctx context.Context, in *ListRacesByMeetingDateRequest, out *ListRacesByMeetingDateResponse) error {
	return h.RacingServiceHandler.ListRacesByMeetingDate(ctx, in, out)
}

func (h *racingServiceHandler) AddMeetings(ctx context.Context, in *AddMeetingsRequest, out *AddMeetingsResponse) error {
	return h.RacingServiceHandler.AddMeetings(ctx, in, out)
}

func (h *racingServiceHandler) AddRaces(ctx context.Context, in *AddRacesRequest, out *AddRacesResponse) error {
	return h.RacingServiceHandler.AddRaces(ctx, in, out)
}

func (h *racingServiceHandler) UpdateRace(ctx context.Context, in *UpdateRaceRequest, out *UpdateRaceResponse) error {
	return h.RacingServiceHandler.UpdateRace(ctx, in, out)
}

func (h *racingServiceHandler) GetNextRace(ctx context.Context, in *GetNextRaceRequest, out *GetNextRaceResponse) error {
	return h.RacingServiceHandler.GetNextRace(ctx, in, out)
}
