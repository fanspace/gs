package api

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"serverB/core"
	log "serverB/core"
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

func TokenInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		log.Error("---------> the unaryServerInterceptor: " + info.FullMethod)

		err := authToken(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

// 简单实现，用常量替代，这里如果不使用tls,还是jwt或自己实现加密更好些
func authToken(ctx context.Context, fullmethod string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("客户端校验失败")
	}
	var (
		appid  string
		appkey string
	)
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}

	if appid != core.GRPC_TOKEN_APPID || appkey != core.GRPC_TOKEN_APPKEY {
		return errors.New("Token认证信息无效")
	}
	return nil
}
