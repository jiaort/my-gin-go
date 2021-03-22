package router

import (
	v1 "my-gin-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitDetectionRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("detection")
	{
		UserRouter.POST("imgDetection", v1.ImgDetection)
		UserRouter.GET("ping", v1.Ping)
	}
}
