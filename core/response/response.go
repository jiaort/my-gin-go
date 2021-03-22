package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func OkWithDataMessage(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(FAILED, map[string]interface{}{}, "失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(FAILED, map[string]interface{}{}, message, c)
}

func FailWithCode(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, "失败", c)
}

func FailWithCodeMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}
