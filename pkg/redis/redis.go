package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
)

var clientMapLock sync.Mutex
var clientMap = make(map[ClientName]*redis.Client)

// InitClient 初始化客户端
func InitClient(client ClientName, rOpt *redis.Options) error {
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
func Client(name ClientName) *redis.Client {
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
