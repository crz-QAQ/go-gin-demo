package log_api

import (
	"go-gin-demo/pkg/response"
	"go-gin-demo/service"

	"github.com/gin-gonic/gin"
)

func OperateList(c *gin.Context) {
	type Param struct {
		Phone    string `form:"phone"`
		Status   *int8  `form:"status"`
		Page     int    `form:"page" binding:"required,min=1"`
		PageSize int    `form:"pageSize" binding:"required,min=1,max=50"`
	}

	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}

	list, total, err := service.OperateLogList(param.PageSize, param.Page, param.Phone, param.Status)
	if err != nil {
		response.Error(c, "获取操作日志列表失败", err)
		return
	}
	response.Success(c, gin.H{
		"list":     list,
		"total":    total,
		"page":     param.Page,
		"pageSize": param.PageSize,
	}, "获取操作日志列表成功")
}
