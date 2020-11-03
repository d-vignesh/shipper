package main

import (
	"log"
	"os"

	pb "github.com/d-vignesh/shipper/user-service/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/cli"
	"golang.org/x/net/context"
)

func main() {

	// create a new user service client
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name: "name",
				Usage: "your full name",
			},
			cli.StringFlag{
				Name: "email",
				Usage: "you email",
			},
			cli.StringFlag{
				Name: "password",
				Usage: "your password",
			},
			cli.StringFlag{
				Name: "company",
				Usage: "your company",
			},
		),
	)

	client := pb.NewUserService("go.micro.srv.user", service.Client())
	
	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			// call our user service
			resp, err := client.Create(context.TODO(), &pb.User{
				Name: name,
				Email: email,
				Password: password,
				Company: company,
			})
			if err != nil {
				log.Fatalf("could not create user : %v", err)
			}
			log.Printf("Created user : %s", resp.User.Id)

			getAllResp, err := client.GetAll(context.Background(), &pb.Request{})
			if err := nil {
				log.Fatalf("could not get users list : %v", err)
			}

			for _, user := range getAllResp.Users {
				log.Println(user)
			}
			os.Exit(0)
		}),
	)	
	
	// run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
	
}