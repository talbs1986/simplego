package configs

import (
	"context"
	"fmt"
)

type IConfigs[T interface{}] interface {
	Get(context.Context) (*T, error)
}

type ConfigParser[T interface{}] interface {
	Parse(context.Context, *T) (*T, error)
}

type ConfigsOpt[T interface{}] func(*configsImpl[T])

type configsImpl[T interface{}] struct {
	instance *T
	parsers  []ConfigParser[T]
}

func NewConfigs[T interface{}](ctx context.Context, opts ...ConfigsOpt[T]) (IConfigs[T], error) {
	s := &configsImpl[T]{}
	for _, opt := range opts {
		opt(s)
	}
	if s.instance == nil {
		if err := s.initalize(ctx); err != nil {
			return nil, fmt.Errorf("simplego configs: failed to initialize config instance, due to: %w", err)
		}
	}
	return s, nil
}

func (s *configsImpl[T]) initalize(ctx context.Context) error {
	if len(s.parsers) < 1 {
		return nil
	}
	cfg := new(T)
	var err error
	for _, p := range s.parsers {
		cfg, err = p.Parse(ctx, cfg)
		if err != nil {
			return err
		}
	}
	s.instance = cfg
	return nil
}
func (s *configsImpl[T]) Get(context.Context) (*T, error) {
	return s.instance, nil
}
