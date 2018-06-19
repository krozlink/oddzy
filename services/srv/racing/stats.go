package main

import (
	"gopkg.in/alexcesaro/statsd.v2"
)

var stats statsClient

type statsdClient struct {
	c statsd.Client
}

type statsClient interface {
	NewTiming() statsd.Timing
	Increment(bucket string)
	Send(bucket string)
}

const (
	statsDefaultAddress = "statsd:9125"
)

func getStats(address string) (*statsd.Client, error) {
	return statsd.New(
		statsd.Address(address),
	)
}

func (s *statsdClient) NewTiming() statsd.Timing {
	return s.c.NewTiming()
}

func (s *statsdClient) Increment(bucket string) {
	s.c.Increment(bucket)
}

func (s *statsdClient) Send(bucket string) {
	s.c.Send(bucket)
}
