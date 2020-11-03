package main

import (
	"fmt"
	"os"
	"log"

	pb "github.com/d-vignesh/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro/v2"
)

const (
	defaultHost = "localhost:27017"
)

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "v002", Name: "light duty", MaxWeight: 200, Capacity: 3},
		{Id: "v001", Name: "heavy duty", MaxWeight: 20000, Capacity: 500},
	}

	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Fatalf("error connect to datastore on host %s - %v", host, err)
	}

	repo := &VesselRepository{session.Copy()}
	createDummyData(repo)

	vh := &VesselHandler{session}

	ms := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	ms.Init()

	pb.RegisterVesselServiceHandler(ms.Server(), vh)

	if err := ms.Run(); err != nil {
		fmt.Println(err)
	}
}
