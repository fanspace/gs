package main

import (
	_ "gate/controller"
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
	gs.Property("grpc.endpoint.userserver.address", "127.0.0.1:12100")
	gs.Property("grpc.endpoint.articleserver.address", "127.0.0.1:12200")
	//gs.Property("spring.application.name", "GreeterClient")
	//fmt.Println("application exit: ", gs.Web(false).Run())
	log.Fatal(gs.Run())
}
