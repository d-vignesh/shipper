package main

import (
	"golang.org/x/net/context"
	pb "github.com/d-vignesh/shipper/user-service/proto/user"
)

type UserHandler struct {
	repo Repository
	// tokenService Authable
}

func (uh *UserHandler) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	user, err := uh.repo.Get(req.Id)
	if err != nil {
		return err
	}
	resp.User = user
	return nil
}

func (uh *UserHandler) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	users, err := uh.repo.GetAll()
	if err != nil {
		return err
	}
	resp.Users = users
	return nil
}

func (uh *UserHandler) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	_, err := uh.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	resp.Token = "testingtok"
	return nil
}

func (uh *UserHandler) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	err := uh.repo.Create(req)
	if err != nil {
		return err
	}
	resp.User = req
	return nil
}

func (uh *UserHandler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	return nil
}