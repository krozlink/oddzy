package main

import (
	model "github.com/krozlink/oddzy/services/srv/racing/model"
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	dbName               = "oddzy"
	raceCollection       = "races"
	meetingCollection    = "meetings"
	selectionCollection  = "selections"
	competitorCollection = "competitors"
)

// Repository interface used to access racing data
type Repository interface {
	AddMeetings(meetings []*proto.Meeting) error
	GetMeeting(meetingID string) (*proto.Meeting, error)
	ListMeetingsByDate(start, end int64) ([]*proto.Meeting, error)
	ListRacesByMeetingDate(start, end int64) ([]*proto.Race, error)

	AddRaces(races []*proto.Race) error
	GetRace(raceID string) (*proto.Race, error)
	ListRacesByMeetingID(meetingID string) ([]*proto.Race, error)
	UpdateRace(race *proto.RaceUpdatedMessage) error

	AddSelections(races []*proto.Selection) error
	ListSelectionsByRaceID(raceID string) ([]*proto.Selection, error)
	UpdateSelection(selection *proto.Selection) error

	Close()
	NewSession() Repository
}

// RacingRepository implements the Repository interface to access racing data
type RacingRepository struct {
	session *mgo.Session
}

// NewSession returns an instance of the repository with a new session
func (repo *RacingRepository) NewSession() Repository {
	return &RacingRepository{repo.session.Clone()}
}

// ListMeetingsByDate will return all meetings between the provided start and end dates
func (repo *RacingRepository) ListMeetingsByDate(start, end int64) ([]*proto.Meeting, error) {
	var results []*model.Meeting

	err := repo.collection(meetingCollection).Find(
		bson.M{
			"scheduled_start": bson.M{
				"$gte": time.Unix(start, 0),
				"$lte": time.Unix(end, 0),
			},
		},
	).All(&results)

	if err != nil {
		return nil, err
	}

	meetings := model.MeetingModelToProtoCollection(results)
	return meetings, nil
}

// ListRacesByMeetingDate will return all races between the provided start and end dates
func (repo *RacingRepository) ListRacesByMeetingDate(start, end int64) ([]*proto.Race, error) {
	var results []*model.Race

	log := logWithField("function", "ListRacesByMeetingDate")

	err := repo.collection(raceCollection).Find(
		bson.M{
			"meeting_start": bson.M{
				"$gte": time.Unix(start, 0),
				"$lte": time.Unix(end, 0),
			},
		},
	).All(&results)

	if err != nil {
		log.Errorf("An error occurred finding meetings starting between %v and %v - %v", start, end, err)
		return nil, err
	}

	races := model.RaceModelToProtoCollection(results)

	return races, nil
}

// ListRacesByMeetingID returns all races for the provided meeting id
func (repo *RacingRepository) ListRacesByMeetingID(meetingID string) ([]*proto.Race, error) {
	var results []*model.Race
	log := logWithField("function", "ListRacesByMeetingID")

	err := repo.collection(raceCollection).Find(bson.M{"meeting_id": meetingID}).All(&results)
	if err != nil {
		log.Errorf("An error occurred finding meeting %v - %v", meetingID, err)
		return nil, err
	}
	r := model.RaceModelToProtoCollection(results)
	return r, nil
}

// AddMeetings will add the provided meetings to the repository
func (repo *RacingRepository) AddMeetings(meetings []*proto.Meeting) error {

	m := model.MeetingProtoToModelCollection(meetings)
	log := logWithField("function", "AddMeetings")

	in := make([]interface{}, len(m))
	for i, v := range m {
		v.LastUpdated = time.Now()
		v.DateCreated = time.Now()
		in[i] = v
	}
	err := repo.collection(meetingCollection).Insert(in...)
	if err != nil {
		log.Errorf("An error occurred adding meetings - %v", err)
	}
	return err
}

// AddRaces will add the provided races to the repository
func (repo *RacingRepository) AddRaces(races []*proto.Race) error {
	r := model.RaceProtoToModelCollection(races)
	log := logWithField("function", "AddRaces")

	in := make([]interface{}, len(r))
	for i, v := range r {
		v.LastUpdated = time.Now()
		v.DateCreated = time.Now()
		in[i] = v
	}

	err := repo.collection(raceCollection).Insert(in...)
	if err != nil {
		log.Errorf("An error occurred adding races - %v", err)
	}
	return err
}

// AddSelections will add the provided selections to the repository
func (repo *RacingRepository) AddSelections(selections []*proto.Selection) error {
	s := model.SelectionProtoToModelCollection(selections)
	log := logWithField("function", "AddSelections")

	in := make([]interface{}, len(s))
	for i, v := range s {
		v.LastUpdated = time.Now()
		v.DateCreated = time.Now()
		in[i] = v
	}

	err := repo.collection(selectionCollection).Insert(in...)
	if err != nil {
		log.Errorf("An error occurred adding selections - %v", err)
	}

	return err
}

// GetRace retrieves a race using the provided race id
func (repo *RacingRepository) GetRace(raceID string) (*proto.Race, error) {
	r := &proto.Race{}
	log := logWithField("function", "GetRace")
	err := repo.collection(raceCollection).FindId(raceID).One(r)
	if err != nil {
		log.Errorf("An error occurred retrieving race with id %v - %v", raceID, err)
	}
	return r, err
}

// GetMeeting retrieves a meeting using the provided meeting id
func (repo *RacingRepository) GetMeeting(meetingID string) (*proto.Meeting, error) {
	m := &proto.Meeting{}
	log := logWithField("function", "GetMeeting")
	err := repo.collection(meetingCollection).FindId(meetingID).One(m)
	if err != nil {
		log.Errorf("An error occurred retrieving meeting with id %v - %v", meetingID, err)
	}
	return m, err
}

// ListSelectionsByRaceID retrieves all of the selections for the provided race id
func (repo *RacingRepository) ListSelectionsByRaceID(raceID string) ([]*proto.Selection, error) {
	var s []*model.Selection
	log := logWithField("function", "ListSelectionsByRaceID")
	err := repo.collection(selectionCollection).Find(bson.M{"race_id": raceID}).All(&s)
	if err != nil {
		log.Errorf("An error occurred retrieving selections with race id %v - %v", raceID, err)
	}

	p := model.SelectionModelToProtoCollection(s)

	return p, err
}

// UpdateRace updates the race record
func (repo *RacingRepository) UpdateRace(race *proto.RaceUpdatedMessage) error {
	updated := model.RaceUpdateProtoToModel(race)
	log := logWithField("function", "UpdateRace")

	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"scheduled_start": updated.ScheduledStart,
			"actual_start":    updated.ActualStart,
			"status":          updated.Status,
			"results":         updated.Results,
			"last_updated":    time.Now(),
		},
		},
	}

	original := &model.Race{}
	_, err := repo.collection(raceCollection).FindId(race.RaceId).Apply(change, original)
	if err != nil {
		log.Errorf("An error occurred updating race - %v", err)
	}

	return err
}

// UpdateSelection updates the selection record
func (repo *RacingRepository) UpdateSelection(s *proto.Selection) error {

	log := logWithField("function", "UpdateSelection")
	updated := model.SelectionProtoToModel(s)

	var change mgo.Change

	if updated.Scratched {
		change = mgo.Change{
			Update: bson.M{"$set": bson.M{
				"scratched":    updated.Scratched,
				"last_updated": time.Now(),
			},
			},
		}

	} else {
		change = mgo.Change{
			Update: bson.M{"$set": bson.M{
				"barrier_number":       updated.BarrierNumber,
				"jockey":               updated.Jockey,
				"number":               updated.Number,
				"source_competitor_id": updated.SourceCompetitorID,
				"name":                 updated.Name,
				"last_updated":         updated.LastUpdated,
			},
			},
		}
	}

	o := &model.Selection{}
	_, err := repo.collection(selectionCollection).FindId(s.SelectionId).Apply(change, o)
	if err != nil {
		log.Errorf("An error occurred updating selections - %v", err)
	}

	return err
}

// Close will close the connection to the repository
func (repo *RacingRepository) Close() {
	repo.session.Close()
}

func (repo *RacingRepository) collection(name string) *mgo.Collection {
	return repo.session.DB(dbName).C(name)
}
