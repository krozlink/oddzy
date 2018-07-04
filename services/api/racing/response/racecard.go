package response

// RaceCard is the top level object returned when calling /racing/racecard
type RaceCard struct {
	RaceID         string              `json:"race_id"`
	Name           string              `json:"name"`
	Number         string              `json:"number"`
	ScheduledStart int                 `json:"scheduled_start"`
	Results        string              `json:"results"`
	LastUpdated    int                 `json:"last_update"`
	Status         string              `json:"status"`
	Meeting        RaceCardMeeting     `json:"meeting"`
	Selections     []RaceCardSelection `json:"selections"`
}

// RaceCardMeeting contains the meeting data shown on a race card
type RaceCardMeeting struct {
	MeetingID string   `json:"meeting_id"`
	Name      string   `json:"name"`
	RaceIds   []string `json:"race_ids"`
}

// RaceCardSelection contains the selection data shown on a race card
type RaceCardSelection struct {
	SelectionID string `json:"selection_id"`
	Name        string `json:"name"`
	Barrier     int    `json:"barrier"`
	Number      int    `json:"number"`
	Jockey      string `json:"jockey"`
	IsScratched bool   `json:"is_scratched"`
}
