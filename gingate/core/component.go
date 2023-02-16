package core

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/gomodule/redigo/redis"
	"xorm.io/xorm"
)

var Orm *xorm.Engine
var Rpool *redis.Pool
var BCache *bigcache.BigCache
var err error

func InitComponent() {
	if Cfg.Database.MysqlSettings != nil {
		Orm, err = InitXorm()
		if err != nil {
			Error(err.Error())
		} else {
			Info("Mysql connected ... ")
		}
	}

	if Cfg.Database.RedisSettings != nil {
		Rpool, err = InitRedis()
		if err != nil {
			Error(err.Error())
		} else {
			Info("Redis connected ... ")
		}
	}

	if Cfg.GrpcSettings != nil {
		InitGrpcs()
	}
	// 开启bigcache
	BCache, err = InitLocalCache()
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
	}
}

func loadFromRedis() error {
	total, blacks, err := getAllBanUsers()
	if err != nil {
		Error(err.Error())
		return err
	} else {
		if total > 0 {
			for _, v := range blacks {
				banUserCache(v)
			}
		}
	}
	return nil
}

func banUserCache(username string) error {
	BCache.Set(fmt.Sprintf("ban_%s", username), []byte("ban"))
	Info("username： " + username + "已被加入黑名单")
	return nil
}

func getAllBanUsers() (int64, []string, error) {
	myredis := Rpool.Get()
	defer myredis.Close()
	res := make([]string, 0)
	keyname := "zj_banuser"
	total, err := redis.Int64(myredis.Do("SCARD", keyname))
	if err != nil {
		Error(err.Error())
		return 0, res, err
	}
	if total > 0 {
		res, err = redis.Strings(myredis.Do("SMEMBERS", keyname))
		if err != nil {
			Error(err.Error())
			return 0, res, err
		}
	}
	return total, res, nil
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
	Info("shutting down local cache ...")
	if BCache != nil {
		BCache.Close()
	}
}
