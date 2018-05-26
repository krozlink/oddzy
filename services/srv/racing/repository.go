package main

import (
	"fmt"
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

	err := repo.updateRaceModel(race)
	if err != nil {
		return err
	}

	err = repo.updateSelectionModels(race.RaceId, selections)
	if err != nil {
		return err
	}

	return nil
}

func (repo *RacingRepository) Close() {
	repo.session.Close()
}

func (repo *RacingRepository) collection(name string) *mgo.Collection {
	return repo.session.DB(dbName).C(name)
}

func (repo *RacingRepository) updateRaceModel(race *proto.Race) error {
	r := model.RaceProtoToModel(race)

	if r.RaceID == "" {
		return fmt.Errorf("Expected race to have race id")
	}

	c := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"scheduled_start": r.ScheduledStart,
			"actual_start":    r.ActualStart,
			"status":          r.Status,
			"results":         r.Results,
			"meeting_start":   r.MeetingStart,
			"last_updated":    time.Now().Unix(),
		},
		},
	}

	original := &model.Race{}
	_, err := repo.collection(raceCollection).FindId(race.RaceId).Apply(c, original)

	if err != nil {
		fmt.Printf("Failed to update race: %s", err)
		return err
	}

	//compare original to existing
	//if changes, publish to race updated topic

	return nil
}

func (repo *RacingRepository) updateSelectionModels(raceID string, selections []*proto.Selection) error {

	s := model.SelectionProtoToModelCollection(selections)

	for _, v := range s {
		if v.SelectionID == "" {
			return fmt.Errorf("Expected selection to have selection id")
		}

		if v.RaceID == "" {
			return fmt.Errorf("Expected selection to have race id")			
		}
	}

	// lookup selections
	var existing []model.Selection
	err := repo.collection(selectionCollection).Find(bson.M{"race_id": raceID}).All(&existing)

	if err != nil {
		fmt.Printf("Failed to read selections: %s", err)
		return err
	}

	if len(existing) == 0 {
		err = repo.collection(selectionCollection).Insert(s)
	} else if len(existing) == len(s) {

		//loop through each item
		//find its matching existing item
		//update with existing selection id
		//compare new to existing
		//if changes, perform update (and also publish to selection updated topic)
	} else {
		return fmt.Errorf("Unexpected race update - number of selections has changed from %v to %v", len(existing), len(s))
	}

	return err
}
