package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"sync"
	"testing"
	"time"
)

func getMockScraper(h requestHandler) *OddsScraper {
	return &OddsScraper{
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

func TestScrapeRaceSchedule(t *testing.T) {
	// testdata contains an encoded and a decoded schedule
	// this test performs a scrape that uses the encoded schedule and validates that the returned
	// schedule matches the decoded schedule
	encoded, err := ioutil.ReadFile("./testdata/race_schedule_encoded.json")
	if err != nil {
		t.Error(err)
	}
	decoded, err := ioutil.ReadFile("./testdata/race_schedule_decoded.json")
	if err != nil {
		t.Error(err)
	}

	h := &mockHTTPHandler{
		response: encoded,
	}

	ctx := context.Background()
	scraper := getMockScraper(h)
	cal, err := scraper.ScrapeRaceSchedule(ctx, "test", "test")
	if err != nil {
		t.Error(err)
	}

	original := &RaceSchedule{}
	err = json.Unmarshal(decoded, original)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(cal, original) {
		t.Error("scraped encoded schedule does not match the decoded schedule")
	}
}

func TestScrapeRaceScheduleHandlesUnexpectedRaceResults(t *testing.T) {

	decoded, err := ioutil.ReadFile("./testdata/race_schedule_unexpected_results.json")
	if err != nil {
		t.Error(err)
	}

	original := &RaceSchedule{}
	err = json.Unmarshal(decoded, original)
	if err != nil {
		t.Error(err)
	}
}

func TestScraperUsesThrottling(t *testing.T) {
	encoded, _ := ioutil.ReadFile("./testdata/race_schedule_encoded.json")

	h := &mockHTTPHandler{
		response: encoded,
	}

	scraper := &OddsScraper{
		http:        h,
		lastRequest: time.Time{},
		interval:    500,
		mux:         sync.Mutex{},
	}

	ctx := context.Background()

	// expect the difference to be at least 500ms as this is the throttled time delay
	start := time.Now()
	scraper.ScrapeRaceSchedule(ctx, "test", "test")
	scraper.ScrapeRaceSchedule(ctx, "test", "test")
	end := time.Now()

	diffMS := int(end.Sub(start).Nanoseconds() / 1000000)
	if diffMS < scraper.interval {
		t.Errorf("Expected at least %v elapsed between scrapings but got %v", scraper.interval, diffMS)
	}
}
