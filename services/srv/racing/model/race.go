package model

import (
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	"time"
)

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
}

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
	}
}

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
	}
}

func RaceProtoToModelCollection(p []*proto.Race) []*Race {
	return genericMap(p, RaceProtoToModel).([]*Race)
}

func RaceModelToProtoCollection(p []*Race) []*proto.Race {
	return genericMap(p, RaceModelToProto).([]*proto.Race)
}
