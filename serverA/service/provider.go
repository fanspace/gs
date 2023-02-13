package service

import (
	"github.com/go-spring/spring-core/gs"
	pb "serverA/pb"
)

func init() {
	//gs.Object(new(ArticleService)).On(cond.OnBean((*xorm.Engine)(nil))).Name("ArticleService")
	gs.Object(new(ArticleService)).
		Export((*ArticleProvider)(nil)).Name("ArticleProvider")
}

type ArticleProvider interface {
	QueryArticles(req *pb.ArticleReq) (*pb.ArticleListRes, error)
	GetArticle(req *pb.ArticleReq) (*pb.ArticleRes, error)
}
