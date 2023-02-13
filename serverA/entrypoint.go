package main

import (
	"context"
	pb "serverA/pb"
	"serverA/service"
)

type ArticleServer struct {
	ArticleProvider service.ArticleProvider `autowire:"ArticleProvider"`
	AppName         string                  `value:"${cfg.appName}"`
}

func (s *ArticleServer) QueryArticles(ctx context.Context, req *pb.ArticleReq) (*pb.ArticleListRes, error) {
	res, err := s.ArticleProvider.QueryArticles(req)
	if err != nil {
		res.Msg = err.Error()
		return res, err
	}
	res.Success = true
	return res, nil
}

func (s *ArticleServer) GetArticle(ctx context.Context, req *pb.ArticleReq) (*pb.ArticleRes, error) {
	res, err := s.ArticleProvider.GetArticle(req)
	if err != nil {
		res.Msg = err.Error()
		return res, err
	}
	res.Success = true
	return res, nil
}
