package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type maxMsgSizeKey struct{}

type Options struct {
	Name           string
	Address        string
	Router         *gin.Engine
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

// Name Servername
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

// Address to bind to - host:port
func Address(a string) Option {
	return func(o *Options) {
		o.Address = a
	}
}

// Router set router
func Router(a *gin.Engine) Option {
	return func(o *Options) {
		o.Router = a
	}
}

// ReadTimeout set read timeout
func ReadTimeout(a time.Duration) Option {
	return func(o *Options) {
		o.ReadTimeout = a
	}
}

// WriteTimeout set write timeout
func WriteTimeout(a time.Duration) Option {
	return func(o *Options) {
		o.WriteTimeout = a
	}
}

func newOptions(opt ...Option) Options {
	opts := Options{
		Name:           DefaultName,
		Address:        DefaultAddress,
		Router:         DefaultRouter,
		ReadTimeout:    DefaultReadTimeout,
		WriteTimeout:   DefaultWriteTimeout,
		MaxHeaderBytes: DefaultMaxHeaderBytes,
	}

	for _, o := range opt {
		o(&opts)
	}

	return opts
}
