package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/stevenroose/gonfig"
	"gorm.io/driver/mysql"
	"time"
)

type Config struct {
	App struct {
		Name string `id:"name" default:""`
	} `id:"app" desc:"application config"`

	Logger struct {
		Level    string `id:"level" default:""`
		SavePath string `id:"save_path" default:""`
		SaveDay  int    `id:"save_day" default:"0"`
	} `id:"logger" desc:"log config"`

	HttpServer struct {
		RunMode      string        `id:"run_mode" default:"debug"`
		Port         int           `id:"port" default:"80"`
		ReadTimeout  time.Duration `id:"read_timeout" default:"30"`
		WriteTimeout time.Duration `id:"write_timeout" default:"30"`
		Domain       string        `id:"domain" default:""`
	} `id:"http_server" desc:"http config"`

	MySql struct {
		Default struct {
			Addr        string `id:"addr" default:""`
			Name        string `id:"name" default:""`
			MaxConnNum  int    `id:"max_conn_num" default:"30"`
			MaxIdleConn int    `id:"max_idle_conn" default:"30"`
			Username    string `id:"username" default:"root"`
			Password    string `id:"password" default:"123456"`
		}
	} `id:"mysql" desc:"mysql config"`

	Redis struct {
		Default struct {
			Addr        string        `id:"addr" default:""`
			Auth        string        `id:"auth" default:""`
			DB          int           `id:"db" default:"0"`
			PoolSize    int           `id:"pool_size" default:"30"`
			MinIdleConn int           `id:"min_idle_conn" default:"30"`
			PoolTimeout time.Duration `id:"poll_timeout" default:"30"`
		}
	} `id:"redis" desc:"redis config"`
}

var config *Config

// InitConfig Setup initialize the configuration instance
func InitConfig() *Config {
	config = &Config{}
	if err := gonfig.Load(config, gonfig.Conf{
		FileDefaultFilename: "./config.json",
		FlagDisable:         true,
		FileDecoder:         gonfig.DecoderJSON,
	}); err != nil {
		panic(err)
	}

	return config
}

// Get 获取配置
func Get() *Config {
	return config
}

func (c *Config) GetRedisDefaultOpts() *redis.Options {
	return &redis.Options{
		Addr:         c.Redis.Default.Addr,
		Password:     c.Redis.Default.Auth,
		DB:           c.Redis.Default.DB,
		MinIdleConns: c.Redis.Default.MinIdleConn,
		PoolSize:     c.Redis.Default.PoolSize,
		PoolTimeout:  c.Redis.Default.PoolTimeout * time.Second,
	}
}

func (c *Config) GetMysqlDefaultOpts() mysql.Config {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", c.MySql.Default.Username, c.MySql.Default.Password, c.MySql.Default.Addr, c.MySql.Default.Name)
	return mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         0,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  //  用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}
}
