package main

import (
	"context"
	proto "github.com/krozlink/oddzy/services/srv/race-scraper/proto"
)

type scrapingMonitor struct {
	process *scrapeProcess
}

func newMonitor(p *scrapeProcess) *scrapingMonitor {
	return &scrapingMonitor{
		process: p,
	}
}

func (s *scrapingMonitor) GetWorkQueue(ctx context.Context, req *proto.GetWorkQueueRequest, resp *proto.GetWorkQueueResponse) error {
	return nil
}

func (s *scrapingMonitor) GetWorkHistory(ctx context.Context, req *proto.GetWorkHistoryRequest, resp *proto.GetWorkHistoryResponse) error {
	return nil
}

func (s *scrapingMonitor) GetStatus(ctx context.Context, req *proto.GetStatusRequest, resp *proto.GetStatusResponse) error {
	resp.Status = s.process.status
	return nil
}
