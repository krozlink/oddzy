package response

// RaceSchedule is the top level object returned when calling /racing/schedule
type RaceSchedule struct {
	HasRaces bool                  `json:"has_races"`
	RaceDate string                `json:"date"`
	Meetings []RaceScheduleMeeting `json:"meetings"`
	Races    []RaceScheduleRace    `json:"races"`
}

// RaceScheduleMeeting contains the meeting data shown on a race schedule
type RaceScheduleMeeting struct {
	MeetingID      string   `json:"meeting_id"`
	Name           string   `json:"name"`
	Country        string   `json:"country"`
	RaceType       string   `json:"race_type"`
	ScheduledStart int      `json:"scheduled_start"`
	RaceIds        []string `json:"race_ids"`
	LastUpdated    int      `json:"last_update"`
}

// RaceScheduleRace contains the race data shown on a race schedule
type RaceScheduleRace struct {
	RaceID         string `json:"race_id"`
	Name           string `json:"name"`
	Number         int    `json:"number"`
	Status         string `json:"status"`
	Results        string `json:"results"`
	ScheduledStart int    `json:"scheduled_start"`
	LastUpdated    int    `json:"last_update"`
}
