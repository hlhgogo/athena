package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/hlhgogo/athena/pkg/config"
	"sync"
	"time"
)

var clientMapLock sync.Mutex
var clientMap = make(map[string]*redis.Client)

// Load 连接redis
func Load() error {
	configMap, err := config.GetRedisOptions()
	if err != nil {
		return err
	}
	for connName, v := range configMap {
		conf := &redis.Options{
			Addr:         v.Addr,
			Password:     v.Auth,
			DB:           v.DB,
			MinIdleConns: v.MinIdleConn,
			PoolSize:     v.PoolSize,
			PoolTimeout:  v.PoolTimeout * time.Second,
		}
		if err := initClient(connName, conf); err != nil {
			return err
		}
	}
	return nil
}

// initClient 初始化客户端
func initClient(client string, rOpt *redis.Options) error {
	clientMapLock.Lock()
	defer clientMapLock.Unlock()

	conn, err := connect(rOpt)
	if err != nil {
		return err
	}

	clientMap[client] = conn

	return nil

}

// Client 获取客户端
func Client(name string) *redis.Client {
	return clientMap[name]
}

// 连接到redis
func connect(rOpt *redis.Options) (*redis.Client, error) {
	client := redis.NewClient(rOpt)
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
