package article

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

func (c *Controller) QueryArticles(ctx web.Context) {
	req := new(pb.ArticleReq)
	if err := ctx.Bind(req); err != nil {
		log.Error(err.Error())
		ctx.String("Wrong params \n")
		return
	}
	res, err := c.gp.QueryArticles(req)
	if err != nil {
		ctx.String(err.Error())
		return
	}
	if res.Success {
		for _, v := range res.Articles {
			fmt.Println(v.Title)
		}
		ctx.String("succeed in query articles ! \n")
		return
	} else {
		ctx.String(res.Msg)
		return
	}

}

func (c *Controller) GetArticle(ctx web.Context) {
	req := new(pb.ArticleReq)
	req.Id, _ = strconv.ParseInt(ctx.PathParam("id"), 10, 64)
	fmt.Println(req)
	res, err := c.gp.GetArticle(req)
	if err != nil {
		log.Error(err.Error())
		ctx.String(err.Error())
		return
	}
	if res.Success {
		ctx.String(fmt.Sprintf("succeed in get article : %s \n", res.Article.Title))
		return
	}
	ctx.String(fmt.Sprintf("failed to get article  due to : %s", res.Msg))
	return
}
