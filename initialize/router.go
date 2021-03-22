package initialize

import (
	_ "my-gin-go/docs"
	"my-gin-go/global"
	"my-gin-go/middleware"
	"my-gin-go/router"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 静态文件路径注册
	Router.StaticFS(global.G_CONFIG.Local.Path, http.Dir(global.G_CONFIG.Local.Path))
	Router.StaticFS(global.G_CONFIG.Local.Result, http.Dir(global.G_CONFIG.Local.Result))
	// 中间件
	// middleware.Cors() 跨域 middleware.GinRecovery(true) 全局异常捕获 middleware.LoadTls() 支持https
	Router.Use(middleware.Cors(), middleware.GinRecovery(true))

	// 文档界面访问URL swag init
	// http://.../swagger/index.html
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ApiGroup := Router.Group("/api/v1")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	router.InitDetectionRouter(ApiGroup)
	return Router
}
