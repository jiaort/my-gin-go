// +build windows

package core

import (
	"my-gin-go/global"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    time.Duration(global.G_CONFIG.System.Timeout) * time.Second,
		WriteTimeout:   time.Duration(global.G_CONFIG.System.Timeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
