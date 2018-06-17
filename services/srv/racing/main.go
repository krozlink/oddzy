package main

import (
	"fmt"
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/consul"
	"os"
)

const (
	defaultHost    = "db-mongo:27017"
	serviceName    = "racing"
	serviceVersion = "latest"
)

func main() {
	var err error
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	log = getLog()

	stats, err = getStats(os.Getenv("STATSD"))
	if err != nil {
		log.Fatalf("Unable to get statsd client - %v", err)
	}

	log.Infof("Connecting to %s", host)
	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Fatalf("Error connecting to datastore %s: %v", host, err)
	}

	log.Info("Successfully connected")

	if err = session.Ping(); err != nil {
		log.Fatalf("Error on ping: %v", err)
	}

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
