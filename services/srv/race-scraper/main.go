package main

import (
	proto "github.com/krozlink/oddzy/services/srv/race-scraper/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/consul"
	"log"
)

const (
	racingService = "racing"
)

func main() {
	process := newScrapeProcess()
	registerProcessMonitor(&process)

	process.run()

	<-process.done
}

func registerProcessMonitor(process *scrapeProcess) {
	monitor := newStatusMonitor(process)

	srv := micro.NewService(
		micro.Name("race-scraper"),
		micro.Version("latest"),
	)

	srv.Init()

	proto.RegisterMonitorServiceHandler(srv.Server(), monitor)

	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalln(err)
		}
	}()
}
