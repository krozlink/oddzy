package main

import (
	proto "github.com/krozlink/oddzy/services/srv/race-scraper/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/consul"
	"sync"
	"time"
)

const (
	serviceName       = "go.micro.srv.race-scraper"
	serviceVersion    = "0.1.0"
	racingServiceName = "go.micro.srv.racing"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		baseLog = getLog()
		wg.Done()
	}()

	go func() {
		stats = getStats()
		wg.Done()
	}()
	wg.Wait()

	process := newScrapeProcess()
	registerProcessMonitor(&process)

	process.run()

	<-process.done
}

func registerProcessMonitor(process *scrapeProcess) {
	monitor := newStatusMonitor(process)

	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version(serviceVersion),
		micro.RegisterTTL(time.Second*60),
		micro.RegisterInterval(time.Second*10),
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
