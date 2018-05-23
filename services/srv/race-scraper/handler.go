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

func (s *service) GetWorkHistory(ctx context.Context, req *proto.GetWorkHistoryRequest, resp *proto.GetWorkHistoryResponse) error {
	return nil
}

func (s *service) GetStatus(ctx context.Context, req *proto.GetStatusRequest, resp *proto.GetStatusResponse) error {
	resp.Status = status
	return nil
}
