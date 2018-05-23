package main

type raceCard struct {
	eventId       int32
	eventName     string
	eventNameFull string
	eventDesc     string
	eventDistance string
	startTime     int64
	sportId       int32
	isRacing      bool
	isGreyhounds  bool
	isHarness     bool
	isHorseRacing bool
	resultState   string
	status        string
	selections    []struct {
		name             string
		selectionId      string
		competitorId     string
		flucs            string
		profileUrl       string
		competitorNumber string
		barrierNumber    string
		imageUrl         string
		jockeyName       string
		jockeyWeight     string
		jockeyUrl        string
		weight           string
		prices           []struct {
			bookmaker string
			bookmakerLower string
			betType string
			oddsKey string
			hasOdds bool
			odds float32
			isBest bool
		}
	}
}
