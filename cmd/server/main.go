package main

import (
	"context"
	"github.com/hlhgogo/athena/internal/http"
	"github.com/hlhgogo/athena/pkg/config"
	"github.com/hlhgogo/athena/pkg/log"
	"github.com/hlhgogo/athena/pkg/mysql"
	"github.com/hlhgogo/athena/pkg/redis"
	"github.com/hlhgogo/athena/pkg/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
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
	conf := config.InitConfig()
	log.Setup()

	if err := redis.InitClient(redis.Default, conf.GetRedisDefaultOpts()); err != nil {
		panic("Redis Error:" + err.Error())
	}

	if err := mysql.InitClient(mysql.Default, conf.GetMysqlDefaultOpts(), conf.MySql.Default.MaxConnNum, conf.MySql.Default.MaxIdleConn); err != nil {
		panic("Mysql Error:" + err.Error())
	}
	g.Go(func(context.Context) (err error) {
		if err := http.Serve(); err != nil {
			return err
		}
		return nil
	})

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
