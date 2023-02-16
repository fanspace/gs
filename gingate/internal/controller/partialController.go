package controller

import (
	"gingate/commons"
	log "gingate/core"
	"gingate/internal/model"
	pb "gingate/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ErrSolver(c *gin.Context, err error) {
	if strings.Index(err.Error(), "desc =") > 0 {
		log.Error(strings.Split(err.Error(), "desc = ")[1])
		c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4007}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: err.Error()}})
	return
}

func GrpcRessolvegr(c *gin.Context, value any) {
	var res any
	switch value.(type) {
	case *pb.UserListRes:
		res = value.(*pb.UserListRes)
	case *pb.UserRes:
		res = value.(*pb.UserRes)
	case *pb.ArticleRes:
		res = value.(*pb.ArticleRes)
	case *pb.ArticleListRes:
		res = value.(*pb.ArticleListRes)
	default:
		c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4002}})
		return
	}
	if res != nil {
		c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": res})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4004}})
	return
}
