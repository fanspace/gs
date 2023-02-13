package StarterMySqlXorm

import (
	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/database"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/gs/arg"
	"github.com/go-spring/spring-core/gs/cond"
	"os"
	"time"
	"xorm.io/xorm"
)

type Factory struct {
	Logger *log.Logger `logger:""`
}

func (factory *Factory) CreateDB(config database.Config) (*xorm.Engine, error) {
	db, err := xorm.NewEngine(config.Type, config.Url)
	if err != nil {
		return nil, err
	}
	//显示sql
	db.ShowSQL(config.ShowSql)

	//设置时区
	db.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	return db, nil
}

func init() {
	gs.Object(&Factory{})
	gs.Provide((*Factory).CreateDB, arg.R1("${xorm}")).
		Name("XormDB").
		On(cond.OnMissingBean(gs.BeanID((*xorm.Engine)(nil), "XormDB")))
}
