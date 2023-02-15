package component

import (
	"gingate/component/database/orm"
	"gingate/component/database/redis"
	"gingate/component/grpc"
	"gingate/core"
)

func InitComponent() {
	if core.Cfg.Database.MysqlSettings != nil {
		orm.InitXorm()
	}

	if core.Cfg.Database.RedisSettings != nil {
		redis.InitRedis()
	}

	if core.Cfg.GrpcSettings != nil {
		grpc.InitGrpcs()
	}
}

func CloseComponent() {
	if orm.Orm != nil {
		orm.Orm.Close()
	}
	if redis.Rpool != nil {
		redis.Rpool.Close()
	}
}
