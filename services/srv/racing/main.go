package main

import (
	"fmt"
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/consul"
	"log"
	"os"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	log.Printf("Connecting to %s", host)
	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Fatalf("Error connecting to datastore %s: %v", host, err)
	}

	log.Printf("Successfully connected")

	if err = session.Ping(); err != nil {
		log.Fatalf("Error on ping: %v", err)
	}

	srv := micro.NewService(
		micro.Name("oddzy.services.racing"),
		micro.Version("latest"),
	)

	srv.Init()

	repo := &RacingRepository{session}
	b := srv.Server().Options().Broker
	racing := NewRacingService(repo, b)

	proto.RegisterRacingServiceHandler(srv.Server(), racing)
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
