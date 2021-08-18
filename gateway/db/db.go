package db

type Client interface {
	Options() Options
	Init(...Option) error

	Start() error
	Stop() error
}

type Option func(*Options)

// NewServer returns a new gin with options passed in
func NewServer(opt ...Option) Client {
	return newClient(opt...)
}
