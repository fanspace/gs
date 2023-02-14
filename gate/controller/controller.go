package controller

import (
	"gate/controller/article"
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
		gs.PostMapping("/auth/users/query", c.QueryUsers)
		gs.GetMapping("/auth/user/:id", c.GetUser)
	})

	gs.Object(new(article.Controller)).Init(func(c *article.Controller) {
		// 注册路由
		gs.PostMapping("/articles/query", c.QueryArticles)
		gs.GetMapping("/article/:id", c.GetArticle)
	})

}

type Controller struct {
	ArticleCtrl article.Controller
	UserCtrl    user.Controller
	PubCtrl     pub.Controller
}
