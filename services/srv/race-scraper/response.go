package main

import "encoding/json"

// RaceCalendar is the format odds.com.au uses to represent a race calendar
type RaceCalendar struct {
	HasResults   bool          `json:"hasResults"`
	RegionGroups []RegionGroup `json:"regionGroups"`
}

// RegionGroup is the format odds.com.au uses to represent a type of races in a region, for example "Australia Greyhounds"
type RegionGroup struct {
	GroupName string    `json:"regionGroup"`
	Meetings  []Meeting `json:"meetings"`
}

// Meeting is the format odds.com.au uses to represent a race meet in their race calendar response
type Meeting struct {
	MeetingName       string  `json:"meetingName"`
	RegionDescription string  `json:"regionDescription"`
	RegionIconURL     string  `json:"regionIconURL"`
	Events            []Event `json:"events"`
}

// Event is the format odds.com.au uses to represent a race in their race calendar response
type Event struct {
	EventID      int32  `json:"eventId"`
	EventNumber  int32  `json:"eventNumber"`
	EventName    string `json:"eventName"`
	StartTime    int64  `json:"startTime"`
	DateWithYear string `json:"dateWithYear"`
	IsImminent   bool   `json:"isImminent"`
	IsAbandoned  int32  `json:"isAbandoned"`
	Resulted     int32  `json:"resulted"`
	Results      string `json:"results"`
	EventURL     string `json:"eventUrl"`
}

// RaceCard is the format odds.com.au uses to represent a race card
type RaceCard struct {
	EventID          int32           `json:"eventId"`
	EventName        string          `json:"eventName"`
	EventNameFull    string          `json:"eventNameFull"`
	EventDescription string          `json:"eventDesc"`
	EventDistance    json.Number     `json:"eventDistance,Number"`
	StartTime        int64           `json:"startTime"`
	SportID          int32           `json:"sportId"`
	IsRacing         bool            `json:"isRacing"`
	IsGreyhounds     bool            `json:"isGreyhounds"`
	IsHarness        bool            `json:"isHarness"`
	IsHorseRacing    bool            `json:"isHorseRacing"`
	ResultState      string          `json:"resultState"`
	Status           string          `json:"status"`
	Selections       []RaceSelection `json:"selections"`
}

// RaceSelection is the format odds.com.au uses to represent an entrant in a race
type RaceSelection struct {
	Name             string      `json:"name"`
	SelectionID      string      `json:"selectionId"`
	CompetitorID     string      `json:"competitorId"`
	Result           string      `json:"result"`
	ResultOrdinal    string      `json:"resultOrdinal"`
	CompetitorNumber string      `json:"competitorNumber"`
	BarrierNumber    string      `json:"barrierNumber"`
	ImageURL         string      `json:"imageUrl"`
	JockeyName       string      `json:"jockeyName"`
	Prices           []RacePrice `json:"prices"`
}

// RacePrice is the format odds.com.au uses to represent the price for a market offered by a particular bookmaker
type RacePrice struct {
	Bookmaker      string  `json:"bookmaker"`
	BookmakerLower string  `json:"bookmakerLower"`
	BetType        string  `json:"betType"`
	OddsKey        string  `json:"oddsKey"`
	HasOdds        bool    `json:"hasOdds"`
	Odds           float32 `json:"odds"`
	IsBest         bool    `json:"isBest"`
}
