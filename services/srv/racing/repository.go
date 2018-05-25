package main

import (
	model "github.com/krozlink/oddzy/services/srv/racing/model"
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName               = "oddzy"
	raceCollection       = "races"
	meetingCollection    = "meetings"
	selectionCollection  = "selections"
	competitorCollection = "competitors"
)

type Repository interface {
	ListMeetingsByDate(start, end int64) ([]*proto.Meeting, error)
	ListRacesByMeetingDate(start, end int64) ([]*proto.Race, error)
	AddMeetings(meetings []*proto.Meeting) error
	AddRaces(races []*proto.Race) error
	UpdateRace(race *proto.Race, selections []*proto.Selection) error
	Close()
}

type RacingRepository struct {
	session *mgo.Session
}

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

func (repo *RacingRepository) AddMeetings(meetings []*proto.Meeting) error {
	m := model.MeetingProtoToModelCollection(meetings)
	return repo.collection(meetingCollection).Insert(m)
}

func (repo *RacingRepository) AddRaces(races []*proto.Race) error {
	r := model.RaceProtoToModelCollection(races)
	return repo.collection(raceCollection).Insert(r)
}

func (repo *RacingRepository) UpdateRace(race *proto.Race, selections []*proto.Selection) error {

	// read race
	// update new data
	// lookup selection ids ?????
	// upsert selections

	r := model.RaceProtoToModel(race)
	raceErr := repo.collection(raceCollection).Update(bson.M{"_id": race.RaceId}, r)
	// repo.collection(selectionCollection).UpdateAll()
	return raceErr
}

func (repo *RacingRepository) Close() {
	repo.session.Close()
}

func (repo *RacingRepository) collection(name string) *mgo.Collection {
	return repo.session.DB(dbName).C(name)
}
