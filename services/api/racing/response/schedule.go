package response

// RaceSchedule is the top level object returned when calling /racing/schedule
type RaceSchedule struct {
	HasRaces bool
	RaceDate string
	Meetings []RaceScheduleMeeting `json:"meetings"`
	Races    []RaceScheduleRace    `json:"races"`
}

// RaceScheduleMeeting contains the meeting data shown on a race schedule
type RaceScheduleMeeting struct {
	MeetingID      string
	Name           string
	Country        string
	RaceType       string
	ScheduledStart int
	LastUpdated    int
	RaceIds        []string
}

// RaceScheduleRace contains the race data shown on a race schedule
type RaceScheduleRace struct {
	RaceID         string
	Name           string
	Number         int
	Status         string
	Results        string
	ScheduledStart int
	LastUpdated    int
}
