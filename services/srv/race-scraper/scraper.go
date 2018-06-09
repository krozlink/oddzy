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
	defaultInterval = 1
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
	interval    float64
	mux         *sync.Mutex
}

type response struct {
	r string
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

	url := fmt.Sprintf(meetingDataURL, eventType, date)
	encodedResponse, err := o.http.getResponse(url)
	if err != nil {
		return nil, fmt.Errorf("error retrieving race calendar response - %v", err)
	}

	odds := response{}
	json.Unmarshal(encodedResponse, odds)
	calendar := &RaceCalendar{}
	err = json.Unmarshal([]byte(odds.r), calendar)

	if err != nil {
		return nil, fmt.Errorf("unable to decode response into race calendar - %v", err)
	}

	return calendar, nil
}

// ScrapeRaceCard reads and parses a race card for the provided odds.com.au event id
func (o *OddscomauScraper) ScrapeRaceCard(eventID string) (*RaceCard, error) {
	throttle(o)
	return nil, nil
}

func throttle(o *OddscomauScraper) {
	o.mux.Lock()
	diff := time.Since(o.lastRequest)
	o.lastRequest = time.Now()
	o.mux.Unlock()

	if diff.Seconds() < o.interval {
		<-time.After(diff)
	}
}

func decode(response []byte) []byte {

	rot := make([]byte, len(response))
	result := make([]byte, len(response))
	for i, b := range response {
		rot[i] = b
	}
	base64.StdEncoding.Decode(result, rot)
	return result
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
