package main

import (
	"gopkg.in/alexcesaro/statsd.v2"
	"log"
	"os"
	"time"
)

const (
	statsDefaultAddress = "statsd:9125"
	statsRetryInterval  = 5
	statsMaxAttempts    = 10
)

var stats statsClient

type statsClient interface {
	Increment(stat string)
	NewTiming() statsTiming
}

type statsTiming interface {
	Send(bucket string)
}

type statsdClient struct {
	c *statsd.Client
}

func getStats() *statsdClient {
	add := os.Getenv("STATSD")
	if add == "" {
		add = statsDefaultAddress
	}

	var client *statsd.Client
	var err error
	for count := 1; count <= statsMaxAttempts; count++ {
		client, err = statsd.New(statsd.Address(add))
		if err != nil {
			if count == statsMaxAttempts {
				log.Print("Unable to connect to statsd - max attempts exceeded")
				log.Fatal(err)
			}

			log.Printf("Unable to connect to statsd - attempt %v of %v. Retrying in %v seconds...", count, statsMaxAttempts, statsRetryInterval)
			time.Sleep(time.Second * statsRetryInterval)
			continue
		}
		break
	}

	log.Print("Successfully connected to statsd")
	return &statsdClient{c: client}
}

func (s *statsdClient) NewTiming() statsTiming {
	return s.c.NewTiming()
}

func (s *statsdClient) Increment(bucket string) {
	s.c.Increment(bucket)
}
