package core

import (
	"context"
	"errors"
	"fmt"
	"gingate/commons"
	"google.golang.org/grpc"
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  GRPC_TOKEN_APPID,
		"appkey": GRPC_TOKEN_APPKEY,
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	/*
		if OpenTLS {
			return true
		}*/
	// 不使用tls
	return false
}

func dialgrpc(address string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		Error(err.Error())
		return nil, err
	}
	return conn, err
}

var GrpcPools map[string]*grpc.ClientConn

func initGrpcs() {
	if Cfg.GrpcSettings != nil && len(Cfg.GrpcSettings.EndPoint) > 0 {
		GrpcPools = make(map[string]*grpc.ClientConn)
		for grpcname, gprcaddr := range Cfg.GrpcSettings.EndPoint {
			fmt.Println(grpcname, gprcaddr)
			conn, err := dialgrpc(gprcaddr)
			if err != nil {
				Error(err.Error())
			}
			GrpcPools[grpcname] = conn
		}
	}
}

// 因使用k8s，为保证由k8s管理负载均衡，去除pool方式，采用单例连接
// 如需pool方式，可使用 https://github.com/shimingyah/pool
func GetGrpcPool(grpcname string) (*grpc.ClientConn, error) {
	if Cfg.DaprMode {
		grpcname = "daprrpc"
	}
	if addr, ok := Cfg.GrpcSettings.EndPoint[grpcname]; ok {
		return dialgrpc(addr)
	} else {
		return nil, errors.New(commons.CUS_ERR_4004)
	}
}
