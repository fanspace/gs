package grpc

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "serverA/pb"
)

type GrpcProvider struct {
	UserServerClient pb.UserServerClient `autowire:""`
}

func (gp *GrpcProvider) QueryUsers(req *pb.UserReq) (*pb.UserListRes, error) {
	req.Limit = 3
	res, err := gp.UserServerClient.QueryUsers(context.TODO(), req)
	if err != nil {
		log.Error(err.Error())
	}
	return res, err
}

func (gp *GrpcProvider) GetUser(req *pb.UserReq) (*pb.UserRes, error) {
	res, err := gp.UserServerClient.GetUser(context.TODO(), req)
	if err != nil {
		log.Error(err.Error())
	}
	return res, err
}
