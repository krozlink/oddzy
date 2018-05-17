package main

import (
	"context"
	proto "github.com/krozlink/oddzy/services/venue/proto"
	"gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VenueRepository{s.session.Clone()}
}

func (s *service) Get(ctx context.Context, req *proto.GetRequest, resp *proto.GetResponse) error {
	return nil
}

func (s *service) List(ctx context.Context, req *proto.ListRequest, resp *proto.ListResponse) error {
	repo := s.GetRepo()
	defer repo.Close()

	venues, err := repo.List()
	resp.Venues = venues
	return err
}

func (s *service) Add(ctx context.Context, req *proto.AddRequest, resp *proto.AddResponse) error {
	repo := s.GetRepo()
	defer repo.Close()

	venue := &proto.Venue{
		Id:       req.Id,
		Category: req.Category,
		Country:  req.Country,
		Name:     req.Name,
		State:    req.State,
	}

	err := repo.Add(venue)
	if err != nil {
		resp.Created = false
		return err
	}

	resp.Created = true
	resp.Venue = venue

	return nil
}

func (s *service) Delete(ctx context.Context, req *proto.DeleteRequest, resp *proto.DeleteResponse) error {
	return nil
}
