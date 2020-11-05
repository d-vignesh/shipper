package main

import (
	"fmt"
	"log"
	"context"
	"os"

	pb "github.com/d-vignesh/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/d-vignesh/shipper/vessel-service/proto/vessel"
	"github.com/micro/go-micro/v2"
)

func main() {
	
	service := micro.NewService(
		micro.Name("shipper.consignment.service")
	)
	service.Init()

	uri := os.Getenv("DB_HOST")

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselService("shipper.vessel.client", service.Client())
	ch := &ConsignmentHandler{repository, vesselClient}

	pb.RegisterShippingServiceHandler(service.Server(), ch)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
