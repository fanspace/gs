package controller

import (
	"gingate/commons"
	log "gingate/core"
	"gingate/internal/model"
	"gingate/internal/service"
	pb "gingate/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetArticle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": commons.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4002}})
		return
	}
	req := new(pb.ArticleReq)
	req.Id = id

	value, err := service.DealGrpcCall(req, "GetArticle", "articleserver")
	if err != nil {
		ErrSolver(c, err)
		return
	}
	if value != nil {
		res := value.(*pb.ArticleRes)
		if res != nil {
			c.JSON(http.StatusOK, gin.H{"code": commons.WEB_STATUS_BACK, "result": res})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": commons.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4007}})
	return
}

func QueryArticles(c *gin.Context) {
	req := new(pb.ArticleReq)
	if err := c.ShouldBindJSON(req); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": commons.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4002}})
		return
	}
	value, err := service.DealGrpcCall(req, "QueryArticles", "articleserver")
	if err != nil {
		ErrSolver(c, err)
		return
	}
	if value != nil {
		res := value.(*pb.ArticleListRes)
		if res != nil {
			c.JSON(http.StatusOK, gin.H{"code": commons.WEB_STATUS_BACK, "result": res})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": commons.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4007}})
	return
}
