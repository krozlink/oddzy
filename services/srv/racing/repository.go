package main

import (
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
	// repo.collection(meetingCollection).Find(
	// 	bson.M{
	// 		"scheduled_start": bson.M{
	// 			"$gte": start,
	// 			"$lte": end,
	// 		},
	// 	},
	// ).All()
	return nil, nil
}

func (repo *RacingRepository) ListRacesByMeetingDate(start, end int64) ([]*proto.Race, error) {

	return nil, nil
}

func (repo *RacingRepository) AddMeetings(meetings []*proto.Meeting) error {
	return repo.collection(meetingCollection).Insert(meetings)
}

func (repo *RacingRepository) AddRaces(races []*proto.Race) error {
	return repo.collection(raceCollection).Insert(races)
}

func (repo *RacingRepository) UpdateRace(race *proto.Race, selections []*proto.Selection) error {
	raceErr := repo.collection(raceCollection).Update(bson.M{"_id": race.RaceId}, race)
	// repo.collection(selectionCollection).UpdateAll()
	return raceErr
}

func (repo *RacingRepository) Close() {
	repo.session.Close()
}

func (repo *RacingRepository) collection(name string) *mgo.Collection {
	return repo.session.DB(dbName).C(name)
}
