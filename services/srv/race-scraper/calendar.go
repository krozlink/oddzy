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
