package main

// RaceCard is the format odds.com.au uses to represent a race card
type RaceCard struct {
	EventID          int32  `json:"eventId"`
	EventName        string `json:"eventName"`
	EventNameFull    string `json:"eventNameFull"`
	EventDescription string `json:"eventDesc"`
	EventDistance    string `json:"eventDistance"`
	StartTime        int64  `json:"startTime"`
	SportID          int32  `json:"sportId"`
	IsRacing         bool   `json:"isRacing"`
	IsGreyhounds     bool   `json:"isGreyhounds"`
	IsHarness        bool   `json:"isHarness"`
	IsHorseRacing    bool   `json:"isHorseRacing"`
	ResultState      string `json:"resultState"`
	Status           string `json:"status"`
	Selections       []struct {
		Name             string `json:"name"`
		SelectionID      string `json:"selectionId"`
		CompetitorID     string `json:"competitorId"`
		Flucs            string `json:"flucs"`
		ProfileURL       string `json:"profileUrl"`
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
	} `json:"selections"`
}
