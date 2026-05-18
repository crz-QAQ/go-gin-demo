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

// Login 用户登陆
func Login(c *gin.Context) {
	type Param struct {
		Phone    string `form:"phone" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	ip := base.GetClientIP(c)
	resp, err := service.LoginAccount(param.Phone, param.Password, ip)
	if err != nil {
		response.Error(c, "登陆失败", err)
		return
	}
	response.Success(c, resp, "登陆成功")
}

// PersonalMsg 获取用户基础数据
func PersonalMsg(c *gin.Context) {
	token, _ := c.Get("token")
	account, err := service.PersonalMsgService(token.(string))
	if err != nil {
		response.Error(c, "用户基础数据获取失败", err)
		return
	}
	response.Success(c, account, "用户基础数据")
}

// LogOut 登出
func LogOut(c *gin.Context) {
	token, _ := c.Get("token")
	err := service.LogOutService(token.(string))
	if err != nil {
		response.Error(c, "登出失败", err)
		return
	}
	response.Success(c, true, "登出成功")
}

// CreateDetail 创建用户详情
func CreateDetail(c *gin.Context) {
	type Param struct {
		IdNo    string `form:"id_no" binding:"required"`
		Sex     int8   `form:"sex"`
		Age     int8   `form:"age"`
		Hobby   string `form:"hobby"`
		Address string `form:"address"`
		Nation  string `form:"nation"`
	}

	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	token, _ := c.Get("token")
	detail, err := service.CreateDetailService(token.(string), param.IdNo, param.Sex, param.Age, param.Hobby, param.Address, param.Nation)
	if err != nil {
		response.Error(c, "创建详情失败", err)
		return
	}
	response.Success(c, detail, "用户详情创建成功")
}

// FindDetail 查找用户详情
func FindDetail(c *gin.Context) {
	token, _ := c.Get("token")
	detail, err := service.FindDetailService(token.(string))
	if err != nil {
		response.Error(c, "查询失败", err)
		return
	}
	response.Success(c, detail, "查询成功")
}

// UpdateDetail 更新用户详情
func UpdateDetail(c *gin.Context) {
	type Param struct {
		IdNo    string `form:"id_no"`
		Sex     int8   `form:"sex"`
		Age     int8   `form:"age"`
		Hobby   string `form:"hobby"`
		Address string `form:"address"`
		Nation  string `form:"nation"`
	}

	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}

	token, _ := c.Get("token")
	// 更新用户详情
	detail, err := service.UpdateDetailService(token.(string), param.IdNo, param.Sex, param.Age, param.Hobby, param.Address, param.Nation)
	if err != nil {
		response.Error(c, "更新用户详情失败", err)
		return
	}
	response.Success(c, detail, "更新用户详情成功")
}

// DeleteDetail 删除用户详情
func DeleteDetail(c *gin.Context) {
	token, _ := c.Get("token")
	result, err := service.DeleteDetailService(token.(string))
	if err != nil {
		response.Error(c, "删除用户详情失败", err)
		return
	}
	response.Success(c, result, "删除用户详情成功")
}

// DeleteAccount 注销用户
func DeleteAccount(c *gin.Context) {
	token, _ := c.Get("token")
	result, err := service.DeleteAccountService(token.(string))
	if err != nil {
		response.Error(c, "注销用户失败", err)
		return
	}
	response.Success(c, result, "注销用户成功")
}

// RestoreDeleteAccount 恢复软删除的数据
func RestoreDeleteAccount(c *gin.Context) {
	type Param struct {
		Phone string `form:"phone" binding:"required"`
	}

	var param Param
	if err := c.ShouldBind(&param); err != nil {
		response.Error(c, "参数错误", err)
		return
	}

	// 恢复软删除的数据
	account, err := service.RestoreDeleteAccountService(param.Phone)
	if err != nil {
		response.Error(c, "参数错误", err)
		return
	}
	response.Success(c, account, "恢复软删除数据")
}
