package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

const (
	meetingDataURL  = "https://www.odds.com.au/api/web/public/Racing/getUpcomingRaces/?sport=%s&date=%s"
	raceDataURL     = "https://www.odds.com.au/api/web/public/Odds/getOddsComparisonCacheable/?eventId=%s&includeTAB=1&includeOdds=1&arrangeOdds=0&betType=FixedWin&includeTote=true&allowGet=true"
	defaultInterval = 1000
)

// Scraper reads racing data from a source
type Scraper interface {
	ScrapeRaceCard(sourceID string) (*RaceCard, error)
	ScrapeRaceCalendar(eventType string, date string) (*RaceCalendar, error)
}

// OddscomauScraper scrapes racing data from odds.com.au
type OddscomauScraper struct {
	http        requestHandler
	lastRequest time.Time
	interval    int
	mux         *sync.Mutex
}

type response struct {
	Value string `json:"r"`
}

// NewOddsScraper returns a new odds.com.au scraper that uses the provided request handler
func NewOddsScraper(h requestHandler) *OddscomauScraper {
	return &OddscomauScraper{
		http:        h,
		lastRequest: time.Time{},
		interval:    defaultInterval,
		mux:         &sync.Mutex{},
	}
}

// ScrapeRaceCalendar reads and parses a racing calendar for the provided event type and date
func (o *OddscomauScraper) ScrapeRaceCalendar(eventType string, date string) (*RaceCalendar, error) {
	throttle(o)

	logContext := log.WithField("method", "ScrapeRaceCalendar").WithField("parameters", fmt.Sprintf("Event Type: %v, Date: %v", eventType, date))

	url := fmt.Sprintf(meetingDataURL, eventType, date)
	logContext.Debugf("Requesting race calendar from %v", url)
	encodedResponse, err := o.http.getResponse(url)
	if err != nil {
		msg := fmt.Sprintf("error retrieving race calendar response - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	if len(encodedResponse) == 0 {
		msg := fmt.Sprintf("No response retrieved")
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	logContext.Debugf("Retrieved encoded response: %s", encodedResponse)

	odds := &response{}
	err = json.Unmarshal(encodedResponse, odds)
	if err != nil {
		msg := fmt.Sprintf("Unable to parse encoded response - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	d, err := decode(odds.Value)
	if err != nil {
		msg := fmt.Sprintf("Unable to decode response - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	logContext.Debugf("Decoded response: %v", string(d))

	calendar := &RaceCalendar{}
	err = json.Unmarshal(d, calendar)

	if err != nil {
		msg := fmt.Sprintf("unable to unmarshal decoded response into race calendar - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	return calendar, nil
}

// ScrapeRaceCard reads and parses a race card for the provided odds.com.au event id
func (o *OddscomauScraper) ScrapeRaceCard(eventID string) (*RaceCard, error) {
	throttle(o)

	logContext := logWithField("method", "ScrapeRaceCard")

	url := fmt.Sprintf(raceDataURL, eventID)
	encodedResponse, err := o.http.getResponse(url)
	if err != nil {
		msg := fmt.Sprintf("error retrieving race card response - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	odds := &response{}
	err = json.Unmarshal(encodedResponse, odds)
	if err != nil {
		msg := fmt.Sprintf("Unable to parse encoded race card response - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	d, err := decode(odds.Value)
	if err != nil {
		msg := fmt.Sprintf("Unable to decode race card response - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	card := &RaceCard{}
	err = json.Unmarshal(d, card)

	if err != nil {
		msg := fmt.Sprintf("unable to unmarshal decoded response into race card - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	return card, nil
}

func throttle(o *OddscomauScraper) {
	o.mux.Lock()
	diff := time.Since(o.lastRequest)
	o.lastRequest = time.Now()
	o.mux.Unlock()

	remainingMS := int(diff.Nanoseconds() / 1000000)
	if remainingMS < o.interval {
		<-time.After(time.Millisecond * time.Duration(o.interval-remainingMS))
	}
}

func decode(response string) ([]byte, error) {

	rot := make([]rune, len(response))
	result := make([]byte, len(response))
	for i, b := range response {
		rot[i] = rot13(b)
	}
	result, err := base64.StdEncoding.DecodeString(string(rot))
	return result, err
}

func rot13(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+13)%26
	}
	return r
}
