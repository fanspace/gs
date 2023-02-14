package service

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
	"serverA/grpc"
	"serverA/model"
	pb "serverA/pb"
	"xorm.io/xorm"
)

type ArticleService struct {
	db *xorm.Engine       `autowire:"wr-db"`
	gp *grpc.GrpcProvider `autowire:"gp"`
}

func (s *ArticleService) GetArticle(req *pb.ArticleReq) (*pb.ArticleRes, error) {
	res := new(pb.ArticleRes)
	Article := new(model.Article)
	_, err := s.db.ID(req.Id).Get(Article)
	if err != nil {
		log.Error(err.Error())
		return res, err
	}
	res.Article = new(pb.Article)
	copier.Copy(res.Article, Article)
	userreq := &pb.UserReq{Id: 1, Limit: 0}
	user, err := s.gp.GetUser(userreq)
	if err != nil {
		log.Error(err.Error())
		return res, err
	}
	fmt.Println(user.User.Showname)
	return res, nil
}

func (s *ArticleService) QueryArticles(req *pb.ArticleReq) (*pb.ArticleListRes, error) {
	res := new(pb.ArticleListRes)
	Articles := make([]*model.Article, 0)
	err := s.db.Limit(int(req.Limit)).Find(&Articles)
	if err != nil {
		log.Error(err.Error())
		return res, err
	}
	res.Articles = make([]*pb.Article, 0)
	for _, v := range Articles {
		it := new(pb.Article)
		copier.Copy(it, v)
		res.Articles = append(res.Articles, it)
	}

	return res, nil
}
