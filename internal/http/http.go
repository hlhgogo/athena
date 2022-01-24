package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hlhgogo/config"
	"github.com/hlhgogo/gin-ext/log"
	"github.com/hlhgogo/gin-ext/middlewares"
	"net/http"
	"time"
)

var httpServer *http.Server

func router() http.Handler {
	r := gin.New()

	r.Use(middlewares.Trace())
	r.Use(middlewares.Sentry())
	r.Use(middlewares.Cors())
	r.Use(middlewares.Recovery())
	r.Use(middlewares.LoggerWithFormatter())
	r.NoRoute(middlewares.PageNotFound)

	err := initRouter(r)
	if err != nil {
		log.Error("init router", err)
		return r
	}

	return r
}

func Serve() error {
	// Do Stuff Here
	gin.SetMode(config.Get().HttpServer.RunMode)
	//gin.DefaultWriter = io.MultiWriter(log.GetGinLogIoWriter(), os.Stdout)

	readTimeout := config.Get().HttpServer.ReadTimeout
	writeTimeout := config.Get().HttpServer.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.Get().HttpServer.Port)
	maxHeaderBytes := 1 << 20

	httpServer = &http.Server{
		Addr:           endPoint,
		Handler:        router(),
		ReadTimeout:    readTimeout * time.Second,
		WriteTimeout:   writeTimeout * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	fmt.Printf("%s [INFO] [%14s] launch http listen [:%d] Run %s Model\n", config.Get().App.Name,
		time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		config.Get().HttpServer.Port, config.Get().HttpServer.RunMode)

	if err := httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
