package main

import (
	"context"
	proto "github.com/krozlink/oddzy/services/venue/proto"
)

type service struct{}

func (s *service) Get(ctx context.Context, req *proto.GetRequest, resp *proto.GetResponse) error {
	return nil
}

func (s *service) List(ctx context.Context, req *proto.ListRequest, resp *proto.ListResponse) error {
	resp.Venues = []*proto.Venue{
		&proto.Venue{
			Category: "cat1234",
			Country:  "AUS",
			Id:       "234523432",
			Name:     "Test Venue",
		}}
	return nil
}

func (s *service) Add(ctx context.Context, req *proto.AddRequest, resp *proto.AddResponse) error {
	return nil
}

func (s *service) Delete(ctx context.Context, req *proto.DeleteRequest, resp *proto.DeleteResponse) error {
	return nil
}
