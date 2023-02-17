package service

import (
	"context"
	"errors"
	"fmt"
	"gingate/commons"
	"gingate/core"
	log "gingate/core"
	pb "gingate/pb"
	"google.golang.org/grpc/metadata"
	"reflect"
	"time"
)

func DealGrpcCall[T any](req T, methodName string, grpcName string) (any, error) {
	var res []reflect.Value
	pool, err := log.GetGrpcPool(grpcName)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer pool.Close()
	if pool == nil {
		log.Error(fmt.Sprintf("connect to %s failed", grpcName))
		return nil, errors.New(commons.CUS_ERR_4007)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*core.Cfg.GrpcSettings.TimeOut)
	defer cancel()
	if core.Cfg.DaprMode {
		ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", fmt.Sprintf("grpc-%s", grpcName))
	}
	var c any
	switch grpcName {
	case "userserver":
		c = pb.NewUserServerClient(pool)
	case "articleserver":
		c = pb.NewArticleServerClient(pool)
	default:
		return nil, errors.New(commons.CUS_ERR_4002)
	}
	value := reflect.ValueOf(c)
	f := value.MethodByName(methodName)
	var parms []reflect.Value
	parms = []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(req)}
	res = f.Call(parms)
	if len(res) != 2 {
		return nil, errors.New(commons.CUS_ERR_4007)
	}
	if res[1].Interface() != nil {
		err = res[1].Interface().(error)
		return nil, err
	}
	if res[0].Interface() != nil {
		return res[0].Interface(), nil
	}
	return nil, err
}
