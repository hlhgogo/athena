package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	v1prefix         = "/api/v1"
	v1InternalPrefix = "/internal/api/v1"
)

// initRouter 初始化路由
func initRouter(router *gin.Engine) error {
	router.GET("/ready", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	router.GET("/healthy", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	return nil
}
