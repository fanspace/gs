package core

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

func InitXorm() (*xorm.Engine, error) {

	orm, err := xorm.NewEngine(Cfg.Database.MysqlSettings.DriverName, Cfg.Database.MysqlSettings.Url)
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
	if Cfg.Database.MysqlSettings.UseCache {
		orm.SetDefaultCacher(NewRedisCacher(Cfg.Database.RedisSettings.Addr, Cfg.Database.RedisSettings.Password, Cfg.Database.RedisSettings.DB, DEFAULT_EXPIRATION, nil))
	}
	return orm, nil
}
