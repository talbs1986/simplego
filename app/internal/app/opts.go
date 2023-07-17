package app

import (
	"context"

	"github.com/talbs1986/simplego/logger/pkg/logger"
)

func WithLoggerConfig(cfg *logger.Config) AppOpt {
	return func(s *App) {
		s.Logger = defaultLogger(cfg)
	}
}

func WithLogger(l logger.ILogger) AppOpt {
	return func(s *App) {
		s.Logger = l
	}
}

func WithContext(ctx context.Context) AppOpt {
	return func(s *App) {
		s.ctx, s.cancel = context.WithCancel(ctx)
	}
}

func (s *App) WithCloseableServices(services ...CloseableService) {
	for _, c := range services {
		service := c
		s.closeableServices = append(s.closeableServices, service)
	}
}
