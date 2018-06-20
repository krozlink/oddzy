package main

import (
	"gopkg.in/alexcesaro/statsd.v2"
	"os"
)

const (
	statsDefaultAddress = "statsd:9125"
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

func getStats() (*statsdClient, error) {
	add := os.Getenv("STATSD")
	if add == "" {
		add = statsDefaultAddress
	}

	client, err := statsd.New(statsd.Address(add))

	return &statsdClient{c: client}, err
}

func (s *statsdClient) NewTiming() statsTiming {
	return s.c.NewTiming()
}

func (s *statsdClient) Increment(bucket string) {
	s.c.Increment(bucket)
}
