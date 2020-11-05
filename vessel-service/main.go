package main

import (
	"fmt"
	"os"
	"log"

	pb "github.com/d-vignesh/shipper/vessel-service/proto/vessel"
	"github.com/micro/go-micro/v2"
)

func main() {

	service := micro.NewService(
		micro.Name("shipper.vessel.service")
	)
	service.Init()

	uri := os.Getenv("DB_HOST")

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shippy").Collection("vessels")
	repository := &MongoRepository{vesselCollection}

	vh := &VesselHandler{repository}

	if err := pb.RegisterVesselServiceHandler(service.Server(), vh); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
