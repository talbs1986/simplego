package app

import (
	"context"

	"github.com/talbs1986/simplego/app/pkg/logger"
)

func WithContext(ctx context.Context) AppOpt {
	return func(s *App) {
		s.CTX, s.cancel = context.WithCancel(ctx)
	}
}

func (s *App) WithCloseableServices(services ...CloseableService) {
	for _, c := range services {
		service := c
		s.closeableServices = append(s.closeableServices, service)
	}
}

func WithDefaultLoggerConfig(cfg *logger.Config) AppOpt {
	return func(s *App) {
		s.Logger = logger.NewFMTLogger(cfg)
	}
}

func WithLogger(l logger.ILogger) AppOpt {
	return func(s *App) {
		s.Logger = l
	}
}
