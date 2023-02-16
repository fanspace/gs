package core

import (
	"encoding/json"
	"fmt"
	"gingate/core/rbac"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/xorm-adapter/v2"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

func initCasbin() (*casbin.Enforcer, error) {
	a, err := xormadapter.NewAdapter("mysql", Cfg.Database.MysqlSettings.Url, true)
	if err != nil {
		Error(err.Error())
		return nil, err
	}
	casbin, err := casbin.NewEnforcer("config/authz_model.conf", a)
	if err != nil {
		Error(err.Error())
		return nil, err
	}
	if !Cfg.ReleaseMode {
		casbin.EnableLog(true)
	} else {
		casbin.EnableLog(false)
	}
	wc := rbac.WatcherOptions{
		Options: redis.Options{
			Network:  "tcp",
			Password: Cfg.Database.RedisSettings.Password,
		},
		Channel:    fmt.Sprintf("/%s_rbac", Cfg.AppName),
		LocalID:    Cfg.Smark,
		IgnoreSelf: true,
	}
	w, err := rbac.NewWatcher(Cfg.Database.RedisSettings.Addr, wc)
	if err != nil {
		Error(err.Error())
		return nil, err
	}
	casbin.SetWatcher(w)
	w.SetUpdateCallback(updateCallback)
	initCasbinRoot(casbin)
	return casbin, nil
}

func initCasbinRoot(casbin *casbin.Enforcer) error {
	cas := &rbac.CasbinRule{
		Id: 1,
	}
	hasdata, err := Orm.Exist(cas)
	if err != nil {
		Error(err.Error())
		return err
	}
	if hasdata {
		return nil
	} else {
		caslist := make([]*rbac.CasbinRule, 0)
		casarr := [...]*rbac.CasbinRule{
			&rbac.CasbinRule{
				Id:    1,
				PType: "p",
				V0:    "root",
				V1:    "back",
				V2:    "*",
				V3:    "*",
			}, &rbac.CasbinRule{
				Id:    2,
				PType: "g",
				V0:    "rootAdm",
				V1:    "root",
				V2:    "back",
			},
		}
		caslist = casarr[:]
		_, err := Orm.Insert(&caslist)
		if err != nil {
			Error(err.Error())
			return err
		}
		casbin.AddNamedGroupingPolicy("p", "root", "back", "*", "*")
		casbin.AddRoleForUserInDomain("rootAdm", "root", "back")
	}
	return nil
}

func updateCallback(msg string) {
	timea := time.Now()
	begin := timea.UnixNano()
	sm := Cfg.Smark
	msgcon := strings.Replace(msg, `\`, "", -1)
	msgs := new(rbac.CasbinUpdateMsg)
	err := json.Unmarshal([]byte(msgcon), msgs)
	if err != nil {
		Error(err.Error())
		err := Casbin.LoadPolicy()
		if err != nil {
			Error(err.Error())
		}
		end := time.Now().UnixNano()
		Info(fmt.Sprintf("********更新casbin规则：sender is %s - reciever is %s  共耗时 ********   %d ms", msg, sm, (end-begin)/1000/1000))
	} else if msgs.ID == "" {
		err := Casbin.LoadPolicy()
		if err != nil {
			Error(err.Error())
		}
		end := time.Now().UnixNano()
		Info(fmt.Sprintf("********更新casbin规则：sender is %s - reciever is %s  共耗时 ********   %d ms", msg, sm, (end-begin)/1000/1000))
	} else {
		if sm == msgs.ID {
			sm = "self"
			end := time.Now().UnixNano()
			Info(fmt.Sprintf("********更新casbin规则：sender is %s - %s  共耗时 ********   %d ms", msgs.ID, sm, (end-begin)/1000/1000))
		} else {
			err := Casbin.LoadPolicy()
			if err != nil {
				Error(err.Error())
			}
			end := time.Now().UnixNano()
			Info(fmt.Sprintf("********更新casbin规则：sender is %s - reciever is %s  共耗时 ********   %d ms", msgs.ID, sm, (end-begin)/1000/1000))
		}
	}

}
