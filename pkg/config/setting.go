package config

import (
	"github.com/stevenroose/gonfig"
	"time"
)

var config *Config

type Config struct {
	App struct {
		Name      string `id:"name" default:"Athena"`
		ShowTrace bool   `id:"show_trace" default:"true"`
	} `id:"app" desc:"application config"`

	Sentry struct {
		Dsn         string `id:"dsn" default:""`
		Environment string `id:"environment" default:"development"`
		Release     string `id:"release" default:"v1.0.0"`
	} `id:"sentry" desc:"sentry config"`

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

	MySql map[string]interface{} `id:"mysql" desc:"mysql config"`
	Redis map[string]interface{} `id:"redis" desc:"redis config"`
}

type MySqlConfig struct {
	Addr        string `id:"addr" default:""`
	Name        string `id:"name" default:""`
	MaxConnNum  int    `id:"max_conn_num" default:"30"`
	MaxIdleConn int    `id:"max_idle_conn" default:"30"`
	Username    string `id:"username" default:"root"`
	Password    string `id:"password" default:"123456"`
}

type RedisConfig struct {
	Addr        string        `id:"addr" default:""`
	Auth        string        `id:"auth" default:""`
	DB          int           `id:"db" default:"0"`
	PoolSize    int           `id:"pool_size" default:"30"`
	MinIdleConn int           `id:"min_idle_conn" default:"30"`
	PoolTimeout time.Duration `id:"poll_timeout" default:"30"`
}

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

// GetRedisOptions 获取Redis连接配置
func GetRedisOptions() (map[string]RedisConfig, error) {
	options := map[string]RedisConfig{}
	for connName, value := range config.Redis {
		m, ok := value.(map[string]interface{})
		if !ok {
			continue
		}

		v := RedisConfig{}
		if err := gonfig.LoadMap(&v, m, gonfig.Conf{}); err != nil {
			return options, err
		}

		options[connName] = v
	}

	return options, nil
}

// GetMysqlOptions 获取Mysql连接配置
func GetMysqlOptions() (map[string]MySqlConfig, error) {
	options := map[string]MySqlConfig{}
	for connName, value := range config.MySql {
		m, ok := value.(map[string]interface{})
		if !ok {
			continue
		}

		v := MySqlConfig{}
		if err := gonfig.LoadMap(&v, m, gonfig.Conf{}); err != nil {
			return options, err
		}

		options[connName] = v
	}

	return options, nil
}

// Get 获取配置
func Get() *Config {
	return config
}
