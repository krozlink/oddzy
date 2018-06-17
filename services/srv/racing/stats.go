package main

import (
	"gopkg.in/alexcesaro/statsd.v2"
)

var stats *statsd.Client

const (
	statsDefaultAddress = "statsd:9125"
)

func getStats(address string) (*statsd.Client, error) {
	return statsd.New(
		statsd.Address(address),
	)
}
