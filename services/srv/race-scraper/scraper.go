package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

const (
	meetingDataURL          = "https://www.odds.com.au/api/web/public/Racing/getUpcomingRaces/?sport=%s&date=%s"
	raceDataURL             = "https://www.odds.com.au/api/web/public/Odds/getOddsComparisonCacheable/?eventId=%s&includeTAB=1&includeOdds=1&arrangeOdds=0&betType=FixedWin&includeTote=true&allowGet=true"
	defaultInterval         = 1000
	oddsComAuRequestSuccess = "race-scraper.service.odds-request.success"
	oddsComAuRequestFailed  = "race-scraper.service.odds-request.failed"
)

// Scraper reads racing data from a source
type Scraper interface {
	ScrapeRaceCard(ctx context.Context, sourceID string) (*RaceCard, error)
	ScrapeRaceSchedule(ctx context.Context, eventType string, date string) (*RaceSchedule, error)
}

// OddscomauScraper scrapes racing data from odds.com.au
type OddscomauScraper struct {
	http        requestHandler
	lastRequest time.Time
	interval    int
	mux         sync.Mutex
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
		mux:         sync.Mutex{},
	}
}

// ScrapeRaceSchedule reads and parses a racing schedule for the provided event type and date
func (o *OddscomauScraper) ScrapeRaceSchedule(ctx context.Context, eventType string, date string) (*RaceSchedule, error) {
	throttle(o)

	logContext := logWithContext(ctx, "ScrapeRaceSchedule").WithField("parameters", fmt.Sprintf("Event Type: %v, Date: %v", eventType, date))

	url := fmt.Sprintf(meetingDataURL, eventType, date)
	logContext.Debugf("Requesting race schedule from %v", url)
	encodedResponse, err := o.http.getResponse(url)
	if err != nil {
		stats.Increment(oddsComAuRequestFailed)
		msg := fmt.Sprintf("error retrieving race calenscheduledar response - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}
	stats.Increment(oddsComAuRequestSuccess)

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

	schedule := &RaceSchedule{}
	err = json.Unmarshal(d, schedule)

	if err != nil {
		msg := fmt.Sprintf("unable to unmarshal decoded response into race schedule - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	return schedule, nil
}

// ScrapeRaceCard reads and parses a race card for the provided odds.com.au event id
func (o *OddscomauScraper) ScrapeRaceCard(ctx context.Context, eventID string) (*RaceCard, error) {
	throttle(o)

	logContext := logWithContext(ctx, "ScrapeRaceCard")

	url := fmt.Sprintf(raceDataURL, eventID)
	encodedResponse, err := o.http.getResponse(url)
	if err != nil {
		stats.Increment(oddsComAuRequestFailed)
		msg := fmt.Sprintf("error retrieving race card response - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}
	stats.Increment(oddsComAuRequestSuccess)

	odds := &response{}
	err = json.Unmarshal(encodedResponse, odds)
	if err != nil {
		msg := fmt.Sprintf("Unable to parse encoded race card response - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	d, err := decode(odds.Value)
	if err != nil {
		msg := fmt.Sprintf("Unable to decode race card response. Event ID: %v URL: %v Response: %v Error: %v", eventID, url, string(encodedResponse), err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	card := &RaceCard{}
	err = json.Unmarshal(d, card)

	if err != nil {
		msg := fmt.Sprintf("unable to unmarshal decoded response into race card. Event ID: %v URL: %v Response: %v - Error - %v", eventID, url, string(d), err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	return card, nil
}

func throttle(o *OddscomauScraper) {
	o.mux.Lock()
	diff := time.Since(o.lastRequest)
	remainingMS := o.interval - int(diff.Nanoseconds()/1000000)
	o.mux.Unlock()

	if remainingMS > 0 {
		time.Sleep(time.Millisecond * time.Duration(remainingMS))
	}

	o.mux.Lock()
	o.lastRequest = time.Now()
	o.mux.Unlock()
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
