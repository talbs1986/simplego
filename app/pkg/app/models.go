package app

import (
	"context"
	"time"

	"github.com/talbs1986/simplego/app/pkg/logger"
)

type CloseableService interface {
	Close(ctx context.Context) error
}

type AppOpt func(*App)

type AppConfig struct {
	Name                string
	Version             string
	ServiceCloseTimeout time.Duration
}

type App struct {
	Logger logger.ILogger
	CTX    context.Context

	cancel            context.CancelFunc
	stopTimeout       time.Duration
	slog              logger.LogLine
	closeableServices []CloseableService
}
