package main

import (
	"google.golang.org/grpc"
	"net"
	"os"
	"serverB/api"
	"serverB/core"
	log "serverB/core"
	pb "serverB/pb"
)

func main() {
	var ArticleService = api.ArticleService{}
	listen, err := net.Listen("tcp", "0.0.0.0:"+core.Cfg.GrpcSettings.Port)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	s := grpc.NewServer()
	pb.RegisterArticleServerServer(s, ArticleService)
	log.Info("Account Grpc Server Starting on " + core.Cfg.GrpcSettings.Port)
	core.InitComponent()
	defer core.CloseComponent()
	s.Serve(listen)
}
