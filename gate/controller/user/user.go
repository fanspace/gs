package user

import "github.com/go-spring/spring-core/web"

type Controller struct {
}

func (c *Controller) QueryUsers(ctx web.Context) {
	ctx.String("This is query users! \n")
}
