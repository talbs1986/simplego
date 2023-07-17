package app

import (
	"context"

	"github.com/talbs1986/simplego/configs/pkg/configs"
	"github.com/talbs1986/simplego/logger/pkg/logger"
)

func WithLoggerConfig[T interface{}](cfg *logger.Config) AppOpt[T] {
	return func(s *App[T]) {
		s.Logger = DefaultLogger(cfg)
	}
}

func WithLogger[T interface{}](l logger.ILogger) AppOpt[T] {
	return func(s *App[T]) {
		s.Logger = l
	}
}

func WithContext[T interface{}](ctx context.Context) AppOpt[T] {
	return func(s *App[T]) {
		s.CTX, s.cancel = context.WithCancel(ctx)
	}
}

func (s *App[T]) WithCloseableServices(services ...CloseableService) {
	for _, c := range services {
		service := c
		s.closeableServices = append(s.closeableServices, service)
	}
}

func WithConfigs[T interface{}](c configs.IConfigs[T]) AppOpt[T] {
	return func(s *App[T]) {
		s.Config = c
	}
}
