package response

type RaceCard struct {
	RaceID         string      `json:"race_id"`
	Name           string      `json:"name"`
	Number         string      `json:"number"`
	ScheduledStart int         `json:"scheduled_start"`
	Results        string      `json:"results"`
	LastUpdated    int         `json:"last_update"`
	Status         string      `json:"status"`
	Meeting        Meeting     `json:"meeting"`
	Selections     []Selection `json:"selections"`
}

type Meeting struct {
	MeetingID string   `json:"meeting_id"`
	Name      string   `json:"name"`
	RaceIds   []string `json:"race_ids"`
}

type Selection struct {
	SelectionID string `json:"selection_id"`
	Name        string `json:"name"`
	Barrier     int    `json:"barrier"`
	Number      int    `json:"number"`
	Jockey      string `json:"jockey"`
	IsScratched bool   `json:"is_scratched"`
}
