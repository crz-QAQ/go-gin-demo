package redis_api

import (
	"go-gin-demo/pkg/response"
	"go-gin-demo/service"

	"github.com/gin-gonic/gin"
)

// SetRedisApi 设置redis
func SetRedisApi(c *gin.Context) {
	type Param struct {
		ID    int    `form:"id" binding:"required"`
		Value string `form:"value" binding:"required"`
	}

	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}

	err := service.SetUser(param.ID, param.Value)
	if err != nil {
		response.Error(c, "redis设置失败", err)
		return
	}
	response.Success(c, true, "redis设置成功")
}

// GetRedisApi 测试获取redis
func GetRedisApi(c *gin.Context) {
	type Param struct {
		ID int `form:"id" binding:"required"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	user, err := service.GetUser(param.ID)
	if err != nil {
		response.Error(c, "获取redis的value失败", err)
		return
	}
	response.Success(c, user, "redis获取成功")
}

// DeleRedisApi 测试删除redis
func DeleRedisApi(c *gin.Context) {
	type Param struct {
		ID int `form:"id" binding:"required"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	err := service.DeleUser(param.ID)
	if err != nil {
		response.Error(c, "删除redis失败", err)
		return
	}
	response.Success(c, true, "redis删除成功")
}
