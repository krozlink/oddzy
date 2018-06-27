package main

import (
	"fmt"
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/consul"
	"sync"
	"time"
)

const (
	defaultHost    = "db-mongo:27017"
	serviceName    = "go.micro.srv.racing"
	serviceVersion = "0.1.0"
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

	session := getDBSession()
	defer session.Close()

	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version(serviceVersion),
		micro.WrapHandler(logWrapper),
		micro.RegisterTTL(time.Second*60),
		micro.RegisterInterval(time.Second*10),
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
