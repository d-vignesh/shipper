package main

import (
	"log"
	"golang.org/x/net/context"
	pb "github.com/d-vignesh/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/d-vignesh/shipper/vessel-service/proto/vessel"
)

type ConsignmentHandler struct {
	session			*mgo.Session
	vesselClient	vesselProto.VesselService
}

// GetRepo - provides a new Repository with the handlers mongo session
func (ch *ConsignmentHandler) GetRepo() Repository {
	return &ConsignmentRepository{ch.session.Clone()}
}

// CreateConsignment - store the request consignment into the database
func (ch *ConsignmentHandler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := ch.GetRepo()
	defer repo.Close()

	// make a request to the vessel-service to check for a vessel 
	// that is satisfying the given consignment specification
	vesselResp, err := ch.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}
	log.Printf("found vessel: %s \n", vesselResp.Vessel.Name)

	req.VesselId = vesselResp.Vessel.Id 

	// save the consignment
	err = repo.Create(req)
	if err != nil {
		return err
	}

	// update the response
	res.Created = true
	res.Consignment = req
	return nil
}

func (ch *ConsignmentHandler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := ch.GetRepo()
	defer repo.Close()

	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}