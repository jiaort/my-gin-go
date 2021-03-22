package middleware

import (
	"my-gin-go/core/response"
	"my-gin-go/global"
	"net/http/httputil"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				// 是否开启控制台打印
				if stack {
					global.G_LOG.Error("[Revovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					global.G_LOG.Error("[Revovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				response.FailWithCodeMessage(response.UNKNOWN, "未知错误", c)
			}
		}()
		c.Next()
	}
}
