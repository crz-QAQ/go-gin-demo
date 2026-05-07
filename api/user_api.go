package user_api

import (
	"go-gin-demo/pkg/response"
	"go-gin-demo/service"

	"github.com/gin-gonic/gin"
)

// SoftDeleteUser 软删除用户
func SoftDeleteUser(c *gin.Context) {
	type Param struct {
		ID uint `form:"id" binding:"required"`
	}

	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}

	if err := service.SoftDeleteUser(param.ID); err != nil {
		response.Error(c, "删除失败", err)
		return
	}

	response.Success(c, true, "删除成功")
}

// DeleteUser 软删除用户
func DeleteUser(c *gin.Context) {
	type Param struct {
		ID uint `form:"id" binding:"required"`
	}

	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}

	if err := service.DeleteUser(param.ID); err != nil {
		response.Error(c, "删除失败", err)
		return
	}
	response.Success(c, true, "删除成功")
}
