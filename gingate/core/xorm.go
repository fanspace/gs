package core

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

func initXorm() (*xorm.Engine, error) {

	orm, err := xorm.NewEngine(Cfg.MysqlSettings.DriverName, Cfg.MysqlSettings.Url)
	if err != nil {
		Error(err.Error())
		return nil, err
	}
	//显示sql
	if Cfg.ReleaseMode {
		orm.ShowSQL(false)
	} else {
		orm.ShowSQL(true)
	}
	//设置时区
	orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		Error(err.Error())
		return nil, err
	}
	// cache
	if Cfg.MysqlSettings.UseCache {
		orm.SetDefaultCacher(NewRedisCacher(Cfg.RedisSettings.Addr, Cfg.RedisSettings.Password, Cfg.RedisSettings.DB, DEFAULT_EXPIRATION, nil))
	}
	return orm, nil
}
