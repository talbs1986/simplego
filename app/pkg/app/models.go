package app

import (
	"context"
	"time"

	"github.com/talbs1986/simplego/app/pkg/logger"
)

// Env defines the environment type the app is running on
type Env string

const (
	// EnvLocal env local
	EnvLocal Env = "local"
	// EnvDev env ci
	EnvCI Env = "ci"
	// EnvDev env test
	EnvTest Env = "test"
	// EnvDev env dev
	EnvDev Env = "dev"
	// EnvStg env staging
	EnvStg Env = "stg"
	// EnvProd env production
	EnvProd Env = "prd"
)

// CloseableService defines an api for a closeable service
type CloseableService interface {
	// Close cleans and closes resources
	Close(ctx context.Context) error
}

// AppOpt defines the application options function
type AppOpt func(*App)

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
