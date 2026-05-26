package router

import (
	"go-gin-demo/api/account_api"
	"go-gin-demo/api/message_api"
	"go-gin-demo/api/redis_api"
	"go-gin-demo/api/test_user_api"
	"go-gin-demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.Default()

	// 用户模块
	userGroup := r.Group("/user")
	{
		userGroup.POST("/soft-delete", test_user_api.SoftDeleteUser)
		userGroup.POST("/delete", test_user_api.DeleteUser)
		userGroup.POST("/create", test_user_api.CreateUser)
		userGroup.GET("/findAll", test_user_api.FindUserEasyList)
		userGroup.GET("/findDetail", test_user_api.FindUserEasyListReady)
		userGroup.GET("/findInfo", test_user_api.FindUserInfoList)
		userGroup.POST("/findWhere", test_user_api.FindWhere)
		userGroup.POST("/findStruct", test_user_api.StructFind)
		userGroup.POST("/findMap", test_user_api.MapFind)
		userGroup.POST("/Update/Save", test_user_api.UpdateSave)
		userGroup.POST("/Update", test_user_api.UpdateApi)
		userGroup.GET("/UnscopedFind", test_user_api.UnscopedFindApi)
	}

	redisGroup := r.Group("/redis")
	{
		redisGroup.POST("/setRedis", redis_api.SetRedisApi)
		authGroup := redisGroup.Use(middleware.AuthLogin())
		{
			authGroup.POST("/getRedis", redis_api.GetRedisApi)
			authGroup.POST("/deleRedis", redis_api.DeleRedisApi)
		}

	}

	accountGroup := r.Group("/account")
	{
		accountGroup.POST("/register", account_api.Register)
		accountGroup.POST("/login", account_api.Login)
		accountGroup.POST("/restore", account_api.RestoreDeleteAccount)
		accountGroup.POST("/forgetPasswrod", account_api.ForgetPassword)

		authGroup := accountGroup.Use(middleware.AuthLogin())
		{
			authGroup.GET("/logout", account_api.LogOut)
			authGroup.GET("/personMsg", account_api.PersonalMsg)
			authGroup.POST("/create/detail", account_api.CreateDetail)
			authGroup.GET("/search/detail", account_api.FindDetail)
			authGroup.POST("/update/detail", account_api.UpdateDetail)
			authGroup.DELETE("/delete/detail", account_api.DeleteDetail)
			authGroup.DELETE("/delete", account_api.DeleteAccount)
			authGroup.POST("/changePassword", account_api.UpdatePasswordToken)
			authGroup.POST("/changeNickname", account_api.UpdateNickname)
		}

	}

	messageGroup := r.Group("/message")
	{
		authGroup := messageGroup.Use(middleware.AuthLogin())
		{
			authGroup.POST("/create/message", message_api.CreateMessage)
			authGroup.POST("/list/personal_message", message_api.PersonalGetMessageList)
			authGroup.POST("/detail/message", message_api.MessageDetail)
		}

		adminGroup := messageGroup.Use(middleware.AuthLogin(), middleware.AdminAuth())
		{
			adminGroup.POST("/admin/get_message", message_api.AdminGetMessage)
		}
	}

	return r
}
