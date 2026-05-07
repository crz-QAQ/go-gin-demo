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
	}

	return r
}
