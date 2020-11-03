package main

import (
	pb "github.com/d-vignesh/shipper/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName =  "shippy"
	vesselCollection = "vessels"
)

type Repository interface {
	Create(*pb.Vessel) error
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Close()
}

type VesselRepository struct {
	session *mgo.Session 
}

// Create - stores the provided vessel into the database
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

// FindAvailable - checks the given specification against all available vessels
// and returns the vessel satisfying the specification
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	err := repo.collection().Find(bson.M{
		"capacity": bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)

	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}