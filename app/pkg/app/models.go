package app

import (
	"context"
	"time"

	"github.com/talbs1986/simplego/app/pkg/logger"
)

// CloseableService defines an api for a closeable service
type CloseableService interface {
	// Close cleans and closes resources
	Close(ctx context.Context) error
}

// AppOpt defines the application options function
type AppOpt func(*App)

// AppConfig defines the configurations object for the App
type AppConfig struct {
	Name                string        `env:"APP_NAME, default=serviceA"`
	Version             string        `env:"APP_VERSION, default=1.0.0"`
	ServiceCloseTimeout time.Duration `env:"APP_SERVICE_CLOSE_TIMEOUT, default=30s"`
}

// App defines the application object
type App struct {
	Logger logger.ILogger
	CTX    context.Context

	cancel            context.CancelFunc
	stopTimeout       time.Duration
	slog              logger.LogLine
	closeableServices []CloseableService

	appServices map[string]interface{}
}
