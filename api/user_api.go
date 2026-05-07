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

	if err := user_service.SoftDeleteUser(param.ID); err != nil {
		response.Error(c, "删除失败", err)
		return
	}

	response.Success(c, true, "删除成功")
}

// DeleteUser 硬删除用户
func DeleteUser(c *gin.Context) {
	type Param struct {
		ID uint `form:"id" binding:"required"`
	}

	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}

	if err := user_service.DeleteUser(param.ID); err != nil {
		response.Error(c, "删除失败", err)
		return
	}
	response.Success(c, true, "删除成功")
}

func CreateUser(c *gin.Context) {
	type Param struct {
		Name    string `form:"name" binding:"required"`
		Age     int64  `form:"age"`
		IdNo    string `form:"id_no" binding:"required"`
		Phone   int64  `form:"phone"`
		Sex     int    `form:"sex"`
		Hobby   string `form:"hobby"`
		Address string `form:"address"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}

	user, err := user_service.CreateUser(param.Name, param.Age, param.IdNo, param.Phone, param.Sex, param.Hobby, param.Address)
	if err != nil {
		response.Error(c, "用户创建失败", err)
		return
	}
	response.Success(c, user, "创建成功")

}
