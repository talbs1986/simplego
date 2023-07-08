package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/talbs1986/simplego/logger/pkg/logger"
	zerolog "github.com/talbs1986/simplego/zerolog-logger/pkg/logger"
)

const (
	DefaultServiceCloseTimeout = time.Second * 5
)

type AppOpt func(*App)

type AppConfig struct {
	Name                string
	Version             string
	ServiceCloseTimeout time.Duration
}

type App struct {
	Logger logger.ILogger

	ctx               context.Context
	cancel            context.CancelFunc
	stopTimeout       time.Duration
	slog              logger.LogLine
	closeableServices []CloseableService
}

func NewApp(cfg *AppConfig, opts ...AppOpt) *App {
	if len(cfg.Name) < 1 {
		fmt.Fprintf(os.Stderr, "simplego app: failed to initialize app, service name is empty")
		os.Exit(1)
	}

	s := &App{
		closeableServices: []CloseableService{},
	}
	for _, opt := range opts {
		opt(s)
	}
	if s.Logger == nil {
		l, err := zerolog.NewSimpleZerolog(nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "simplego app: failed to initialize default logger, due to: %s", err.Error())
			os.Exit(1)
		}
		s.Logger = l
	}
	if s.ctx == nil {
		s.ctx, s.cancel = context.WithCancel(context.Background())
	}
	if s.stopTimeout <= 0 {
		s.stopTimeout = DefaultServiceCloseTimeout
	}
	s.slog = s.Logger.With(&logger.LogFields{"service": cfg.Name, "version": cfg.Version})
	s.slog.Info("simplego app: initialized successfully :) , GL HF")
	return s
}

func defaultLogger(cfg *logger.Config) logger.ILogger {
	l, err := zerolog.NewSimpleZerolog(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "simplego app: failed to initialize default logger, due to: %s", err.Error())
		os.Exit(1)
	}
	return l
}

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
