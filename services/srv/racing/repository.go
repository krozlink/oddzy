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

// Repository interface used to access racing data
type Repository interface {
	ListMeetingsByDate(start, end int64) ([]*proto.Meeting, error)
	ListRacesByMeetingDate(start, end int64) ([]*proto.Race, error)
	AddMeetings(meetings []*proto.Meeting) error
	AddRaces(races []*proto.Race) error
	UpdateRace(race *proto.Race, selections []*proto.Selection) error
	Close()
	NewSession() Repository
}

// RacingRepository implements the Repository interface to access racing data
type RacingRepository struct {
	session *mgo.Session
}

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

// UpdateRace will update the race and selection data for the provided race
func (repo *RacingRepository) UpdateRace(race *proto.Race, selections []*proto.Selection) error {

	raceUpdated, err := repo.updateRaceModel(race)
	if err != nil {
		return err
	}

	selectionsUpdated, err := repo.updateSelectionModels(race.RaceId, selections)
	if err != nil {
		return err
	}

	if raceUpdated || selectionsUpdated {
		onRaceUpdated(race, selections)
	}

	return nil
}

// Close will close the connection to the repository
func (repo *RacingRepository) Close() {
	repo.session.Close()
}

func (repo *RacingRepository) collection(name string) *mgo.Collection {
	return repo.session.DB(dbName).C(name)
}

func (repo *RacingRepository) updateRaceModel(race *proto.Race) (bool, error) {
	updated := model.RaceProtoToModel(race)

	c := mgo.Change{
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
	_, err := repo.collection(raceCollection).FindId(race.RaceId).Apply(c, original)

	if err != nil {
		fmt.Printf("Failed to update race: %s", err)
		return false, err
	}

	raceUpdated := (original.ScheduledStart != updated.ScheduledStart ||
		original.ActualStart != updated.ActualStart ||
		original.Status != updated.Status ||
		original.Results != updated.Results)

	return raceUpdated, nil
}

func (repo *RacingRepository) updateSelectionModels(raceID string, selections []*proto.Selection) (bool, error) {

	updated := model.SelectionProtoToModelCollection(selections)

	// lookup selections
	var existing []model.Selection
	err := repo.collection(selectionCollection).Find(bson.M{"race_id": raceID}).All(&existing)

	if err != nil {
		fmt.Printf("Failed to read selections: %s", err)
		return false, err
	}

	raceUpdated := false
	if len(existing) == 0 {
		raceUpdated = true
		err = repo.collection(selectionCollection).Insert(updated)
	} else if len(existing) == len(updated) {
		for _, v := range updated {
			original := getSelectionByID(v.SelectionID, &existing)
			if original == nil {
				return false, fmt.Errorf("Unexpected update of race - selection %v (%v)does not exist", v.SelectionID, v.SourceID)
			}

			if getSelectionBySourceID(v.SourceID, &existing) == nil {
				return false, fmt.Errorf("Unexpected update of race - selection %v (%v)does not exist", v.SelectionID, v.SourceID)
			}

			if selectionChanged(original, v) {
				raceUpdated = true

				c := mgo.Change{
					Update: bson.M{"$set": bson.M{
						"barrier_number": v.BarrierNumber,
						"jockey":         v.Jockey,
						"name":           v.Name,
						"number":         v.Number,
						"last_updated":   time.Now().Unix(),
					},
					},
				}

				repo.collection(selectionCollection).FindId(v.SelectionID).Apply(c, original)
			}
		}
	} else {
		return false, fmt.Errorf("Unexpected race update - number of selections has changed from %v to %v", len(existing), len(updated))
	}

	return raceUpdated, err
}

func getSelectionByID(selectionID string, selections *[]model.Selection) *model.Selection {
	for _, v := range *selections {
		if v.SelectionID == selectionID {
			return &v
		}
	}
	return nil
}

func getSelectionBySourceID(sourceID string, selections *[]model.Selection) *model.Selection {
	for _, v := range *selections {
		if v.SourceID == sourceID {
			return &v
		}
	}
	return nil
}

func selectionChanged(from, to *model.Selection) bool {
	return from.BarrierNumber != to.BarrierNumber ||
		from.Jockey != to.Jockey ||
		from.Name != to.Name ||
		from.Number != to.Number
}

func onRaceUpdated(race *proto.Race, selections []*proto.Selection) {
	// TODO: publish message to Race Updated topic
}
