package grpc

import (
	"context"
	pb "gate/pb"
	"github.com/labstack/gommon/log"
)

type GrpcProvider struct {
	UserServerClient    pb.UserServerClient    `autowire:""`
	ArticleServerClient pb.ArticleServerClient `autowire:""`
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

func (gp *GrpcProvider) QueryArticles(req *pb.ArticleReq) (*pb.ArticleListRes, error) {
	req.Limit = 3
	res, err := gp.ArticleServerClient.QueryArticles(context.TODO(), req)
	if err != nil {
		log.Error(err.Error())
	}
	return res, err
}

func (gp *GrpcProvider) GetArticle(req *pb.ArticleReq) (*pb.ArticleRes, error) {
	//req.Id = 4411
	res, err := gp.ArticleServerClient.GetArticle(context.TODO(), req)
	if err != nil {
		log.Error(err.Error())
	}
	return res, err
}
