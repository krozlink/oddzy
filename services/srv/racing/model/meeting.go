package model

import (
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	"time"
)

type Meeting struct {
	MeetingID      string `bson:"_id,omitempty"`
	SourceID       string `bson:"source_id"`
	Name           string `bson:"name"`
	Country        string `bson:"country"`
	RaceType       string `bson:"race_type"`
	ScheduledStart time.Time `bson:"scheduled_start"`
}

func MeetingProtoToModel(p *proto.Meeting) *Meeting {
	return &Meeting{
		MeetingID:      p.MeetingId,
		Country:        p.Country,
		SourceID:       p.SourceId,
		Name:           p.Name,
		RaceType:       p.RaceType,
		ScheduledStart: time.Unix(p.ScheduledStart, 0),
	}
}


func MeetingModelToProto(m *Meeting) *proto.Meeting {
	return &proto.Meeting{
		MeetingId:      m.MeetingID,
		Country:        m.Country,
		SourceId:       m.SourceID,
		Name:           m.Name,
		RaceType:       m.RaceType,
		ScheduledStart: m.ScheduledStart.Unix(),
	}
}



func MeetingProtoToModelCollection(p []*proto.Meeting) []*Meeting {
	// collection := make([]*Meeting, len(p), len(p))
	// for i, v := range p {
	// 	collection[i] = MeetingProtoToModel(v)
	// }
	// return collection
	return genericMap(p, MeetingProtoToModel).([]*Meeting)
}


func MeetingModelToProtoCollection(m []*Meeting) []*proto.Meeting {
	// collection := make([]*proto.Meeting, len(m), len(m))
	// for i, v := range p {
	// 	collection[i] = MeetingModelToProto(v)
	// }
	// return collection
	return genericMap(m, MeetingModelToProto).([]*proto.Meeting)
}
