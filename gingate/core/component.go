package core

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/casbin/casbin/v2"
	"github.com/gomodule/redigo/redis"
	"xorm.io/xorm"
)

var Orm *xorm.Engine
var Rpool *redis.Pool
var BCache *bigcache.BigCache
var Casbin *casbin.Enforcer
var err error

func InitComponent() {
	if Cfg.Database.MysqlSettings != nil {
		Orm, err = initXorm()
		if err != nil {
			Error(err.Error())
		} else {
			Info("Mysql connected ... ")
		}
	}

	if Cfg.Database.RedisSettings != nil {
		Rpool, err = initRedis()
		if err != nil {
			Error(err.Error())
		} else {
			Info("Redis connected ... ")
		}
	}

	if Cfg.GrpcSettings != nil {
		initGrpcs()
	}
	// 开启bigcache
	BCache, err = initLocalCache()
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

	// 开启 casbin
	Casbin, err = initCasbin()
	if err != nil {
		Error(err.Error())
	} else {
		Info("BigCache init ... ")
	}
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

func loadFromRedis() error {
	total, blacks, err := getAllBanUsers()
	if err != nil {
		Error(err.Error())
		return err
	} else {
		if total > 0 {
			for _, v := range blacks {
				BanUserCache(v)
			}
		}
	}
	return nil
}

func BanUserCache(username string) error {
	BCache.Set(fmt.Sprintf("%s_%s", PREFIX_BCACHE_BAN, username), []byte(PREFIX_BCACHE_BAN))
	Info("username： " + username + "已被加入黑名单")
	return nil
}

func IsUserBaned(username string) bool {
	entry, _ := BCache.Get(fmt.Sprintf("%s_%s", PREFIX_BCACHE_BAN, username))
	if string(entry) != "" {
		return true
	}
	return false
}

func getAllBanUsers() (int64, []string, error) {
	myredis := Rpool.Get()
	defer myredis.Close()
	res := make([]string, 0)
	keyname := PREFIX_REDIS_USER_BAN
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
