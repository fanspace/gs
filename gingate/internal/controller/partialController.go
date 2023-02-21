package controller

import (
	"gingate/commons"
	log "gingate/core"
	"gingate/internal/model"
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

func GrpcResSolver(c *gin.Context, value any) {

	if value != nil {
		c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": value})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4004}})
	return
}
