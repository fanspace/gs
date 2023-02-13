package controller

import (
	"gate/controller/pub"
	"gate/controller/user"
	"github.com/go-spring/spring-core/gs"
)

func init() {

	gs.Object(new(pub.Controller)).Init(func(c *pub.Controller) {
		// 注册路由
		gs.GetMapping("/", c.Home)
	})
	gs.Object(new(user.Controller)).Init(func(c *user.Controller) {
		// 注册路由
		gs.GetMapping("/auth/users/query", c.QueryUsers)
	})
}
