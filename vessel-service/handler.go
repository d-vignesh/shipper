package main

import (
	"context"

	pb "github.com/d-vignesh/shipper/vessel-service/proto/vessel"

)

type VesselHandler struct {
	repository	*MongoRepository
}

// Create - stores the provided vessel into the database
func (vh *VesselHandler) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
	if err := vh.repository.Create(ctx, MarshalVessel(req)); err != nil {
		return err
	}
	resp.Vessel = req
	return nil
}

func (vh *VesselHandler) FindAvailable(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
	
	vessel, err := vh.repository.FindAvailable(ctx, MarshalSpecification(req))
	if err != nil {
		return err
	}

	resp.Vessel = UnmarshalVessel(vessel)
	return nil
}