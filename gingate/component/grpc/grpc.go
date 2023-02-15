package grpc

import (
	"errors"
	"fmt"
	"gingate/commons"
	"gingate/core"
	log "gingate/core"
	"google.golang.org/grpc"
)

func dialgrpc(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return conn, err
}

var GrpcPools map[string]*grpc.ClientConn

func InitGrpcs() {
	if core.Cfg.GrpcSettings != nil && len(core.Cfg.GrpcSettings.EndPoint) > 0 {
		GrpcPools = make(map[string]*grpc.ClientConn)
		for grpcname, gprcaddr := range core.Cfg.GrpcSettings.EndPoint {
			fmt.Println(grpcname, gprcaddr)
			conn, err := dialgrpc(gprcaddr)
			if err != nil {
				log.Error(err.Error())
			}
			GrpcPools[grpcname] = conn
		}
	}
}

// 因使用k8s，为保证由k8s管理负载均衡，去除pool方式，采用单例连接
// 如需pool方式，可使用 https://github.com/shimingyah/pool
func GetGrpcPool(grpcname string) (*grpc.ClientConn, error) {
	if core.Cfg.DaprMode {
		grpcname = "daprrpc"
	}
	if addr, ok := core.Cfg.GrpcSettings.EndPoint[grpcname]; ok {
		return dialgrpc(addr)
	} else {
		return nil, errors.New(commons.CUS_ERR_4004)
	}
}
