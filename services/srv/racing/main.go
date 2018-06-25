package main

import (
	"fmt"
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/consul"
	"sync"
)

const (
	defaultHost    = "db-mongo:27017"
	serviceName    = "racing"
	serviceVersion = "latest"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		baseLog = getLog()
		wg.Done()
	}()

	go func() {
		stats = getStats()
		wg.Done()
	}()
	wg.Wait()

	session := getDBSession()
	defer session.Close()

	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version(serviceVersion),
		micro.WrapHandler(logWrapper),
	)

	srv.Init()

	racing := NewRacingService(
		&RacingRepository{session},
		srv.Server().Options().Broker,
	)

	proto.RegisterRacingServiceHandler(srv.Server(), racing)
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
