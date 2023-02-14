package main

import (
	gp "gate/grpc"
	pb "gate/pb"
	"github.com/go-spring/spring-core/gs"
	_ "github.com/go-spring/starter-gin"
	_ "github.com/go-spring/starter-grpc/client"
	"github.com/labstack/gommon/log"
)

func init() {

	gs.GrpcClient(pb.NewUserServerClient, "userserver")
	gs.GrpcClient(pb.NewArticleServerClient, "articleserver")
	gs.Object(&gp.GrpcProvider{}).Name("gp")

}

func main() {
	gs.Property("grpc.endpoint.userserver.address", "${grpcSettings.userServer}")
	gs.Property("grpc.endpoint.articleserver.address", "${grpcSettings.articleServer}")

	log.Fatal(gs.Run())

}
