package router

import (
	test_user_api "go-gin-demo/api"

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
		redisGroup.POST("/setRedis", test_user_api.SetRedisApi)
		redisGroup.POST("/getRedis", test_user_api.GetRedisApi)
		redisGroup.POST("/deleRedis", test_user_api.DeleRedisApi)
	}

	return r
}
