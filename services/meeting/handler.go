package main

import (
	"context"
	proto "github.com/krozlink/oddzy/services/race/proto"
)

type service struct{}

func (s *service) Get(ctx context.Context, req *proto.GetRequest, resp *proto.GetResponse) error {
	return nil
}

func (s *service) List(ctx context.Context, req *proto.ListRequest, resp *proto.ListResponse) error {
	return nil
}

func (s *service) Add(ctx context.Context, req *proto.AddRequest, resp *proto.AddResponse) error {
	return nil
}

func (s *service) Delete(ctx context.Context, req *proto.DeleteRequest, resp *proto.DeleteResponse) error {
	return nil
}
