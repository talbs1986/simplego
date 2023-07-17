package configs

import "fmt"

type IConfigs[T interface{}] interface {
	Get() (*T, error)
}

type configImpl[T interface{}] struct {
	instance *T
}

type ConfigsOpt func(*configImpl)

func NewConfigs[T interface{}](opts ...ConfigsOpt) (IConfigs[T], error) {
	s := &configImpl[T]{}
	for _, opt := range opts {
		opt(s)
	}
	if s.instance == nil {
		if err := s.initalize(); err != nil {
			return nil, fmt.Errorf("simplego configs: failed to initialize config instance, due to: %w", err)
		}
	}
	return s, nil
}

func (s *configImpl[T]) initalize() error {
	return nil
}
func (s *configImpl[T]) Get() (*T, error) {
	return s.instance, nil
}
