package gateway

import (
	"context"
	"github.com/hlhgogo/athena/gateway/db"
	"github.com/hlhgogo/athena/gateway/gin"
)

type Options struct {
	Server   gin.Server
	Database db.Client

	// Before and After funcs
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context

	Signal bool
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Server:   gin.NewServer(),
		Database: db.NewServer(),
		Context:  context.Background(),
		Signal:   true,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Context specifies a context for the service.
// Can be used to signal shutdown of the service.
// Can be used for extra option values.
func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

// Server option
func Server(s gin.Server) Option {
	return func(o *Options) {
		o.Server = s
	}
}

// Database option
func Database(c db.Client) Option {
	return func(o *Options) {
		o.Database = c
	}
}

// DbDriver ...
func DbDriver(name string, db db.DB) Option {
	return func(o *Options) {
		o.Database.Init()
	}
}

// Address sets the address of the gin
func Address(addr string) Option {
	return func(o *Options) {
		o.Server.Init(gin.Address(addr))
	}
}

// Name of the service
func Name(n string) Option {
	return func(o *Options) {
		o.Server.Init(gin.Name(n))
	}
}

// Before and Afters

func BeforeStart(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

func BeforeStop(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}

func AfterStart(fn func() error) Option {
	return func(o *Options) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

func AfterStop(fn func() error) Option {
	return func(o *Options) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}
