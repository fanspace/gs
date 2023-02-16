package service

import (
	"fmt"
	log "gingate/core"
)

func BanUserCache(username string) error {
	log.BCache.Set(fmt.Sprintf("ban_%s", username), []byte("ban"))
	log.Info("username： " + username + "已被加入黑名单")
	return nil
}

func ReleaseUserCache(username string) error {
	log.BCache.Delete(fmt.Sprintf("ban_%s", username))
	log.Info("username： " + username + "已从黑名单中删除")
	return nil
}

func IsUserBaned(username string) bool {
	entry, _ := log.BCache.Get(fmt.Sprintf("ban_%s", username))
	if string(entry) != "" {
		return true
	}
	return false
}

func BanIpCache(ip string) error {
	log.BCache.Set(fmt.Sprintf("ban_%s", ip), []byte("ban"))
	log.Info("ip： " + ip + "已被加入黑名单，持续1小时")
	return nil
}

func ReleaseIpCache(ip string) error {
	log.BCache.Delete(fmt.Sprintf("ban_%s", ip))
	log.Info("ip： " + ip + "已从黑名单中删除")
	return nil
}
func IsIpBand(ip string) bool {
	entry, _ := log.BCache.Get(fmt.Sprintf("ban_%s", ip))
	if string(entry) != "" {
		return true
	}
	return false
}

func SetBigCache(key string, val string) error {
	log.BCache.Set(key, []byte(val))
	return nil
}
func GetBigCache(key string) (string, error) {
	r, err := log.BCache.Get(key)
	return string(r), err
}
