package middleware

import "github.com/gin-gonic/gin"

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(403, gin.H{"code": 403, "msg": "请先登录"})
			c.Abort()
			return
		}
		if role.(int8) != 1 {
			c.JSON(403, gin.H{"code": 403, "msg": "权限不足"})
			c.Abort()
			return
		}

		c.Next()
	}
}
