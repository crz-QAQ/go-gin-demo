package account_api

import (
	"go-gin-demo/pkg/base"
	"go-gin-demo/pkg/response"
	"go-gin-demo/service"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	type Param struct {
		Name     string `form:"name" binding:"required"`
		Phone    string `form:"phone" binding:"required"`
		Password string `form:"password" binding:"required"`
		Confirm  string `form:"confirm" binding:"required"`
		Nickname string `form:"nickname"`
		Role     int8   `form:"role" binding:"required"`
	}

	// 验证参数
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}

	ip := base.GetClientIP(c)

	account, err := service.RegisterAccount(param.Name, param.Phone, param.Password, param.Nickname, param.Role, param.Confirm, ip)
	if err != nil {
		response.Error(c, "用户注册失败", err)
		return
	}
	response.Success(c, account, "注册成功，token有效时间10分钟")
}
