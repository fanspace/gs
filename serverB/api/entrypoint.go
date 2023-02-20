package api

import (
	"context"
	"serverB/internal/service"
	pb "serverB/pb"
)

type ArticleService struct {
}

func (s ArticleService) GetArticle(ctx context.Context, req *pb.ArticleReq) (*pb.ArticleRes, error) {
	return service.GetArticle(req)
}

func (s ArticleService) QueryArticles(ctx context.Context, req *pb.ArticleReq) (*pb.ArticleListRes, error) {
	return service.QueryArticles(req)
}
