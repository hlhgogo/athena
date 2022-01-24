package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	DemoApi "github.com/hlhgogo/athena/internal/http/controller/demo"
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

	order := router.Group(v1prefix + "/")
	{
		order.GET("demo_list", DemoApi.OrderList)
	}
	return nil
}
