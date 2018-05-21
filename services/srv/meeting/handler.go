package main

import (
	"context"
	proto "github.com/krozlink/oddzy/services/srv/meeting/proto"
)

type service struct{}

func (s *service) Get(ctx context.Context, req *proto.GetRequest, resp *proto.GetResponse) error {
	return nil
}

func (s *service) GetBySourceID(ctx context.Context, req *proto.GetBySourceIDRequest, resp *proto.GetResponse) error {
	return nil
}

func (s *service) List(ctx context.Context, req *proto.ListRequest, resp *proto.ListResponse) error {
	return nil
}

func (s *service) ListByDate(ctx context.Context, req *proto.ListByDateRequest, resp *proto.ListResponse) error {
	return nil
}

func (s *service) Add(ctx context.Context, req *proto.AddRequest, resp *proto.AddResponse) error {
	return nil
}

func (s *service) AddRange(ctx context.Context, req *proto.AddRangeRequest, resp *proto.AddRangeResponse) error {
	return nil
}

func (s *service) Delete(ctx context.Context, req *proto.DeleteRequest, resp *proto.DeleteResponse) error {
	resp.Deleted = true
	return nil
}
