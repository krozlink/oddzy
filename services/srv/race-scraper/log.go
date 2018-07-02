package main

import (
	"context"
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"
	defaultLog "log"
	"net"
	"os"
	"time"
)

var baseLog *logrus.Logger

const (
	loggerEnv           = "ODDZY_LOGGER"
	loggerDefault       = "logstash:5000"
	loggerLevel         = logrus.DebugLevel
	loggerRetryInterval = 10
	loggerMaxAttempts   = 20
)

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		baseLog.Debugf("Server request: %v", req.Method())
		return fn(ctx, req, resp)
	}
}

func getLog() *logrus.Logger {

	env := os.Getenv(loggerEnv)
	if env == "" {
		env = loggerDefault
	}

	l := logrus.New()
	l.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	}

	for count := 1; count <= loggerMaxAttempts; count++ {
		conn, err := net.Dial("tcp", env)
		if err != nil {
			if count == loggerMaxAttempts {
				defaultLog.Print("Unable to connect to logstash - max attempts exceeded")
				defaultLog.Fatal(err)
			}

			defaultLog.Printf("Unable to connect to logstash - attempt %v of %v. Retrying in %v seconds...", count, loggerMaxAttempts, loggerRetryInterval)
			time.Sleep(time.Second * loggerRetryInterval)
			continue
		}

		fmt := logrustash.LogstashFormatter{
			Fields: logrus.Fields{
				"service": serviceName,
			},
			Formatter: l.Formatter,
		}

		hook := logrustash.New(conn, fmt)
		l.Hooks.Add(hook)

		l.SetLevel(loggerLevel)
		break
	}

	defaultLog.Print("Successfully connected to logstash")
	return l
}

func logWithField(key string, value interface{}) *logrus.Entry {
	return baseLog.WithField(key, value)
}

func addField(l *logrus.Entry, key string, value interface{}) *logrus.Entry {
	return l.WithField(key, value)
}

type logField struct {
	key   string
	value interface{}
}
