package db

// DB ...
type DB struct {
	Addr        string `id:"addr" default:""`
	Database    string `id:"db" default:""`
	MaxConnNum  int    `id:"max_conn_num" default:"30"`
	MaxIdleConn int    `id:"max_idle_conn" default:"30"`
	Username    string `id:"username" default:""`
	Password    string `id:"password" default:""`
	Timeout     int    `id:"timeout" default:""`
}

// Options ...
type Options struct {
	drivers map[string]DB
}

// RegisterDataBase ...
func RegisterDataBase(name string, db DB) func(o *Options) {
	return func(o *Options) {
		o.drivers[name] = db
	}
}

func newOptions(opt ...Option) Options {
	opts := Options{}

	for _, o := range opt {
		o(&opts)
	}

	return opts
}
