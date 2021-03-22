package main

import (
	"my-gin-go/core"
	"my-gin-go/global"
	"my-gin-go/initialize"
)

// @title 接口文档
// @version 0.0.1
// @description 图像识别服务
// @termsOfService https://www.topgoer.com

// @contact.name www.topgoer.com
// @contact.url https://www.topgoer.com
// @contact.email me@razeen.me

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 10.0.24.5:8888
// @BasePath /api/v1
func main() {
	// 初始化gorm
	initialize.Gorm()
	// 初始化redis服务
	initialize.Redis()
	db, _ := global.G_DB.DB()
	defer db.Close()
	core.RunServer()
}

// opencv-python==3.4.2.17
// keras==2.3.1
// imageai==2.1.5
// tensorflow==1.13.2
