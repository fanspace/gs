package main

import (
	"context"
	pb "server/pb"
	"server/service"
)

type MyServerServer struct {
	userProvider service.UserProvider `autowire:"userProvider"`
	AppName      string               `value:"${cfg.appName}"`
}

func (s *MyServerServer) QueryUsers(ctx context.Context, req *pb.UserReq) (*pb.UserListRes, error) {
	res, err := s.userProvider.QueryUsers(req)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *MyServerServer) GetUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	res, err := s.userProvider.GetUser(req)
	if err != nil {
		return res, err
	}
	return res, nil
}
