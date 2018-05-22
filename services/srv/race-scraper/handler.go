package main

import (
	"context"
	proto "github.com/krozlink/oddzy/services/srv/race-scraper/proto"
)

type service struct {
	
}

func (s *service) GetWorkQueue(ctx context.Context, req *proto.GetWorkQueueRequest, resp *proto.GetWorkQueueResponse) error {
	return nil
}
