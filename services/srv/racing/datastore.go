package main

import (
	"gopkg.in/mgo.v2"
	"os"
	"time"
)

const (
	sessionRetryInterval = 5
	sessionMaxAttempts   = 10
)

// CreateSession creates a mongodb session with the provided host
func CreateSession(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}

func getDBSession() *mgo.Session {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	baseLog.Infof("Connecting to %s", host)

	var session *mgo.Session
	var err error
	for count := 1; count <= sessionMaxAttempts; count++ {
		session, err = CreateSession(host)
		if err != nil {
			if count == sessionMaxAttempts {
				baseLog.Print("Unable to connect to datastore - max attempts exceeded")
				baseLog.Fatal(err)
			}

			baseLog.Printf("Unable to connect to datastore - attempt %v of %v. Retrying in %v seconds...", count, sessionMaxAttempts, sessionRetryInterval)
			time.Sleep(time.Second * sessionRetryInterval)
			continue
		}

		break
	}

	if err = session.Ping(); err != nil {
		baseLog.Fatalf("Error on ping: %v", err)
	}

	baseLog.Info("Successfully connected to datastore")
	return session
}
