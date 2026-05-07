package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一返回结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 成功返回
func Success(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  msg,
		Data: data,
	})
}

// Error 失败返回
func Error(c *gin.Context, msg string, err error) {
	errMsg := msg
	if err != nil {
		errMsg = msg + "：" + err.Error()
	}

	c.JSON(http.StatusOK, Response{
		Code: 500,
		Msg:  errMsg,
		Data: nil,
	})
}
