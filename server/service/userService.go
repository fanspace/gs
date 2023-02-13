package service

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
	"server/model"
	pb "server/pb"
	"xorm.io/xorm"
)

type UserService struct {
	db *xorm.Engine `autowire:"wr-db"`
}

func (s *UserService) GetUser(req *pb.UserReq) (*pb.UserRes, error) {
	res := new(pb.UserRes)
	user := new(model.User)
	_, err := s.db.ID(req.Id).Get(user)
	if err != nil {
		log.Error(err.Error())
		return res, err
	}
	res.User = new(pb.User)
	copier.Copy(res.User, user)
	return res, nil
}

func (s *UserService) QueryUsers(req *pb.UserReq) (*pb.UserListRes, error) {
	res := new(pb.UserListRes)
	users := make([]*model.User, 0)
	err := s.db.Limit(int(req.Limit)).Find(&users)
	if err != nil {
		log.Error(err.Error())
		return res, err
	}
	res.Users = make([]*pb.User, 0)
	for _, v := range users {
		it := new(pb.User)
		copier.Copy(it, v)
		res.Users = append(res.Users, it)
	}

	return res, nil
}
