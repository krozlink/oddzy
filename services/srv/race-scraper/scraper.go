package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

const (
	meetingDataURL = "https://www.odds.com.au/api/web/public/Racing/getUpcomingRaces/?sport=%s&date=%s"
	raceDataURL    = "https://www.odds.com.au/api/web/public/Odds/getOddsComparisonCacheable/?eventId=%s&includeTAB=1&includeOdds=1&arrangeOdds=0&betType=FixedWin&includeTote=true&allowGet=true"
)

// Scraper reads racing data from a source
type Scraper interface {
	ScrapeRaceCard(sourceID string) (*RaceCard, error)
	ScrapeRaceCalendar(eventType string, date string) (*RaceCalendar, error)
}

// OddscomauScraper scrapes racing data from odds.com.au
type OddscomauScraper struct {
	http requestHandler
}

type response struct {
	r string
}

// NewOddsScraper returns a new odds.com.au scraper that uses the provided request handler
func NewOddsScraper(h requestHandler) *OddscomauScraper {
	return &OddscomauScraper{
		http: h,
	}
}

// ScrapeRaceCalendar reads and parses a racing calendar for the provided event type and date
func (o *OddscomauScraper) ScrapeRaceCalendar(eventType string, date string) (*RaceCalendar, error) {
	url := fmt.Sprintf(meetingDataURL, eventType, date)
	encodedResponse, err := o.http.getResponse(url)
	if err != nil {
		fmt.Printf("Error retrieving race calendar response - %v \n", err)
		return nil, err
	}

	odds := response{}
	json.Unmarshal(encodedResponse, odds)
	calendar := &RaceCalendar{}
	err = json.Unmarshal([]byte(odds.r), calendar)

	if err != nil {
		fmt.Println("Unable to decode response into race calendar")
		fmt.Println(odds.r)
		return nil, err
	}

	return calendar, nil
}

// ScrapeRaceCard reads and parses a race card for the provided odds.com.au event id
func (o *OddscomauScraper) ScrapeRaceCard(eventID string) (*RaceCard, error) {
	return nil, nil
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
