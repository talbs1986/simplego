package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/talbs1986/simplego/configs/pkg/configs"
	"github.com/talbs1986/simplego/logger/pkg/logger"
)

const (
	DefaultServiceCloseTimeout = time.Second * 5
)

type AppOpt[T interface{}] func(*App[T])

type AppConfig struct {
	Name                string
	Version             string
	ServiceCloseTimeout time.Duration
}

type App[T interface{}] struct {
	Logger logger.ILogger
	Config configs.IConfigs[T]
	CTX    context.Context

	cancel            context.CancelFunc
	stopTimeout       time.Duration
	slog              logger.LogLine
	closeableServices []CloseableService
}

func NewApp[T interface{}](cfg *AppConfig, opts ...AppOpt[T]) *App[T] {
	if len(cfg.Name) < 1 {
		fmt.Fprintf(os.Stderr, "simplego app: failed to initialize app, service name is empty")
		os.Exit(1)
	}

	s := &App[T]{
		closeableServices: []CloseableService{},
	}
	for _, opt := range opts {
		opt(s)
	}
	if s.Logger == nil {
		s.Logger = DefaultLogger(logger.DefaultConfig)
	}
	if s.CTX == nil {
		s.CTX, s.cancel = context.WithCancel(context.Background())
	}
	if s.stopTimeout <= 0 {
		s.stopTimeout = DefaultServiceCloseTimeout
	}
	if s.Config == nil {
		s.Config = DefaultConfigurations[T](s.CTX)
	}
	s.slog = s.Logger.With(&logger.LogFields{"service": cfg.Name, "version": cfg.Version})
	s.slog.Info("simplego app: initialized successfully :) , GL HF")
	return s
}
