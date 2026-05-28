package middleware

import (
	"encoding/json"
	"go-gin-demo/pkg/redis"
	"time"

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

		// 用户访问接口 → 自动续期10分钟
		_ = redis.Expire(redisKey, 10*time.Minute)

		var userInfo map[string]interface{}
		err = json.Unmarshal([]byte(account), &userInfo)
		if err != nil {
			c.JSON(401, gin.H{"code": 401, "msg": "用户信息解析失败"})
			c.Abort()
			return
		}

		// 4. 从map取出 role，并转成 int8
		role := int8(userInfo["role"].(float64))

		c.Set("token", authHeader)
		c.Set("role", role)

		// 放行
		c.Next()
	}
}
