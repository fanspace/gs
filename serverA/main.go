package main

import (
	"fmt"
	"github.com/go-spring/spring-core/grpc"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/gs/cond"
	_ "github.com/go-spring/starter-grpc/server"
	pb "serverA/pb"
)

func init() {
	// 创建service.ProductService bean示例并设置初始方法
	gs.Object(new(ArticleServer)).On(cond.OnBean("wr-db").OnBean("ArticleProvider")).Init(func(srv *ArticleServer) {
		//gs.Object(new(service.MyServerServer)).Init(func(srv *service.MyServerServer) {
		// 添加grpc服务
		// gs.GrpcServer(string, *grpc.Server) 其中第一个参数为serviceName, 要求与生成的pb文件Service.Desc中的ServiceName参数值相等
		gs.GrpcServer("myserver.ArticleServer", &grpc.Server{
			// 服务注册方法
			Register: pb.RegisterArticleServerServer,
			// 服务实现对象
			Service: srv,
		})
	})
}

func main() {
	gs.Property("spring.application.name", "${cfg.appName}")
	gs.Property("grpc.server.port", "${cfg.port}")
	fmt.Println("application exit: ", gs.Web(false).Run())
}
