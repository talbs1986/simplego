package app

import (
	"context"
	"time"

	"github.com/talbs1986/simplego/app/pkg/logger"
)

// Env defines the enviornment type the app is running on
type Env string

const (
	// EnvLocal env local
	EnvLocal = "local"
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

// AppConfig defines the configurations object for the App
type AppConfig struct {
	Name                string
	Version             string
	ServiceCloseTimeout time.Duration
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

// RegisterAppService registers a service to the App
func (s *App) RegisterAppService(key string, service interface{}) {
	s.appServices[key] = service
}

// GetAppService gets a service by key from the registers App services
func (s *App) GetAppService(key string) interface{} {
	return s.appServices[key]
}
