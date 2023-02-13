package user

import (
	"fmt"
	"gate/grpc"
	pb "gate/pb"
	"github.com/go-spring/spring-core/web"
	"github.com/labstack/gommon/log"
	"strconv"
)

type Controller struct {
	gp *grpc.GrpcProvider `autowire:"gp"`
}

func (c *Controller) QueryUsers(ctx web.Context) {
	req := new(pb.UserReq)
	if err := ctx.Bind(req); err != nil {
		log.Error(err.Error())
		ctx.String("Wrong params \n")
		return
	}
	res, err := c.gp.QueryUsers(req)
	if err != nil {
		ctx.String(err.Error())
		return
	}
	if res.Success {
		for _, v := range res.Users {
			fmt.Println(v.Showname)
		}
		ctx.String("succeed in query Users ! \n")
		return
	} else {
		ctx.String(res.Msg)
		return
	}
}

func (c *Controller) GetUser(ctx web.Context) {
	req := new(pb.UserReq)
	req.Id, _ = strconv.ParseInt(ctx.PathParam("id"), 10, 64)
	res, err := c.gp.GetUser(req)
	if err != nil {
		ctx.String(err.Error())
		return
	}
	if res.Success {
		ctx.String(fmt.Sprintf("succeed in get User : %s \n", res.User.Showname))
		return
	}
	ctx.String(fmt.Sprintf("failed to get User  due to : %s", res.Msg))
	return
}
