package main

import (
	proto "github.com/krozlink/oddzy/services/api/racing/proto"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"sync"
)

const (
	serviceName       = "go.micro.api.racing"
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

	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(serviceVersion),
	)

	service.Init(
		micro.WrapHandler(RacingWrapper(service), CorrelationWrapper(service)),
	)

	proto.RegisterRacingHandler(service.Server(), new(Racing))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
