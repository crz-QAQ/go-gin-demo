package middleware

import (
	"bytes"
	"encoding/json"
	"go-gin-demo/model"
	mq "go-gin-demo/pkg/rabbitmq"
	"go-gin-demo/pkg/redis"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func OperateLogMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跳过无需记录的接口
		skipList := []string{"/log/list"}
		path := c.FullPath()
		for _, skip := range skipList {
			if strings.Contains(path, skip) {
				c.Next()
				return
			}
		}

		// 基础请求信息
		reqUrl := path
		reqMethod := c.Request.Method
		clientIp := c.ClientIP()

		// 默认游客
		var userId int64 = 0
		var userName string = "游客"

		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			redisKey := "login:token:" + authHeader
			accountStr, err := redis.Get(redisKey)
			if err == nil {
				var userInfo map[string]interface{}
				_ = json.Unmarshal([]byte(accountStr), &userInfo)
				// 你SetToken里存的是 username 不是 name！
				if idVal, ok := userInfo["id"].(float64); ok {
					userId = int64(idVal)
				}
				if nameVal, ok := userInfo["username"].(string); ok {
					userName = nameVal
				}
			}
		}

		// 获取请求参数
		var reqParam string
		if strings.Contains(reqMethod, "POST") {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			reqParam = string(bodyBytes)
			// 放回body，不影响后续
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		} else {
			reqParam = c.Request.URL.RawQuery
		}
		c.Next()

		// 接口状态
		status := 1
		if c.Writer.Status() >= http.StatusBadRequest {
			status = 0
		}

		operateDesc := getOpDesc(path)

		logMsg := model.OperateLogMsg{
			UserId:      userId,
			UserName:    userName,
			Url:         reqUrl,
			Method:      reqMethod,
			ClientIp:    clientIp,
			Param:       reqParam,
			OperateDesc: operateDesc,
			Status:      status,
			OperateTime: time.Now().Format("2006-01-02 15:04:05.000"),
		}

		go func(m model.OperateLogMsg) {
			// 放大延迟，直观看出异步差距
			time.Sleep(2 * time.Second)
			msgBytes, err := json.Marshal(m)
			if err == nil {
				_ = mq.SendMsg("operate_log_queue", msgBytes)
			}
		}(logMsg)

	}
}

func getOpDesc(url string) string {
	switch {
	case strings.Contains(url, "/account/login"):
		return "用户登录系统"
	case strings.Contains(url, "/account/register"):
		return "用户注册账号"
	case strings.Contains(url, "/sign/sign"):
		return "用户每日签到"
	case strings.Contains(url, "/account/logout"):
		return "用户退出登录"
	default:
		return "访问系统功能接口"
	}
}
