package app

import (
	"context"

	"github.com/talbs1986/simplego/app/pkg/logger"
)

// WithContext application option to replace the current application context
func WithContext(ctx context.Context) AppOpt {
	return func(s *App) {
		s.CTX, s.cancel = context.WithCancel(ctx)
	}
}

// WithCloseableServices application option to add services to global app close
func (s *App) WithCloseableServices(services ...CloseableService) {
	for _, c := range services {
		service := c
		s.closeableServices = append(s.closeableServices, service)
	}
}

// WithDefaultLoggerConfig application option to set the app logger as the default FMT logger
func WithDefaultLoggerConfig(cfg *logger.Config) AppOpt {
	return func(s *App) {
		s.Logger = logger.NewFMTLogger(cfg)
	}
}

// WithLogger application option to set the app logger
func WithLogger(l logger.ILogger) AppOpt {
	return func(s *App) {
		s.Logger = l
	}
}
