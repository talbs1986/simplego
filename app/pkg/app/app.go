package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/talbs1986/simplego/app/pkg/logger"
)

const (
	// DefaultServiceCloseTimeout the default time for waiting for services to close shutdown
	DefaultServiceCloseTimeout = time.Second * 5
)

// NewApp creates a new application based on the config and options
func NewApp(cfg *AppConfig, opts ...AppOpt) *App {
	if len(cfg.Name) < 1 {
		fmt.Fprintf(os.Stderr, "simplego app: failed to initialize app, service name is empty")
		os.Exit(1)
	}

	s := &App{
		closeableServices: []CloseableService{},
		appServices:       map[string]interface{}{},
	}
	for _, opt := range opts {
		opt(s)
	}
	if s.Logger == nil {
		s.Logger = logger.NewFMTLogger(logger.DefaultConfig)
	}
	if s.CTX == nil {
		s.CTX, s.cancel = context.WithCancel(context.Background())
	}
	if s.stopTimeout <= 0 {
		s.stopTimeout = DefaultServiceCloseTimeout
	}
	s.slog = s.Logger.With(&logger.LogFields{"service": cfg.Name, "version": cfg.Version})
	s.slog.Info("simplego app: initialized successfully :) , GL HF")
	return s
}

// RegisterAppService registers a service to the App
func (s *App) RegisterAppService(key string, service interface{}) {
	s.appServices[key] = service
	if closeable, ok := service.(CloseableService); ok {
		s.closeableServices = append(s.closeableServices, closeable)
	}
}

// GetAppService gets a service by key from the registers App services or nil if missing
func (s *App) GetAppService(key string) interface{} {
	return s.appServices[key]
}
