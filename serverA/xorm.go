package main

import (
	"github.com/go-spring/spring-core/gs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"os"
	"serverA/model"
	"time"
	"xorm.io/xorm"
)

func init() {
	// 创建一个配置监听器, 当配置加载后将执行该方法

	gs.OnProperty("dbs", func(config []model.DBConfig) {
		// 遍历配置项
		for _, dbConfig := range config {
			db, err := CreateDB(dbConfig)
			if err != nil {
				log.Fatal(err)
			}
			// 将*gorm.DB注入Spring bean中
			gs.Object(db).Destroy(CloseDB).Name(dbConfig.Name)
		}
	})

}

func CreateDB(dbconfig model.DBConfig) (*xorm.Engine, error) {
	db, err := xorm.NewEngine(dbconfig.Type, dbconfig.Url)
	if err != nil {
		return nil, err
	}
	//显示sql
	db.ShowSQL(dbconfig.ShowSql)

	//设置时区
	db.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	return db, nil
}

func CloseDB(db *xorm.Engine) {
	log.Info("close xorm mysql")
	if err := db.Close(); err != nil {
		log.Error(err)
	}
}
