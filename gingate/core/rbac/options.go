package rbac

import (
	"fmt"
	rds "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type WatcherOptions struct {
	rds.Options
	Channel    string
	IgnoreSelf bool
	LocalID    string
	Password   string
}

func initConfig(option *WatcherOptions) {
	if option.LocalID == "" {
		option.LocalID = uuid.New().String()
	}
	if option.Channel == "" {
		option.Channel = fmt.Sprintf("/%s_rbac", "Unknown")
	}
}
