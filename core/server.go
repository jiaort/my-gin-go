package core

import (
	"my-gin-go/global"
	"my-gin-go/initialize"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.G_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.G_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	欢迎使用 图像识别平台
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, address)
	global.G_LOG.Error(s.ListenAndServe().Error())
}
