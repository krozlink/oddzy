package response

// RaceCard is the top level object returned when calling /racing/racecard
type RaceCard struct {
	RaceID         string              `json:"race_id"`
	Name           string              `json:"name"`
	Number         int32               `json:"number"`
	ScheduledStart int64               `json:"scheduled_start"`
	Results        string              `json:"results"`
	LastUpdated    int64               `json:"last_update"`
	Status         string              `json:"status"`
	Meeting        RaceCardMeeting     `json:"meeting"`
	Selections     []RaceCardSelection `json:"selections"`
}

// RaceCardMeeting contains the meeting data shown on a race card
type RaceCardMeeting struct {
	MeetingID string   `json:"meeting_id"`
	Name      string   `json:"name"`
	Country   string   `json:"country"`
	RaceType  string   `json:"race_type"`
	RaceIds   []string `json:"race_ids"`
}

// RaceCardSelection contains the selection data shown on a race card
type RaceCardSelection struct {
	SelectionID  string `json:"selection_id"`
	Name         string `json:"name"`
	Barrier      int32  `json:"barrier"`
	Number       int32  `json:"number"`
	Jockey       string `json:"jockey"`
	IsScratched  bool   `json:"is_scratched"`
	Weight       string `json:"weight"`
	JockeyWeight string `json:"jockey_weight"`
	ImageURL     string `json:"image_url"`
}
