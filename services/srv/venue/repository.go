package main

import (
	"errors"
	proto "github.com/krozlink/oddzy/services/srv/venue/proto"
	"gopkg.in/mgo.v2"
)

const (
	dbName          = "oddzy"
	venueCollection = "venues"
)

type Repository interface {
	Get(id string) (*proto.Venue, error)
	List() ([]*proto.Venue, error)
	Delete(id string) error
	Add(*proto.Venue) error
	Close()
}

type VenueRepository struct {
	session *mgo.Session
}

func (repo *VenueRepository) Get(id string) (*proto.Venue, error) {
	venue := &proto.Venue{
		Id: id,
	}

	if err := repo.collection().FindId(id).One(venue).Error(); err != "" {
		return nil, errors.New(err)
	}

	return venue, nil
}

func (repo *VenueRepository) Close() {
	repo.session.Close()
}

func (repo *VenueRepository) Delete(id string) error {
	return repo.collection().RemoveId(id)
}

func (repo *VenueRepository) Add(venue *proto.Venue) error {
	return repo.collection().Insert(venue)
}

func (repo *VenueRepository) List() ([]*proto.Venue, error) {
	var results []*proto.Venue
	err := repo.collection().Find(nil).All(&results)
	return results, err
}

func (repo *VenueRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(venueCollection)
}
