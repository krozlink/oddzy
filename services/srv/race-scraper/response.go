package main

// RaceCalendar is the format odds.com.au uses to represent a race calendar
type RaceCalendar struct {
	HasResults   bool `json:"hasResults"`
	RegionGroups []struct {
		RegionGroup string `json:"regionGroup"`
		Meetings    []struct {
			MeetingName       string `json:"meetingName"`
			RegionDescription string `json:"regionDescription"`
			RegionIconURL     string `json:"regionIconURL"`
			Events            []struct {
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
			} `json:"events"`
		} `json:"meetings"`
	} `json:"regionGroups"`
}

// RaceCard is the format odds.com.au uses to represent a race card
type RaceCard struct {
	EventID          int32           `json:"eventId"`
	EventName        string          `json:"eventName"`
	EventNameFull    string          `json:"eventNameFull"`
	EventDescription string          `json:"eventDesc"`
	EventDistance    string          `json:"eventDistance"`
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

type RaceSelection struct {
	Name             string `json:"name"`
	SelectionID      string `json:"selectionId"`
	CompetitorID     string `json:"competitorId"`
	Flucs            string `json:"flucs"`
	ProfileURL       string `json:"profileUrl"`
	Result           string `json:"result"`
	ResultOrdinal    string `json:"resultOrdinal"`
	CompetitorNumber string `json:"competitorNumber"`
	BarrierNumber    string `json:"barrierNumber"`
	ImageURL         string `json:"imageUrl"`
	JockeyName       string `json:"jockeyName"`
	JockeyWeight     string `json:"jockeyWeight"`
	JockeyURL        string `json:"jockeyUrl"`
	Weight           string `json:"weight"`
	Prices           []struct {
		Bookmaker      string  `json:"bookmaker"`
		BookmakerLower string  `json:"bookmakerLower"`
		BetType        string  `json:"betType"`
		OddsKey        string  `json:"oddsKey"`
		HasOdds        bool    `json:"hasOdds"`
		Odds           float32 `json:"odds"`
		IsBest         bool    `json:"isBest"`
	} `json:"prices"`
}
