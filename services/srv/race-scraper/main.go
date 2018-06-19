package main

import (
	proto "github.com/krozlink/oddzy/services/srv/race-scraper/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/consul"
)

const (
	racingService = "racing"
	serviceName   = "race-scraper"
)

func main() {
	var err error
	baseLog = getLog()
	stats, err = getStats()
	if err != nil {
		baseLog.Fatalf("Unable to get statsd client - %v", err)
	}

	process := newScrapeProcess()
	registerProcessMonitor(&process)

	process.run()

	<-process.done
}

func registerProcessMonitor(process *scrapeProcess) {
	monitor := newStatusMonitor(process)

	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
	)

	srv.Init()

	proto.RegisterMonitorServiceHandler(srv.Server(), monitor)
	baseLog.Infof("Registered monitor successfully")

	go func() {
		if err := srv.Run(); err != nil {
			baseLog.Fatalln(err)
		}
	}()
}
