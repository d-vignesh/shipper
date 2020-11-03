package main

import (
	pb "github.com/d-vignesh/shipper/vessel-service/proto/vessel"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"

)

type VesselHandler struct {
	session		*mgo.Session
}

func (vh *VesselHandler) GetRepo() Repository {
	return &VesselRepository{vh.session.Clone()}
}

// Create - stores the provided vessel into the database
func (vh *VesselHandler) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
	repo := vh.GetRepo()
	defer repo.Close()

	err := repo.Create(req)
	if err != nil {
		return err
	}

	resp.Created = true
	resp.Vessel = req
	return nil
}

func (vh *VesselHandler) FindAvailable(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
	repo := vh.GetRepo()
	defer repo.Close()

	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err 
	}

	resp.Vessel = vessel
	return nil
}