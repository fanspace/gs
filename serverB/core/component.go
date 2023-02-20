package core

import (
	"github.com/gomodule/redigo/redis"
	"xorm.io/xorm"
)

var Orm *xorm.Engine
var Rpool *redis.Pool

//var BCache *bigcache.BigCache
var err error

func InitComponent() {
	if Cfg.MysqlSettings != nil {
		Orm, err = initXorm()
		if err != nil {
			Error(err.Error())
		} else {
			Info("Mysql connected ... ")
		}
	}

	if Cfg.RedisSettings != nil {
		Rpool, err = initRedis()
		if err != nil {
			Error(err.Error())
		} else {
			Info("Redis connected ... ")
		}
	}

	/*if Cfg.GrpcSettings != nil {
		initGrpcs()
	}*/
	// 开启bigcache
	/*BCache, err = initLocalCache()
	if err != nil {
		Error(err.Error())
	} else {
		Info("BigCache started ... ")
	}
	if BCache != nil {
		err = loadFromRedis()
		if err != nil {
			Error(err.Error())
		} else {
			Info("BlackList added ... ")
		}
	}*/

	// 开启 casbin
	/*Casbin, err = initCasbin()
	if err != nil {
		Error(err.Error())
	} else {
		Info("BigCache init ... ")
	}*/
}

func CloseComponent() {
	Info("shutting down xrom ...")
	if Orm != nil {
		Orm.Close()
	}
	Info("shutting down redis ...")
	if Rpool != nil {
		Rpool.Close()
	}

	//Info("shutting down local cache ...")
	//if BCache != nil {
	//	BCache.Close()
	//}
}
