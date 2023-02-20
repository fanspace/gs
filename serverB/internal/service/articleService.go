package service

import (
	"fmt"
	"github.com/jinzhu/copier"
	"serverB/core"
	log "serverB/core"
	"serverB/internal/model"
	pb "serverB/pb"
)

func GetArticle(req *pb.ArticleReq) (*pb.ArticleRes, error) {
	res := new(pb.ArticleRes)
	Article := new(model.Article)
	_, err := core.Orm.ID(req.Id).Get(Article)
	if err != nil {
		log.Error(err.Error())
		return res, err
	}
	res.Article = new(pb.Article)
	copier.Copy(res.Article, Article)
	//测试配置热更新
	fmt.Println(core.Cfg.Smark)
	return res, nil
}

func QueryArticles(req *pb.ArticleReq) (*pb.ArticleListRes, error) {
	res := new(pb.ArticleListRes)
	Articles := make([]*model.Article, 0)
	err := core.Orm.Limit(int(req.Limit)).Find(&Articles)
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
