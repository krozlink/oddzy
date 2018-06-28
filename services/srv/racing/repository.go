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

const (
	repoSessionOpenedTiming = "racing.service.repository.sessionopened.timing"
	repoSessionOpenedCount  = "racing.service.repository.sessionopened.count"
	repoSessionClosedTiming = "racing.service.repository.sessionclosed.timing"
	repoSessionClosedCount  = "racing.service.repository.sessionclosed.count"

	repoListMeetingsByDateTiming      = "racing.service.repository.listmeetingsbydate.timing"
	repoListMeetingsByDateSuccess     = "racing.service.repository.listmeetingsbydate.success"
	repoListMeetingsByDateFailed      = "racing.service.repository.listmeetingsbydate.failed"
	repoListRacesByMeetingDateTiming  = "racing.service.repository.listracesbymeetingdate.timing"
	repoListRacesByMeetingDateSuccess = "racing.service.repository.listracesbymeetingdate.success"
	repoListRacesByMeetingDateFailed  = "racing.service.repository.listracesbymeetingdate.failed"
	repoListRacesByMeetingIDTiming    = "racing.service.repository.listracesbymeetingid.timing"
	repoListRacesByMeetingIDSuccess   = "racing.service.repository.listracesbymeetingid.success"
	repoListRacesByMeetingIDFailed    = "racing.service.repository.listracesbymeetingid.failed"
	repoAddMeetingsTiming             = "racing.service.repository.addmeetings.timing"
	repoAddMeetingsSuccess            = "racing.service.repository.addmeetings.success"
	repoAddMeetingsFailed             = "racing.service.repository.addmeetings.failed"
	repoAddRacesTiming                = "racing.service.repository.addraces.timing"
	repoAddRacesSuccess               = "racing.service.repository.addraces.success"
	repoAddRacesFailed                = "racing.service.repository.addraces.failed"
	repoAddSelectionsTiming           = "racing.service.repository.addselections.timing"
	repoAddSelectionsSuccess          = "racing.service.repository.addselections.success"
	repoAddSelectionsFailed           = "racing.service.repository.addselections.failed"
	repoGetRaceTiming                 = "racing.service.repository.getrace.timing"
	repoGetRaceSuccess                = "racing.service.repository.getrace.success"
	repoGetRaceFailed                 = "racing.service.repository.getrace.failed"
	repoGetMeetingTiming              = "racing.service.repository.getmeeting.timing"
	repoGetMeetingSuccess             = "racing.service.repository.getmeeting.success"
	repoGetMeetingFailed              = "racing.service.repository.getmeeting.failed"
	repoListSelectionsByRaceIDTiming  = "racing.service.repository.listselectionsbyraceid.timing"
	repoListSelectionsByRaceIDSuccess = "racing.service.repository.listselectionsbyraceid.success"
	repoListSelectionsByRaceIDFailed  = "racing.service.repository.listselectionsbyraceid.failed"
	repoUpdateRaceTiming              = "racing.service.repository.updaterace.timing"
	repoUpdateRaceSuccess             = "racing.service.repository.updaterace.success"
	repoUpdateRaceFailed              = "racing.service.repository.updaterace.failed"
	repoUpdateSelectionTiming         = "racing.service.repository.updateselection.timing"
	repoUpdateSelectionSuccess        = "racing.service.repository.updateselection.success"
	repoUpdateSelectionFailed         = "racing.service.repository.updateselection.failed"
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
	t := stats.NewTiming()
	defer t.Send(repoSessionOpenedTiming)
	stats.Increment(repoSessionOpenedCount)
	return &RacingRepository{repo.session.Clone()}
}

// ListMeetingsByDate will return all meetings between the provided start and end dates
func (repo *RacingRepository) ListMeetingsByDate(start, end int64) ([]*proto.Meeting, error) {
	t := stats.NewTiming()
	defer t.Send(repoListMeetingsByDateTiming)
	var results []*model.Meeting
	log := logWithField("function", "repository.ListMeetingsByDate")

	log.Debugf("Finding meetings starting between %v and %v", start, end)
	log.Debugf("Unix date range is %v and %v", time.Unix(start, 0), time.Unix(end, 0))

	err := repo.collection(meetingCollection).Find(
		bson.M{
			"scheduled_start": bson.M{
				"$gte": time.Unix(start, 0),
				"$lte": time.Unix(end, 0),
			},
		},
	).All(&results)

	if err != nil {
		log.Errorf("An error occurred finding meetings starting between %v and %v - %v", start, end, err)
		stats.Increment(repoListMeetingsByDateFailed)
		return nil, err
	}

	log.Debugf("Meetings found - %v", len(results))

	meetings := model.MeetingModelToProtoCollection(results)
	stats.Increment(repoListMeetingsByDateSuccess)
	return meetings, nil
}

// ListRacesByMeetingDate will return all races between the provided start and end dates
func (repo *RacingRepository) ListRacesByMeetingDate(start, end int64) ([]*proto.Race, error) {
	t := stats.NewTiming()
	defer t.Send(repoListRacesByMeetingDateTiming)
	var results []*model.Race

	log := logWithField("function", "repository.ListRacesByMeetingDate")

	err := repo.collection(raceCollection).Find(
		bson.M{
			"meeting_start": bson.M{
				"$gte": time.Unix(start, 0),
				"$lte": time.Unix(end, 0),
			},
		},
	).All(&results)

	if err != nil {
		stats.Increment(repoListRacesByMeetingDateFailed)
		log.Errorf("An error occurred finding races with meetings starting between %v and %v - %v", start, end, err)
		return nil, err
	}

	log.Debugf("Races found - %v", len(results))

	races := model.RaceModelToProtoCollection(results)
	stats.Increment(repoListRacesByMeetingDateSuccess)

	return races, nil
}

// ListRacesByMeetingID returns all races for the provided meeting id
func (repo *RacingRepository) ListRacesByMeetingID(meetingID string) ([]*proto.Race, error) {
	t := stats.NewTiming()
	defer t.Send(repoListRacesByMeetingIDTiming)
	var results []*model.Race
	log := logWithField("function", "repository.ListRacesByMeetingID")

	err := repo.collection(raceCollection).Find(bson.M{"meeting_id": meetingID}).All(&results)
	if err != nil {
		stats.Increment(repoListRacesByMeetingIDFailed)
		log.Errorf("An error occurred finding meeting %v - %v", meetingID, err)
		return nil, err
	}
	r := model.RaceModelToProtoCollection(results)
	stats.Increment(repoListRacesByMeetingIDSuccess)
	return r, nil
}

// AddMeetings will add the provided meetings to the repository
func (repo *RacingRepository) AddMeetings(meetings []*proto.Meeting) error {
	t := stats.NewTiming()
	defer t.Send(repoAddMeetingsTiming)
	m := model.MeetingProtoToModelCollection(meetings)
	log := logWithField("function", "repository.AddMeetings")

	in := make([]interface{}, len(m))
	for i, v := range m {
		v.LastUpdated = time.Now()
		v.DateCreated = time.Now()
		in[i] = v
	}
	err := repo.collection(meetingCollection).Insert(in...)
	if err != nil {
		stats.Increment(repoAddMeetingsFailed)
		log.Errorf("An error occurred adding meetings - %v", err)
		return err
	}

	stats.Increment(repoAddMeetingsSuccess)
	return nil
}

// AddRaces will add the provided races to the repository
func (repo *RacingRepository) AddRaces(races []*proto.Race) error {
	t := stats.NewTiming()
	defer t.Send(repoAddRacesTiming)
	r := model.RaceProtoToModelCollection(races)
	log := logWithField("function", "repository.AddRaces")

	in := make([]interface{}, len(r))
	for i, v := range r {
		v.LastUpdated = time.Now()
		v.DateCreated = time.Now()
		v.IsScraped = false
		in[i] = v
	}

	err := repo.collection(raceCollection).Insert(in...)
	if err != nil {
		stats.Increment(repoAddRacesFailed)
		log.Errorf("An error occurred adding races - %v", err)
		return err
	}

	stats.Increment(repoAddRacesSuccess)
	return nil
}

// AddSelections will add the provided selections to the repository
func (repo *RacingRepository) AddSelections(selections []*proto.Selection) error {
	t := stats.NewTiming()
	defer t.Send(repoAddSelectionsTiming)
	s := model.SelectionProtoToModelCollection(selections)
	log := logWithField("function", "repository.AddSelections")

	in := make([]interface{}, len(s))
	for i, v := range s {
		v.LastUpdated = time.Now()
		v.DateCreated = time.Now()
		in[i] = v
	}

	err := repo.collection(selectionCollection).Insert(in...)
	if err != nil {
		stats.Increment(repoAddSelectionsFailed)
		log.Errorf("An error occurred adding selections - %v", err)
		return err
	}

	stats.Increment(repoAddSelectionsSuccess)
	return nil
}

// GetRace retrieves a race using the provided race id
func (repo *RacingRepository) GetRace(raceID string) (*proto.Race, error) {
	t := stats.NewTiming()
	defer t.Send(repoGetRaceTiming)
	r := &proto.Race{}
	log := logWithField("function", "repository.GetRace")
	err := repo.collection(raceCollection).FindId(raceID).One(r)
	if err != nil {
		stats.Increment(repoGetRaceFailed)
		log.Errorf("An error occurred retrieving race with id %v - %v", raceID, err)
		return nil, err
	}

	stats.Increment(repoGetRaceSuccess)
	return r, nil
}

// GetMeeting retrieves a meeting using the provided meeting id
func (repo *RacingRepository) GetMeeting(meetingID string) (*proto.Meeting, error) {
	t := stats.NewTiming()
	defer t.Send(repoGetMeetingTiming)
	m := &proto.Meeting{}
	log := logWithField("function", "repository.GetMeeting")
	err := repo.collection(meetingCollection).FindId(meetingID).One(m)
	if err != nil {
		stats.Increment(repoGetMeetingFailed)
		log.Errorf("An error occurred retrieving meeting with id %v - %v", meetingID, err)
		return nil, err
	}

	stats.Increment(repoGetMeetingSuccess)
	return m, nil
}

// ListSelectionsByRaceID retrieves all of the selections for the provided race id
func (repo *RacingRepository) ListSelectionsByRaceID(raceID string) ([]*proto.Selection, error) {
	t := stats.NewTiming()
	defer t.Send(repoListSelectionsByRaceIDTiming)
	var s []*model.Selection
	log := logWithField("function", "repository.ListSelectionsByRaceID")
	err := repo.collection(selectionCollection).Find(bson.M{"race_id": raceID}).All(&s)
	if err != nil {
		stats.Increment(repoListSelectionsByRaceIDFailed)
		log.Errorf("An error occurred retrieving selections with race id %v - %v", raceID, err)
		return nil, err
	}

	p := model.SelectionModelToProtoCollection(s)
	stats.Increment(repoListSelectionsByRaceIDSuccess)
	return p, nil
}

// UpdateRace updates the race record
func (repo *RacingRepository) UpdateRace(race *proto.RaceUpdatedMessage) error {
	t := stats.NewTiming()
	defer t.Send(repoUpdateRaceTiming)
	updated := model.RaceUpdateProtoToModel(race)
	log := logWithField("function", "repository.UpdateRace")

	change := mgo.Change{
		Update: bson.M{"$set": bson.M{
			"scheduled_start": updated.ScheduledStart,
			"actual_start":    updated.ActualStart,
			"status":          updated.Status,
			"results":         updated.Results,
			"is_scraped":      true,
			"last_updated":    time.Now(),
		},
		},
	}

	original := &model.Race{}
	_, err := repo.collection(raceCollection).FindId(race.RaceId).Apply(change, original)
	if err != nil {
		stats.Increment(repoUpdateRaceFailed)
		log.Errorf("An error occurred updating race - %v", err)
		return err
	}

	stats.Increment(repoUpdateRaceSuccess)
	return nil
}

// UpdateSelection updates the selection record
func (repo *RacingRepository) UpdateSelection(s *proto.Selection) error {
	t := stats.NewTiming()
	defer t.Send(repoUpdateSelectionTiming)
	log := logWithField("function", "repository.UpdateSelection")
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
		stats.Increment(repoUpdateSelectionFailed)
		log.Errorf("An error occurred updating selections - %v", err)
		return err
	}

	stats.Increment(repoUpdateSelectionSuccess)
	return nil
}

// Close will close the connection to the repository
func (repo *RacingRepository) Close() {
	t := stats.NewTiming()
	defer t.Send(repoSessionClosedTiming)
	stats.Increment(repoSessionClosedCount)
	repo.session.Close()
}

func (repo *RacingRepository) collection(name string) *mgo.Collection {
	return repo.session.DB(dbName).C(name)
}
