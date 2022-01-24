package main

import (
	"github.com/hlhgogo/athena/pkg/mysql"
	"github.com/hlhgogo/athena/pkg/redis"
	"github.com/hlhgogo/gin-ext/sentry"
	"os"
	"os/signal"
	"syscall"

	"github.com/hlhgogo/athena/application/event"
	"github.com/hlhgogo/athena/internal/http"
	"github.com/hlhgogo/config"
	"github.com/hlhgogo/gin-ext/log"
	"github.com/hlhgogo/gin-ext/sync/errgroup"
)

var (
	g errgroup.Group
	//tracingCloser io.Closer
)

func shutdown() error {
	if err := http.Shutdown(); err != nil {
		return err
	}
	return nil
}

func main() {
	_ = config.InitConfig()
	log.Setup()

	// 初始化redis
	if err := redis.Load(); err != nil {
		panic(err)
	}

	// 初始化mysql
	if err := mysql.Load(); err != nil {
		panic(err)
	}

	// 初始化sentry事件上报
	if err := sentry.Load(); err != nil {
		panic(err)
	}

	// 事件初始化
	event.Init()

	go func() {
		if err := http.Serve(); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			shutdown()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}
