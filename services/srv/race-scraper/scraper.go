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
	meetingDataURL                  = "uggcf://jjj.bqqf.pbz.nh/ncv/jro/choyvp/Enpvat/trgHcpbzvatEnprf/?fcbeg=%f&qngr=%f"
	raceDataURL                     = "uggcf://jjj.bqqf.pbz.nh/ncv/jro/choyvp/Bqqf/trgBqqfPbzcnevfbaPnpurnoyr/?riragVq=%f&vapyhqrGNO=1&vapyhqrBqqf=1&neenatrBqqf=0&orgGlcr=SvkrqJva&vapyhqrGbgr=gehr&nyybjTrg=gehr"
	defaultInterval                 = 1000
	oddsComAuScheduleRequestSuccess = "oddzy.race-scraper.service.scraper.schedule.success"
	oddsComAuScheduleRequestFailed  = "oddzy.race-scraper.service.scraper.schedule.failed"
	oddsComAuRacecardRequestSuccess = "oddzy.race-scraper.service.scraper.racecard.success"
	oddsComAuRacecardRequestFailed  = "oddzy.race-scraper.service.scraper.racecard.failed"
)

// Scraper reads racing data from a source
type Scraper interface {
	ScrapeRaceCard(ctx context.Context, sourceID string) (*RaceCard, error)
	ScrapeRaceSchedule(ctx context.Context, eventType string, date string) (*RaceSchedule, error)
}

// OddsScraper scrapes racing data from odds
type OddsScraper struct {
	http        requestHandler
	lastRequest time.Time
	interval    int
	mux         sync.Mutex
}

type response struct {
	Value string `json:"r"`
}

// NewOddsScraper returns a new odds scraper that uses the provided request handler
func NewOddsScraper(h requestHandler) *OddsScraper {
	return &OddsScraper{
		http:        h,
		lastRequest: time.Time{},
		interval:    defaultInterval,
		mux:         sync.Mutex{},
	}
}

// ScrapeRaceSchedule reads and parses a racing schedule for the provided event type and date
func (o *OddsScraper) ScrapeRaceSchedule(ctx context.Context, eventType string, date string) (*RaceSchedule, error) {
	throttle(o)

	logContext := logWithContext(ctx, "ScrapeRaceSchedule").WithField("parameters", fmt.Sprintf("Event Type: %v, Date: %v", eventType, date))

	url := fmt.Sprintf(rot13String(meetingDataURL), eventType, date)
	logContext.Debugf("Requesting race schedule from %v", url)
	encodedResponse, err := o.http.getResponse(url)
	if err != nil {
		stats.Increment(oddsComAuScheduleRequestFailed)
		msg := fmt.Sprintf("error retrieving race calenscheduledar response - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}
	stats.Increment(oddsComAuScheduleRequestSuccess)

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

// ScrapeRaceCard reads and parses a race card for the provided odds event id
func (o *OddsScraper) ScrapeRaceCard(ctx context.Context, eventID string) (*RaceCard, error) {
	throttle(o)

	logContext := logWithContext(ctx, "ScrapeRaceCard")

	url := fmt.Sprintf(rot13String(raceDataURL), eventID)
	encodedResponse, err := o.http.getResponse(url)
	if err != nil {
		stats.Increment(oddsComAuRacecardRequestFailed)
		msg := fmt.Sprintf("error retrieving race card response - %v", err)
		logContext.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}
	stats.Increment(oddsComAuRacecardRequestSuccess)

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

func throttle(o *OddsScraper) {
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

func rot13String(s string) string {
	result := make([]rune, len(s))
	for i, r := range s {
		result[i] = rot13(r)
	}
	return string(result)
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
