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

func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": commons.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4002}})
		return
	}
	req := new(pb.UserReq)
	req.Id = id
	value, err := service.DealGrpcCall(req, "GetUser", "userserver")
	if err != nil {
		ErrSolver(c, err)
		return
	}
	if value != nil {
		GrpcRessolvegr(c, value)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": commons.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4007}})
	return
}

func QueryUsers(c *gin.Context) {
	req := new(pb.UserReq)
	if err := c.ShouldBindJSON(req); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": commons.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4002}})
		return
	}
	value, err := service.DealGrpcCall(req, "QueryUsers", "userserver")
	if err != nil {
		ErrSolver(c, err)
		return
	}
	if value != nil {
		GrpcRessolvegr(c, value)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": commons.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4007}})
	return
}
