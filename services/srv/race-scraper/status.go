package main

import (
	"context"
	proto "github.com/krozlink/oddzy/services/srv/race-scraper/proto"
)

type statusMonitor struct {
	process *scrapeProcess
}

func newStatusMonitor(p *scrapeProcess) *statusMonitor {
	return &statusMonitor{
		process: p,
	}
}

func (s *statusMonitor) GetWorkQueue(ctx context.Context, req *proto.GetWorkQueueRequest, resp *proto.GetWorkQueueResponse) error {
	return nil
}

func (s *statusMonitor) GetWorkHistory(ctx context.Context, req *proto.GetWorkHistoryRequest, resp *proto.GetWorkHistoryResponse) error {
	return nil
}

func (s *statusMonitor) GetStatus(ctx context.Context, req *proto.GetStatusRequest, resp *proto.GetStatusResponse) error {
	resp.Status = s.process.status
	return nil
}
