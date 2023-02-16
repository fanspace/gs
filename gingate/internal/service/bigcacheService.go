package service

import (
	"fmt"
	"gingate/core"
	log "gingate/core"
)

func BanUserCache(username string) error {
	return core.BanUserCache(username)
}

func ReleaseUserCache(username string) error {
	core.BCache.Delete(fmt.Sprintf("%s_%s", core.PREFIX_BCACHE_BAN, username))
	log.Info("username： " + username + "已从黑名单中删除")
	return nil
}

func IsUserBaned(username string) bool {
	entry, _ := core.BCache.Get(fmt.Sprintf("%s_%s", core.PREFIX_BCACHE_BAN, username))
	if string(entry) != "" {
		return true
	}
	return false
}

func BanIpCache(ip string) error {
	core.BCache.Set(fmt.Sprintf("%s_%s", core.PREFIX_BCACHE_BAN, ip), []byte(core.PREFIX_BCACHE_BAN))
	log.Info("ip： " + ip + "已被加入黑名单，持续1小时")
	return nil
}

func ReleaseIpCache(ip string) error {
	core.BCache.Delete(fmt.Sprintf("%s_%s", core.PREFIX_BCACHE_BAN, ip))
	log.Info("ip： " + ip + "已从黑名单中删除")
	return nil
}
func IsIpBand(ip string) bool {
	entry, _ := core.BCache.Get(fmt.Sprintf("%s_%s", core.PREFIX_BCACHE_BAN, ip))
	if string(entry) != "" {
		return true
	}
	return false
}

func SetBigCache(key string, val string) error {
	core.BCache.Set(key, []byte(val))
	return nil
}
func GetBigCache(key string) (string, error) {
	r, err := core.BCache.Get(key)
	return string(r), err
}
