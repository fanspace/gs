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

// @Summary 文章详情
// @Description 文章详情
// @Tags Article
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true " "
// @Success 200 {object} map[string]any
// @Router /article/{id} [get]
func GetArticle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4002}})
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
		GrpcRessolvegr(c, value)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4007}})
	return
}

// @Summary 文章列表
// @Description 文章列表
// @Tags Article
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param req body pb.ArticleReq true "article list"
// @Success 200 {object} map[string]any
// @Router /article/list [post]
func QueryArticles(c *gin.Context) {
	req := new(pb.ArticleReq)
	if err := c.ShouldBindJSON(req); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4002}})
		return
	}
	value, err := service.DealGrpcCall(req, "QueryArticles", "articleserver")
	if err != nil {
		ErrSolver(c, err)
		return
	}
	if value != nil {
		GrpcRessolvegr(c, value)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4007}})
	return
}
