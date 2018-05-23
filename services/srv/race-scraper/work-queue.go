package main

import (
	proto "github.com/krozlink/oddzy/services/srv/race-scraper/proto"
	"time"
)

var (
	workQueue []proto.ScrapeHistoryItem
)

func addURLToQueue(url string) {
	item := proto.ScrapeHistoryItem{
		Timestamp: time.Now().Unix(),
		Url:       url,
	}
	workQueue = append(workQueue, item)
}
