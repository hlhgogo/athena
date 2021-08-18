package gin

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Server interface {
	Options() Options
	Init(...Option) error
	Start() error
	Stop() error
	String() string
}

var (
	DefaultAddress                    = ":5000"
	DefaultName                       = "Athena"
	DefaultRouter         *gin.Engine = nil
	DefaultReadTimeout                = 60 * time.Second
	DefaultWriteTimeout               = 60 * time.Second
	DefaultMaxHeaderBytes             = 1 << 20
)

type Option func(*Options)

// NewServer returns a new gin with options passed in
func NewServer(opt ...Option) Server {
	return newGinServer(opt...)
}
