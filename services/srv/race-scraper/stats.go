package main

import (
	"gopkg.in/alexcesaro/statsd.v2"
	"os"
)

var stats statsClient

type statsdClient struct {
	c statsd.Client
}

type statsClient interface {
	NewTiming() statsd.Timing
	Increment(bucket string)
}

const (
	statsDefaultAddress = "statsd:9125"
)

func getStats() (*statsd.Client, error) {
	add := os.Getenv("STATSD")
	if add == "" {
		add = statsDefaultAddress
	}
	return statsd.New(
		statsd.Address(add),
	)
}

func (s *statsdClient) NewTiming() statsd.Timing {
	return s.c.NewTiming()
}

func (s *statsdClient) Increment(bucket string) {
	s.c.Increment(bucket)
}
