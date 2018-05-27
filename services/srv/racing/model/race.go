package model

import (
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	"time"
)

// Race represents an instance of a horse/greyhounds/harness race
// This model is used for storing/retrieving data from mongodb
type Race struct {
	RaceID         string    `bson:"_id,omitempty"`
	SourceID       string    `bson:"source_id"`
	Number         int32     `bson:"number"`
	Name           string    `bson:"name"`
	ScheduledStart time.Time `bson:"scheduled_start"`
	ActualStart    time.Time `bson:"actual_start"`
	Status         string    `bson:"status"`
	Results        string    `bson:"results"`
	MeetingStart   time.Time `bson:"meeting_start"`
	DateCreated    time.Time `bson:"date_created"`
	LastUpdated    time.Time `bson:"last_updated"`
}

// RaceProtoToModel converts a Race protobuf object used in service communication to a Race model object used in storage
func RaceProtoToModel(p *proto.Race) *Race {
	return &Race{
		RaceID:         p.RaceId,
		SourceID:       p.SourceId,
		Number:         p.Number,
		Name:           p.Name,
		ScheduledStart: time.Unix(p.ScheduledStart, 0),
		ActualStart:    time.Unix(p.ActualStart, 0),
		Status:         p.Status,
		Results:        p.Results,
		MeetingStart:   time.Unix(p.MeetingStart, 0),
		DateCreated:    time.Unix(p.DateCreated, 0),
		LastUpdated:    time.Unix(p.LastUpdated, 0),
	}
}

// RaceModelToProto converts a Race model object used in storage to a Race protobuf object
// used in service communication
func RaceModelToProto(r Race) *proto.Race {
	return &proto.Race{
		RaceId:         r.RaceID,
		SourceId:       r.SourceID,
		Number:         r.Number,
		Name:           r.Name,
		ScheduledStart: r.ScheduledStart.Unix(),
		ActualStart:    r.ActualStart.Unix(),
		Status:         r.Status,
		Results:        r.Results,
		MeetingStart:   r.MeetingStart.Unix(),
		DateCreated:    r.DateCreated.Unix(),
		LastUpdated:    r.LastUpdated.Unix(),
	}
}

// RaceProtoToModelCollection converts a slice of Race protobuf objects used in service communication
// to a slice of Race model object used in storage
func RaceProtoToModelCollection(p []*proto.Race) []*Race {
	return genericMap(p, RaceProtoToModel).([]*Race)
}

// RaceModelToProtoCollection converts a slice of Race model object used in storage
// to a slice of Race protobuf objects used in service communication
func RaceModelToProtoCollection(p []*Race) []*proto.Race {
	return genericMap(p, RaceModelToProto).([]*proto.Race)
}
