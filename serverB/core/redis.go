package core

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

func initRedis() (*redis.Pool, error) {
	pool := &redis.Pool{
		MaxIdle:     Cfg.RedisSettings.MaxIdle,
		MaxActive:   Cfg.RedisSettings.MaxActive,
		IdleTimeout: time.Duration(Cfg.RedisSettings.IdleTimeout) * time.Second,
		// Other pool configuration not shown in this example.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", Cfg.RedisSettings.Addr)
			if err != nil {
				Error(err.Error())
				return nil, err
			}
			if _, err := c.Do("AUTH", Cfg.RedisSettings.Password); err != nil {
				c.Close()
				Error(err.Error())
				return nil, err
			}
			if _, err := c.Do("SELECT", Cfg.RedisSettings.DB); err != nil {
				c.Close()
				Error(err.Error())
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return pool, nil
}
