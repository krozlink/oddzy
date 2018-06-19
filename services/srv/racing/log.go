package main

import (
	"context"
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"
	defaultLog "log"
	"net"
	"os"
)

var log logClient

const (
	loggerEnv     = "ODDZY_LOGGER"
	loggerDefault = "logstash:5000"
	loggerLevel   = logrus.DebugLevel
)

type Level uint32

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

type logClient interface {
	Infof(msg string, args ...interface{})
	Debugf(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
	Errorf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})
	SetLevel()
}

type logger struct {
	c *logrus.Logger
}

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		log.Printf("Server request: %v", req.Method())
		return fn(ctx, req, resp)
	}
}

func getLog() *logrus.Logger {

	env := os.Getenv(loggerEnv)
	if env == "" {
		env = loggerDefault
	}

	l := logrus.New()
	conn, err := net.Dial("tcp", env)
	if err != nil {
		defaultLog.Fatal(err)
	}

	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{
		"service": serviceName,
	}))
	l.Hooks.Add(hook)

	l.SetLevel(loggerLevel)

	return l
}

func logWithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

func addField(l *logrus.Entry, key string, value interface{}) *logrus.Entry {
	return l.WithField(key, value)
}

type logField struct {
	key   string
	value interface{}
}
