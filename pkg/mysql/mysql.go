package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var ClientMap = make(map[ClientName]*gorm.DB)
var clientMapLock sync.Mutex

func InitClient(client ClientName, conf mysql.Config, maxConnNum int, maxIdleConn int) error {
	clientMapLock.Lock()
	defer clientMapLock.Unlock()

	Conn, err := gorm.Open(mysql.New(conf), &gorm.Config{SkipDefaultTransaction: false})
	if err != nil {
		return err
	}

	sqlDB, err := Conn.DB()
	if err != nil {
		return err
	}

	// 设置数据库连接池参数
	sqlDB.SetMaxOpenConns(maxConnNum)  // 设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(maxIdleConn) // 最多空闲数量
	ClientMap[client] = Conn

	return nil
}

func Client(name ClientName) *gorm.DB {
	return ClientMap[name]
}
