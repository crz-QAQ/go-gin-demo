package middleware

import (
	"go-gin-demo/pkg/redis"

	"github.com/gin-gonic/gin"
)

func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"code": 401, "msg": "未登录，请先登录"})
			c.Abort()
			return
		}

		// 解析token、判断是否过期
		redisKey := "login:token:" + authHeader

		account, err := redis.Get(redisKey)
		if err != nil {
			c.JSON(401, gin.H{"code": 401, "msg": "登录已过期，请重新登录"})
			c.Abort()
			return
		}
		c.Set("token", authHeader)
		c.Set("account", account)

		// 放行
		c.Next()
	}
}
