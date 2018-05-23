package main

type raceCalendar struct {
	hasResults   bool
	regionGroups []struct {
		regionGroup string
		meetings    []struct {
			meetingName       string
			regionDescription string
			regionIconURL     string
			events            []struct {
				eventId      int32
				eventNumber  int32
				eventName    string
				startTime    int64
				dateWithYear string
				isImminent   bool
				isAbandoned  int32
				resulted     int32
				results      string
				eventUrl     string
			}
		}
	}
}
