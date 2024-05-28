package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/talbs1986/simplego/app/pkg/logger"
)

const (
	DefaultServiceCloseTimeout = time.Second * 5
)

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
