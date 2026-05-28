package sign_api

import (
	"go-gin-demo/pkg/response"
	"go-gin-demo/service"

	"github.com/gin-gonic/gin"
)

// Sign 用户签到
func Sign(c *gin.Context) {
	token, _ := c.Get("token")
	sign, err := service.CreateSign(token.(string))
	if err != nil {
		response.Error(c, "签到失败", err)
		return
	}
	response.Success(c, sign, "签到成功")
}

// UserSignList 用户签到列表
func UserSignList(c *gin.Context) {
	type Param struct {
		Page     int `form:"page" binding:"required,min=1"`
		PageSize int `form:"pageSize" binding:"required,min=1,max=50"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	token, _ := c.Get("token")
	list, total, err := service.ListUser(token.(string), param.Page, param.PageSize)
	if err != nil {
		response.Error(c, "查询签到列表失败", err)
		return
	}
	response.Success(c, gin.H{
		"list":     list,
		"total":    total,
		"page":     param.Page,
		"pageSize": param.PageSize,
	}, "查询签到列表成功")
}

// AdminSignList 管理员查询签到列表
func AdminSignList(c *gin.Context) {
	type Param struct {
		Page     int    `form:"page" binding:"required,min=1"`
		PageSize int    `form:"pageSize" binding:"required,min=1,max=50"`
		Phone    string `form:"phone" binding:"max=11"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	list, total, err := service.ListAdmin(param.Phone, param.Page, param.PageSize)
	if err != nil {
		response.Error(c, "查询签到列表失败", err)
		return
	}
	response.Success(c, gin.H{
		"list":     list,
		"total":    total,
		"page":     param.Page,
		"pageSize": param.PageSize,
	}, "查询签到列表成功")
}
