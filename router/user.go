package router

import (
	v1 "my-gin-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("login", v1.Login)
		UserRouter.POST("register", v1.Register)
		UserRouter.GET("getUserList", v1.GetUserList)
	}
}
