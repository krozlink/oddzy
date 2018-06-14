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

	url := fmt.Sprintf(meetingDataURL, eventType, date)
	encodedResponse, err := o.http.getResponse(url)
	if err != nil {
		return nil, fmt.Errorf("error retrieving race calendar response - %v", err)
	}

	odds := &response{}
	err = json.Unmarshal(encodedResponse, odds)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse encoded response - %v", err)
	}

	d, err := decode(odds.Value)
	if err != nil {
		return nil, fmt.Errorf("Unable to decode response - %v", err)
	}

	calendar := &RaceCalendar{}
	err = json.Unmarshal(d, calendar)

	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal decoded response into race calendar - %v", err)
	}

	return calendar, nil
}

// ScrapeRaceCard reads and parses a race card for the provided odds.com.au event id
func (o *OddscomauScraper) ScrapeRaceCard(eventID string) (*RaceCard, error) {
	throttle(o)

	url := fmt.Sprintf(raceDataURL, eventID)
	encodedResponse, err := o.http.getResponse(url)
	if err != nil {
		return nil, fmt.Errorf("error retrieving race card response - %v", err)
	}

	odds := &response{}
	err = json.Unmarshal(encodedResponse, odds)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse encoded race card response - %v", err)
	}

	d, err := decode(odds.Value)
	if err != nil {
		return nil, fmt.Errorf("Unable to decode race card response - %v", err)
	}

	card := &RaceCard{}
	err = json.Unmarshal(d, card)

	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal decoded response into race card - %v", err)
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
