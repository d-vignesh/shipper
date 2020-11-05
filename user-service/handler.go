package main

import (
	"context"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	pb "github.com/d-vignesh/shipper/user-service/proto/user"
)

type authable inteface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type UserHandler struct {
	repository 		Repository
	tokenService 	authable
}

func (uh *UserHandler) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	results, err := uh.repository.Get(ctx)
	if err != nil {
		return err
	}

	users := UnmarshalUserCollection(results)
	resp.Users = users
	return nil
}

func (uh *UserHandler) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	result, err := uh.repository.Get(ctx, req.Id)
	if err != nil {
		return err
	}

	user := UnmarshalUser(result)
	resp.User = user

	return nil
}

func (uh *UserHandler) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	log.Println("user : ", req)
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPass)
	if err := vh.repository.Create(ctx, MarshalUser(req)); err != nil {
		return err
	}

	req.Password = ""
	resp.User = req

	return nil
}

func (uh *UserHandler) Auth(ctx context.Context, req *pb.User, resp *pb.Response) error {
	user, err := uh.repository.GetByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := uh.tokenService.Encode(req)
	if err != nil {
		return err
	}

	resp.Token = token
	return nil
}

func (uh *UserHandler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	claims, err := uh.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	resp.Valid = true
	return nil
}