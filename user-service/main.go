package main

import (
	"fmt"
	"log"
	pb "github.com/d-vignesh/shipper/user-service/proto/user"
	"github.com/micro/go-micro/v2"
)

func main() {
	// create the db connection
	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("could not connect to DB : %v", err)
	}
	defer db.Close()

	// automatically migrate the user struct into database column.
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	ms := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)
	ms.Init()

	pb.RegisterUserServiceHandler(ms.Server(), &UserHandler{repo})

	if err := ms.Run(); err != nil {
		fmt.Println(err)
	}
}