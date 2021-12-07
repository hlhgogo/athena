package main

import (
	"github.com/hlhgogo/athena/pkg/gin-ext/sentry"
	"os"
	"os/signal"
	"syscall"

	"github.com/hlhgogo/athena/application/event"
	"github.com/hlhgogo/athena/internal/http"
	"github.com/hlhgogo/athena/pkg/config"
	"github.com/hlhgogo/athena/pkg/log"
	"github.com/hlhgogo/athena/pkg/sync/errgroup"
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

	//if err := redis.Load(); err != nil {
	//	panic("Redis Error:" + err.Error())
	//}
	//
	//if err := mysql.Load(); err != nil {
	//	panic(err)
	//}

	// 初始化sentry事件上报
	if config.Get().Sentry.Dsn != "" {
		sentry.Load()
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
