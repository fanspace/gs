package engine

import (
	"fmt"
	"gate/controller"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	Engine     *gin.Engine
	Address    string                 `value:"${http.addr:=:8080}"`
	Controller *controller.Controller `autowire:"controllers"`
	Exit       chan struct{}          `autowire:""`
}

func (e *Engine) Init() {
	e.Engine = gin.Default()
	//e.Engine.GET("/article/:id", e.Controller.ArticleCtrl.GetArticle)
	go func() {
		err := e.Engine.Run(e.Address)
		fmt.Println(err)
		e.Exit <- struct{}{}
	}()
}
