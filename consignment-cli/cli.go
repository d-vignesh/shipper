package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/d-vignesh/shipper/consignment-service/proto/consignment"
	"github.com/micro/go-micro/v2"
	"golang.org/x/net/context"
)

const (
	defaultFilename = "consignment.json"
)

func parserFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, nil
}

func main() {

	// create a new shipping service client
	service := micro.NewService()
	service.Init()

	client := pb.NewShippingService("go.micro.srv.consignment", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parserFile(file)

	if err != nil {
		log.Fatalf("could not parse file : %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("could not greet : %v", err)
	}
	log.Printf("consignment creations status : %t", r.Created)

	getResp, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("could not get all consignments: %v", err)
	}
	for _, c := range getResp.Consignments {
		log.Println(c)
	}
}
