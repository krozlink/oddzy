package main

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"sync"
	"testing"
	"time"
)

func getMockScraper(h requestHandler) *OddscomauScraper {
	return &OddscomauScraper{
		http:        h,
		lastRequest: time.Time{},
		interval:    0,
		mux:         sync.Mutex{},
	}
}

type mockHTTPHandler struct {
	response []byte
}

func (h *mockHTTPHandler) getResponse(url string) ([]byte, error) {
	return h.response, nil
}

func TestScrapeRaceCalendar(t *testing.T) {
	// testdata contains an encoded and a decoded calendar
	// this test performs a scrape that uses the encoded calendar and validates that the returned
	// calendar matches the decoded calendar
	encoded, err := ioutil.ReadFile("./testdata/race_calendar_encoded.json")
	if err != nil {
		t.Error(err)
	}
	decoded, err := ioutil.ReadFile("./testdata/race_calendar_decoded.json")
	if err != nil {
		t.Error(err)
	}

	h := &mockHTTPHandler{
		response: encoded,
	}

	scraper := getMockScraper(h)
	cal, err := scraper.ScrapeRaceCalendar("test", "test")
	if err != nil {
		t.Error(err)
	}

	original := &RaceCalendar{}
	err = json.Unmarshal(decoded, original)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(cal, original) {
		t.Error("scraped encoded calendar does not match the decoded calendar")
	}
}

func TestScrapeRaceCalendarHandlesUnexpectedRaceResults(t *testing.T) {

	decoded, err := ioutil.ReadFile("./testdata/race_calendar_unexpected_results.json")
	if err != nil {
		t.Error(err)
	}

	original := &RaceCalendar{}
	err = json.Unmarshal(decoded, original)
	if err != nil {
		t.Error(err)
	}
}

func TestScraperUsesThrottling(t *testing.T) {
	encoded, _ := ioutil.ReadFile("./testdata/race_calendar_encoded.json")

	h := &mockHTTPHandler{
		response: encoded,
	}

	scraper := &OddscomauScraper{
		http:        h,
		lastRequest: time.Time{},
		interval:    500,
		mux:         sync.Mutex{},
	}

	// expect the difference to be at least 500ms as this is the throttled time delay
	start := time.Now()
	scraper.ScrapeRaceCalendar("test", "test")
	scraper.ScrapeRaceCalendar("test", "test")
	end := time.Now()

	diffMS := int(end.Sub(start).Nanoseconds() / 1000000)
	if diffMS < scraper.interval {
		t.Errorf("Expected at least %v elapsed between scrapings but got %v", scraper.interval, diffMS)
	}
}
