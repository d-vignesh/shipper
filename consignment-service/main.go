package main

import (
	"fmt"
	"log"

	pb "github.com/d-vignesh/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/d-vignesh/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro/v2"
	"os"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	// Get database host from env variable
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	// Mgo creates a master session which needs to be closed at function end
	defer session.Close()

	if err != nil {
		log.Panicf("could not connect to datastore on host %s - %v", host, err)
	}

	// create a new micro service
	ms := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// create a new vessel service client
	vc := vesselProto.NewVesselService("go.micro.srv.vessel", ms.Client())

	ms.Init()

	ch := &ConsignmentHandler{session: session, vesselClient: vc}

	pb.RegisterShippingServiceHandler(ms.Server(), ch)

	if err := ms.Run(); err != nil {
		fmt.Println(err)
	}
}
