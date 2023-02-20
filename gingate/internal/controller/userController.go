package controller

import (
	"gingate/commons"
	"gingate/core"
	log "gingate/core"
	"gingate/internal/model"
	"gingate/internal/service"
	pb "gingate/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary 用户详情
// @Description 用户详情
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true " "
// @Success 200 {object} map[string]any
// @Router /user/{id} [get]
func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: core.Cfg.BizErr["err_4002"]}})
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
		GrpcResSolver(c, value)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4007}})
	return
}

// @Summary 用户列表
// @Description 用户列表
// @Tags User
// @Accept json
// @Produce json
// @Param req body pb.UserReq true "User list"
// @Success 200 {object} map[string]any
// @Router /user/list [post]
func QueryUsers(c *gin.Context) {
	req := new(pb.UserReq)
	if err := c.ShouldBindJSON(req); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4002}})
		return
	}
	value, err := service.DealGrpcCall(req, "QueryUsers", "userserver")
	if err != nil {
		ErrSolver(c, err)
		return
	}
	if value != nil {
		GrpcResSolver(c, value)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": log.WEB_STATUS_BACK, "result": &model.SimpleResponse{Success: false, Msg: commons.CUS_ERR_4007}})
	return
}
