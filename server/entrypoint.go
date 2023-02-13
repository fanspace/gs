package main

import (
	"context"
	pb "server/pb"
	"server/service"
)

type UserServer struct {
	userProvider service.UserProvider `autowire:"userProvider"`
	AppName      string               `value:"${cfg.appName}"`
}

func (s *UserServer) QueryUsers(ctx context.Context, req *pb.UserReq) (*pb.UserListRes, error) {
	res, err := s.userProvider.QueryUsers(req)
	if err != nil {
		res.Msg = err.Error()
		return res, err
	}
	res.Success = true
	return res, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	res, err := s.userProvider.GetUser(req)
	if err != nil {
		res.Msg = err.Error()
		return res, err
	}
	res.Success = true
	return res, nil
}
