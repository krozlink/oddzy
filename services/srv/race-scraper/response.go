package main

import (
	"encoding/json"
	"strconv"
)

// RaceSchedule is the format odds.com.au uses to represent a race schedule
type RaceSchedule struct {
	HasResults   bool          `json:"hasResults"`
	RegionGroups []RegionGroup `json:"regionGroups"`
}

// RegionGroup is the format odds.com.au uses to represent a type of races in a region, for example "Australia Greyhounds"
type RegionGroup struct {
	GroupName string    `json:"regionGroup"`
	Meetings  []Meeting `json:"meetings"`
}

// Meeting is the format odds.com.au uses to represent a race meet in their race schedule response
type Meeting struct {
	MeetingName       string  `json:"meetingName"`
	RegionDescription string  `json:"regionDescription"`
	RegionIconURL     string  `json:"regionIconURL"`
	Events            []Event `json:"events"`
}

// Event is the format odds.com.au uses to represent a race in their race schedule response
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

// UnmarshalJSON unmarshalls the event response from odds.com.au to an Event object.
// There are some inconsistencies with the data returned and this manual unmarshal is required
// to handle them. e.g "results" can be null, a string or an int
func (e *Event) UnmarshalJSON(b []byte) error {
	var values map[string]interface{}
	err := json.Unmarshal(b, &values)

	if err != nil {
		return err
	}

	for item, value := range values {
		switch item {
		case "eventId":
			e.EventID = int32(value.(float64))
		case "dateWithYear":
			e.DateWithYear = value.(string)
		case "eventNumber":
			e.EventNumber = int32(value.(float64))
		case "eventUrl":
			e.EventURL = value.(string)
		case "results":
			if value == nil {
				e.Results = ""
			} else if r, ok := value.(string); ok {
				e.Results = r
			} else if r, ok := value.(int); ok {
				e.Results = strconv.Itoa(r)
			}
		case "eventName":
			if value == nil {
				e.EventName = ""
			} else {
				e.EventName = value.(string)
			}
		case "isImminent":
			if value == nil {
				e.IsImminent = false
			} else {
				e.IsImminent = value.(bool)
			}
		case "startTime":
			e.StartTime = int64(value.(float64))
		case "resulted":
			if value == nil {
				e.Resulted = 0
			} else {
				e.Resulted = int32(value.(float64))
			}
		case "isAbandoned":
			if value == nil {
				e.IsAbandoned = 0
			} else {
				e.IsAbandoned = int32(value.(float64))
			}
		}
	}

	return nil
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
	Weight           string      `json:"weight"`
	JockeyWeight     string      `json:"jockeyWeight"`
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
