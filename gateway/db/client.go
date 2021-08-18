package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type client struct {
	sync.RWMutex
	gormConnections map[string]*gorm.DB
	opts            Options
	started         bool
}

func (s *client) Options() Options {
	s.RLock()
	opts := s.opts
	s.RUnlock()
	return opts
}

func (s *client) Init(opts ...Option) error {
	s.configure(opts...)
	return nil
}

func (s *client) Stop() error {
	return nil
}

func (s *client) Start() error {
	s.RLock()
	if s.started {
		s.RUnlock()
		return nil
	}
	s.RUnlock()

	if err := s.openClient(); err != nil {
		return err
	}

	// mark the gin as started
	s.Lock()
	s.started = true
	s.Unlock()

	return nil
}

func (s *client) openClient() error {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	if len(s.opts.drivers) == 0 {
		return nil
	}

	// Register Database Connection to struct
	for dbName, v := range s.opts.drivers {
		Conn, err := gorm.Open(mysql.Open(s.DSN(v)), &gorm.Config{SkipDefaultTransaction: false})
		if err != nil {
			return err
		}

		sqlDB, err := Conn.DB()
		if err != nil {
			return err
		}

		// db connection pool params
		sqlDB.SetMaxOpenConns(v.MaxConnNum)  // 设置数据库连接池最大连接数
		sqlDB.SetMaxIdleConns(v.MaxIdleConn) // 最多空闲数量
		s.gormConnections[dbName] = Conn
	}

	return nil
}

func (s *client) DSN(cfg DB) string {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=PRC",
		cfg.Username, cfg.Password, cfg.Addr, cfg.Database,
	)
	if cfg.Timeout > 0 {
		dsn = fmt.Sprintf("%s&timeout=%ds", dsn, cfg.Timeout)
	}

	return dsn
}

func (s *client) configure(opts ...Option) {
	// Don't reprocess where there's no config
	if len(opts) == 0 && len(s.gormConnections) != 0 {
		return
	}

	for _, o := range opts {
		o(&s.opts)
	}
}

func newClient(opts ...Option) Client {
	options := newOptions(opts...)
	return &client{
		opts: options,
	}
}
