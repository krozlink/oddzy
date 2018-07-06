package model

import (
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	"time"
)

// Meeting represents an series of races held at a particular location on a day.
// This model is used for storing/retrieving data from mongodb
type Meeting struct {
	MeetingID      string    `bson:"_id,omitempty"`
	SourceID       string    `bson:"source_id"`
	Name           string    `bson:"name"`
	Country        string    `bson:"country"`
	RaceType       string    `bson:"race_type"`
	RaceIDs        []string  `bson:"race_ids"`
	ScheduledStart time.Time `bson:"scheduled_start"`
	DateCreated    time.Time `bson:"date_created"`
	LastUpdated    time.Time `bson:"last_updated"`
}

// MeetingProtoToModel converts a Meeting protobuf object used in service communication to a Meeting model object used in storage
func MeetingProtoToModel(p *proto.Meeting) *Meeting {
	m := &Meeting{
		MeetingID:      p.MeetingId,
		Country:        p.Country,
		SourceID:       p.SourceId,
		Name:           p.Name,
		RaceType:       p.RaceType,
		RaceIDs:        make([]string, len(p.RaceIds)),
		ScheduledStart: time.Unix(p.ScheduledStart, 0),
		LastUpdated:    time.Unix(p.LastUpdated, 0),
		DateCreated:    time.Unix(p.DateCreated, 0),
	}

	for i, v := range p.RaceIds {
		m.RaceIDs[i] = v
	}

	return m
}

// MeetingModelToProto converts a Meeting model object used in storage to a Meeting protobuf object
// used in service communication
func MeetingModelToProto(m *Meeting) *proto.Meeting {
	p := &proto.Meeting{
		MeetingId:      m.MeetingID,
		Country:        m.Country,
		SourceId:       m.SourceID,
		Name:           m.Name,
		RaceType:       m.RaceType,
		RaceIds:        make([]string, len(m.RaceIDs)),
		ScheduledStart: m.ScheduledStart.Unix(),
		LastUpdated:    m.LastUpdated.Unix(),
		DateCreated:    m.DateCreated.Unix(),
	}

	for i, v := range m.RaceIDs {
		p.RaceIds[i] = v
	}
	return p
}

// MeetingProtoToModelCollection converts a slice of Meeting protobuf objects used in service communication
// to a slice of Meeting model object used in storage
func MeetingProtoToModelCollection(p []*proto.Meeting) []*Meeting {
	// collection := make([]*Meeting, len(p), len(p))
	// for i, v := range p {
	// 	collection[i] = MeetingProtoToModel(v)
	// }
	// return collection
	return genericMap(p, MeetingProtoToModel).([]*Meeting)
}

// MeetingModelToProtoCollection converts a slice of Meeting model object used in storage
// to a slice of Meeting protobuf objects used in service communication
func MeetingModelToProtoCollection(m []*Meeting) []*proto.Meeting {
	// collection := make([]*proto.Meeting, len(m), len(m))
	// for i, v := range p {
	// 	collection[i] = MeetingModelToProto(v)
	// }
	// return collection
	return genericMap(m, MeetingModelToProto).([]*proto.Meeting)
}
