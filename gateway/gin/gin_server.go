package gin

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type ginServer struct {
	sync.RWMutex

	// http gin
	srv *http.Server

	exit chan chan error
	opts Options

	// marks the serve as started
	started bool

	// graceful exit
	wg *sync.WaitGroup
}

func (s *ginServer) Options() Options {
	s.RLock()
	opts := s.opts
	s.RUnlock()
	return opts
}

func (s *ginServer) Start() error {
	s.RLock()
	if s.started {
		s.RUnlock()
		return nil
	}
	s.RUnlock()

	// micro: go ts.Accept(s.accept)
	go func() {
		fmt.Println("start gin server")
		if err := s.srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	// mark the gin as started
	s.Lock()
	s.started = true
	s.Unlock()

	return nil
}

func (s *ginServer) Stop() error {
	s.RLock()
	if !s.started {
		s.RUnlock()
		return nil
	}
	s.RUnlock()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}

	s.Lock()
	s.started = false
	s.Unlock()

	return nil
}

func (s *ginServer) Init(opts ...Option) error {
	s.configure(opts...)
	return nil
}

func (s *ginServer) String() string {
	return "grpc"
}

func (s *ginServer) configure(opts ...Option) {
	// Don't reprocess where there's no config
	if len(opts) == 0 && s.srv != nil {
		return
	}

	for _, o := range opts {
		o(&s.opts)
	}

	config := s.Options()

	s.srv = &http.Server{
		Addr:           config.Address,
		Handler:        config.Router,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: config.MaxHeaderBytes,
	}
}

func (s *ginServer) GetGinServer() *http.Server {
	return s.srv
}

func newGinServer(opts ...Option) Server {
	options := newOptions(opts...)
	wg := &sync.WaitGroup{}
	return &ginServer{
		opts: options,
		exit: make(chan chan error),
		wg:   wg,
	}
}
