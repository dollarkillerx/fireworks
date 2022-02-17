package server

import "github.com/dollarkillerx/fireworks/internal/conf"

type Backend struct {
	conf *conf.BackendConfig
}

func NewBackend(conf *conf.BackendConfig) *Backend {
	return &Backend{
		conf: conf,
	}
}

func (s *Backend) Run() error {
	return nil
}
