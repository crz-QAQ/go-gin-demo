package base

import "github.com/gin-gonic/gin"

func GetClientIP(c *gin.Context) string {
	ip := c.GetHeader("X-Forwarded-For")
	if ip != "" {
		return ip
	}
	ip = c.GetHeader("X-Real-Ip")
	if ip != "" {
		return ip
	}

	return c.ClientIP()
}
