package message_api

import (
	"go-gin-demo/dao"
	"go-gin-demo/pkg/response"
	"go-gin-demo/service"

	"github.com/gin-gonic/gin"
)

// CreateMessage 留言发布
func CreateMessage(c *gin.Context) {
	type Param struct {
		Content  string `form:"content" binding:"required"`
		Audience int8   `form:"audience" binding:"required"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	// 获取token
	token, _ := c.Get("token")
	message, err := service.CreateMessageService(token.(string), param.Content, param.Audience)
	if err != nil {
		response.Error(c, "留言发布失败", err)
		return
	}
	response.Success(c, message, "留言发布成功")
}

// AdminGetMessage 管理员获取留言列表
func AdminGetMessage(c *gin.Context) {
	type Param struct {
		Page     int   `form:"page" binding:"required,min=1"`
		PageSize int   `form:"pageSize" binding:"required,min=1,max=50"`
		Status   *int8 `form:"status"` // 审核状态，可选
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	list, total, err := service.GetMessageListByAdmin(param.Page, param.PageSize, param.Status)
	if err != nil {
		response.Error(c, "获取留言列表失败", err)
		return
	}

	response.Success(c, gin.H{
		"list":     list,
		"total":    total,
		"page":     param.Page,
		"pageSize": param.PageSize,
	}, "查询成功")
}

// PersonalGetMessageList 查询个人留言列表
func PersonalGetMessageList(c *gin.Context) {
	type Param struct {
		Page     int   `form:"page" binding:"required,min=1"`
		PageSize int   `form:"pageSize" binding:"required,min=1,max=50"`
		Status   *int8 `form:"status"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	// 获取token
	token, _ := c.Get("token")
	list, total, err := service.GetMessageListByPersonal(token.(string), param.Page, param.PageSize, param.Status)
	if err != nil {
		response.Error(c, "获取留言列表失败", err)
		return
	}
	response.Success(c, gin.H{
		"list":     list,
		"total":    total,
		"page":     param.Page,
		"pageSize": param.PageSize,
	}, "查询成功")
}

// MessageDetail 获取留言详情
func MessageDetail(c *gin.Context) {
	type Param struct {
		ID uint `form:"id" binding:"required"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	message, err := dao.GetMessageDetailById(param.ID)
	if err != nil {
		response.Error(c, "获取留言详情失败", err)
		return
	}
	response.Success(c, message, "获取留言详情成功")
}
