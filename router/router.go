package router

import (
	user_api "go-gin-demo/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.Default()

	// 用户模块
	userGroup := r.Group("/user")
	{
		userGroup.POST("/soft-delete", user_api.SoftDeleteUser)
		userGroup.POST("/delete", user_api.DeleteUser)
		userGroup.POST("/create", user_api.CreateUser)
		userGroup.GET("/findAll", user_api.FindUserEasyList)
	}

	return r
}
