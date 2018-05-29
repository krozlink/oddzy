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
	ListMeetingsByDate(start, end int64) ([]*proto.Meeting, error)
	ListRacesByMeetingDate(start, end int64) ([]*proto.Race, error)
	AddMeetings(meetings []*proto.Meeting) error
	AddRaces(races []*proto.Race) error

	GetRace(raceID string) (*proto.Race, error)
	GetSelectionsByRaceID(raceID string) ([]*proto.Selection, error)
	UpdateRace(race *proto.Race) error
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
				"$gte": start,
				"$lte": end,
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

	err := repo.collection(raceCollection).Find(
		bson.M{
			"meeting_start": bson.M{
				"$gte": start,
				"$lte": end,
			},
		},
	).All(&results)

	if err != nil {
		return nil, err
	}

	races := model.RaceModelToProtoCollection(results)

	return races, nil
}

// AddMeetings will add the provided meetings to the repository
func (repo *RacingRepository) AddMeetings(meetings []*proto.Meeting) error {
	m := model.MeetingProtoToModelCollection(meetings)
	return repo.collection(meetingCollection).Insert(m)
}

// AddRaces will add the provided races to the repository
func (repo *RacingRepository) AddRaces(races []*proto.Race) error {
	r := model.RaceProtoToModelCollection(races)
	return repo.collection(raceCollection).Insert(r)
}

// GetRace retrieves a race using the provided race id
func (repo *RacingRepository) GetRace(raceID string) (*proto.Race, error) {
	r := &proto.Race{}
	err := repo.collection(raceCollection).FindId(raceID).One(r)
	return r, err
}

// GetSelectionsByRaceID retrieves all of the selections for the provided race id
func (repo *RacingRepository) GetSelectionsByRaceID(raceID string) ([]*proto.Selection, error) {
	var s []*model.Selection
	err := repo.collection(selectionCollection).Find(bson.M{"race_id": raceID}).All(&s)

	p := model.SelectionModelToProtoCollection(s)

	return p, err
}

// UpdateRace updates the race record
func (repo *RacingRepository) UpdateRace(race *proto.Race) error {
	updated := model.RaceProtoToModel(race)

	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"scheduled_start": updated.ScheduledStart,
			"actual_start":    updated.ActualStart,
			"status":          updated.Status,
			"results":         updated.Results,
			"last_updated":    time.Now().Unix(),
		},
		},
	}

	original := &model.Race{}
	_, err := repo.collection(raceCollection).FindId(race.RaceId).Apply(change, original)

	return err
}

// UpdateSelection updates the selection record
func (repo *RacingRepository) UpdateSelection(s *proto.Selection) error {

	updated := model.SelectionProtoToModel(s)

	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"barrier_number":       updated.BarrierNumber,
			"jockey":               updated.Jockey,
			"number":               updated.Number,
			"source_competitor_id": updated.SourceCompetitorID,
			"name":                 updated.Name,
			"last_updated":         time.Now().Unix(),
		},
		},
	}

	o := &model.Selection{}
	_, err := repo.collection(selectionCollection).FindId(s.SelectionId).Apply(change, o)

	return err
}

// Close will close the connection to the repository
func (repo *RacingRepository) Close() {
	repo.session.Close()
}

func (repo *RacingRepository) collection(name string) *mgo.Collection {
	return repo.session.DB(dbName).C(name)
}
