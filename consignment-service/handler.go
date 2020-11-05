package main

import (
	"context"

	pb "github.com/d-vignesh/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/d-vignesh/shipper/vessel-service/proto/vessel"
	"github.com/pkg/errors"
)

type ConsignmentHandler struct {
	repository 		Repository
	vesselClient	vesselProto.VesselService
}

// CreateConsignment - store the request consignment into the database
func (ch *ConsignmentHandler) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {

	vesselResp, err := ch.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity : int32(len(req.Containers)),
	})
	if vesselResp == nil {
		return errors.New("no vessel found for the specification, got nil response")
	}

	if err != nil {
		return nil
	}

	req.VesselId = vesselResponse.Vessel.Id

	if err = ch.repository.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}

	resp.Created = true
	resp.Consignment = req
	return nil
}

func (ch *ConsignmentHandler) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {

	consignments, err := ch.repository.GetAll(ctx)
	if err != nil {
		return err
	}
	resp.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}