package orm

import (
	"fmt"
	"gingate/core"
	log "gingate/core"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
	"xorm.io/xorm"
)

var Orm *xorm.Engine

func InitXorm() {
	var err error
	fmt.Println(core.Cfg.Database.MysqlSettings)
	Orm, err = xorm.NewEngine(core.Cfg.Database.MysqlSettings.DriverName, core.Cfg.Database.MysqlSettings.Url)
	if err != nil {
		log.Error(err.Error())
	}
	//显示sql
	if core.Cfg.ReleaseMode {
		Orm.ShowSQL(false)
	} else {
		Orm.ShowSQL(true)
	}
	//设置时区
	Orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	// cache
	if core.Cfg.Database.MysqlSettings.UseCache {
		Orm.SetDefaultCacher(NewRedisCacher(core.Cfg.Database.RedisSettings.Addr, core.Cfg.Database.RedisSettings.Password, core.Cfg.Database.RedisSettings.DB, DEFAULT_EXPIRATION, nil))
	}
}
