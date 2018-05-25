package main

import (
	"context"
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	"gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &RacingRepository{s.session.Clone()}
}

func (s *service) ListMeetingsByDate(ctx context.Context, req *proto.ListMeetingsByDateRequest, resp *proto.ListMeetingsByDateResponse) error {
	return nil
}

func (s *service) ListRacesByMeetingDate(ctx context.Context, req *proto.ListRacesByMeetingDateRequest, resp *proto.ListRacesByMeetingDateResponse) error {
	return nil
}

func (s *service) AddMeetings(ctx context.Context, req *proto.AddMeetingsRequest, resp *proto.AddMeetingsResponse) error {
	return nil
}

func (s *service) AddRaces(ctx context.Context, req *proto.AddRacesRequest, resp *proto.AddRacesResponse) error {
	return nil
}

func (s *service) UpdateRace(ctx context.Context, req *proto.UpdateRaceRequest, resp *proto.UpdateRaceResponse) error {
	return nil
}
