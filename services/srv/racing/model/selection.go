package model

import (
	proto "github.com/krozlink/oddzy/services/srv/racing/proto"
)

type Selection struct {
	SelectionID        string `bson:"_id,omitempty"`
	SourceID           string `bson:"source_id"`
	CompetitorID       string `bson:"competitor_id"`
	SourceCompetitorID string `bson:"source_competitor_id"`
	Name               string `bson:"name"`
	Jockey             string `bson:"jockey"`
	Number             int32  `bson:"number"`
	BarrierNumber      int32  `bson:"barrier_number"`
}

func SelectionProtoToModel(p *proto.Selection) *Selection {
	return &Selection{
		SelectionID:        p.SelectionId,
		SourceID:           p.SourceId,
		CompetitorID:       p.CompetitorId,
		SourceCompetitorID: p.SourceCompetitorId,
		Name:               p.Name,
		Jockey:             p.Jockey,
		Number:             p.Number,
		BarrierNumber:      p.BarrierNumber,
	}
}

func SelectionModelToProto(s Selection) *proto.Selection {
	return &proto.Selection{
		SelectionId:        s.SelectionID,
		SourceId:           s.SourceID,
		CompetitorId:       s.CompetitorID,
		SourceCompetitorId: s.SourceCompetitorID,
		Name:               s.Name,
		Jockey:             s.Jockey,
		Number:             s.Number,
		BarrierNumber:      s.BarrierNumber,
	}
}
