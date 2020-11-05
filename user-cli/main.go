package main

import (
	"log"
	"context"
	"log"

	pb "github.com/d-vignesh/shipper/user-service/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/cli/v2"
)

func createUser(ctx context.Context, service micro.Service, user *pb.User) error {
	client := pb.NewUserService("shipper.service.user", service.Client())
	resp, err := client.Create(ctx, user)
	if err != nil {
		return err
	}

	fmt.Println("user created : ", resp.User)
	return nil
}

func main() {

	// create a new service 
	service := micro.NewService(
		micro.Flags(
			&cli.StringFlag{
				Name: "name",
				Usage: "your full name",
			},
			&cli.StringFlag{
				Name: "email",
				Usage: "you email",
			},
			&cli.StringFlag{
				Name: "password",
				Usage: "your password",
			},
			&cli.StringFlag{
				Name: "company",
				Usage: "your company",
			},
		),
	)
	
	service.Init(
		micro.Action(func(c *cli.Context) error {
			log.Println(c)
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			log.Println("test : ", name, email, company, password)

			ctx := context.Background()
			user := &pb.User{
				Name: 	  name,
				Email:    email,
				Company:  company,
				Password: password,
			}

			if err := createUser(ctx, service, user); err != nil {
				log.Println("error creating user: ", err.Error())
				return err
			}
			return nil
		}),
	)	
	
}