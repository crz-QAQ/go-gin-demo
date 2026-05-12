package test_user_api

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

	if err := service.DeleteUser(param.ID); err != nil {
		response.Error(c, "删除失败", err)
		return
	}
	response.Success(c, true, "删除成功")
}

// CreateUser 创建用户信息
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

	user, err := service.CreateUser(param.Name, param.Age, param.IdNo, param.Phone, param.Sex, param.Hobby, param.Address)
	if err != nil {
		response.Error(c, "用户创建失败", err)
		return
	}
	response.Success(c, user, "创建成功")
}

// FindUserEasyList 搜素所有用户基础信息
func FindUserEasyList(c *gin.Context) {
	user, err := service.FindUserEasyList()
	if err != nil {
		response.Error(c, "用户基础信息查询失败", err)
		return
	}
	response.Success(c, user, "搜素所有用户基础信息")
}

// FindUserEasyListReady 查询关联表全部记录
func FindUserEasyListReady(c *gin.Context) {
	user, err := service.FindUserEasyListReady()
	if err != nil {
		response.Error(c, "用户基础查询失败", err)
	}
	response.Success(c, user, "查询关联表全部记录")
}

// FindUserInfoList 查询关联表全部记录（普通join）
func FindUserInfoList(c *gin.Context) {
	user, err := service.FindUserInfoList()
	if err != nil {
		response.Error(c, "用户基础查询失败", err)
	}
	response.Success(c, user, "查询关联表全部记录（普通join）")
}

// FindWhere 查询关联表全部记录Where
func FindWhere(c *gin.Context) {
	type Param struct {
		Name string `form:"name" binding:"required"`
		IdNo string `form:"idNo" binding:"required"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	user, err := service.FindWhere(param.Name, param.IdNo)
	if err != nil {
		response.Error(c, "用户基础查询失败", err)
	}
	response.Success(c, user, "查询关联表全部记录Where")
}

// StructFind 查询关联表全部记录Struct
func StructFind(c *gin.Context) {
	type Param struct {
		Name string `form:"name" binding:"required"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	user, err := service.StructFind(param.Name)
	if err != nil {
		response.Error(c, "用户基础查询失败", err)
	}
	response.Success(c, user, "查询关联表全部记录Struct")
}

// MapFind 查询关联表全部记录Map
func MapFind(c *gin.Context) {
	type Param struct {
		Name string `form:"name" binding:"required"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	user, err := service.MapFind(param.Name)
	if err != nil {
		response.Error(c, "用户基础查询失败", err)
	}
	response.Success(c, user, "查询关联表全部记录Struct")
}

// UpdateSave 全量更新
func UpdateSave(c *gin.Context) {
	type Param struct {
		Name  string `form:"name" binding:"required"`
		IdNo  string `form:"id_no" binding:"required"`
		Phone int64  `form:"phone"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	detail, err := service.UpdateSave(param.Name, param.IdNo, param.Phone)
	if err != nil {
		response.Error(c, "用户基础查询失败", err)
	}
	response.Success(c, detail, "全量更新")
}

// UpdateApi 局部更新
func UpdateApi(c *gin.Context) {
	type Param struct {
		Name  string `form:"name" binding:"required"`
		IdNo  string `form:"id_no"`
		Phone int64  `form:"phone"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	detail, err := service.UpdateService(param.Name, param.IdNo, param.Phone)
	if err != nil {
		response.Error(c, "用户基础查询失败", err)
		return
	}
	response.Success(c, detail, "局部更新")
}

// UnscopedFindApi 查找被软删除的记录
func UnscopedFindApi(c *gin.Context) {
	user, err := service.UnscopedService()
	if err != nil {
		response.Error(c, "用户基础查询失败", err)
		return
	}
	response.Success(c, user, "查找被软删除的记录")
}
