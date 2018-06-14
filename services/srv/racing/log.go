package main

import (
	"context"
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"
	"net"
)

var log *logrus.Logger

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		log.Printf("Server request: %v", req.Method())
		return fn(ctx, req, resp)
	}
}

func getLog() *logrus.Logger {
	l := logrus.New()
	conn, err := net.Dial("tcp", "logstash")
	if err != nil {
		log.Fatal(err)
	}

	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{
		"type": "racing",
	}))
	l.Hooks.Add(hook)

	return l
}
