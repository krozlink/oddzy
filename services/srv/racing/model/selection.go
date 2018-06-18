package model

import (
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
	"time"
)

// Selection represents an instance of an animal running in a particular race
// The data is specific to the race itself. Data specific to the animal is stored in the Competitor entity
// This model is used for storing/retrieving data from mongodb
type Selection struct {
	SelectionID        string    `bson:"_id,omitempty"`
	SourceID           string    `bson:"source_id"`
	CompetitorID       string    `bson:"competitor_id"`
	SourceCompetitorID string    `bson:"source_competitor_id"`
	RaceID             string    `bson:"race_id"`
	Name               string    `bson:"name"`
	Jockey             string    `bson:"jockey"`
	Number             int32     `bson:"number"`
	BarrierNumber      int32     `bson:"barrier_number"`
	LastUpdated        time.Time `bson:"last_updated"`
	Scratched          bool      `bson:"scratched"`
}

// SelectionProtoToModel converts a Selection protobuf object used in service communication
// to a Selection model object used in storage
func SelectionProtoToModel(p *proto.Selection) *Selection {
	return &Selection{
		SelectionID:        p.SelectionId,
		SourceID:           p.SourceId,
		CompetitorID:       p.CompetitorId,
		SourceCompetitorID: p.SourceCompetitorId,
		RaceID:             p.RaceId,
		Name:               p.Name,
		Jockey:             p.Jockey,
		Number:             p.Number,
		BarrierNumber:      p.BarrierNumber,
		Scratched:          p.Scratched,
		LastUpdated:        time.Unix(p.LastUpdated, 0),
	}
}

// SelectionModelToProto converts a Selection model object used in storage to a Selection protobuf object
// used in service communication
func SelectionModelToProto(s *Selection) *proto.Selection {
	return &proto.Selection{
		SelectionId:        s.SelectionID,
		SourceId:           s.SourceID,
		CompetitorId:       s.CompetitorID,
		SourceCompetitorId: s.SourceCompetitorID,
		RaceId:             s.RaceID,
		Name:               s.Name,
		Jockey:             s.Jockey,
		Number:             s.Number,
		BarrierNumber:      s.BarrierNumber,
		Scratched:          s.Scratched,
		LastUpdated:        s.LastUpdated.Unix(),
	}
}

// SelectionProtoToModelCollection converts a slice of Selection protobuf objects used in service communication
// to a slice of Selection model object used in storage
func SelectionProtoToModelCollection(p []*proto.Selection) []*Selection {
	return genericMap(p, SelectionProtoToModel).([]*Selection)
}

// SelectionModelToProtoCollection converts a Selection model object used in storage to a Selection protobuf object
// used in service communication
func SelectionModelToProtoCollection(p []*Selection) []*proto.Selection {
	return genericMap(p, SelectionModelToProto).([]*proto.Selection)
}
