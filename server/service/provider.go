package service

import (
	"github.com/go-spring/spring-core/gs"
	pb "server/pb"
)

func init() {
	gs.Object(new(UserService)).
		Export((*UserProvider)(nil)).Name("userProvider")
}

type UserProvider interface {
	QueryUsers(req *pb.UserReq) (*pb.UserListRes, error)
	GetUser(req *pb.UserReq) (*pb.UserRes, error)
}
