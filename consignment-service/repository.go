package main

import (
	pb "github.com/d-vignesh/shipper/consignment-service/proto/consignment"
	"gopkg.in/mgo.v2"
)

const (
	dbName = "shippy"
	consignmentCollection = "consignments"
)

type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

type ConsignmentRepository struct {
	session *mgo.Session 
}

// Create a new consignment
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

// GetAll consignments
func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment
	err := repo.collection().Find(nil).All(&consignments)
	return consignments, err
}

// Close closes the database session After each query has ran.
func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}

func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(consignmentCollection)
}
