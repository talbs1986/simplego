package configs

import (
	"context"

	"github.com/talbs1986/simplego/app/pkg/app"
)

// ConfigParser defines an api for the configuration parser
type ConfigParser[T interface{}] interface {
	// Parse parses the configuration into an object
	Parse(context.Context) (*T, error)
	// Get gets the current configuration object
	Get(context.Context) (*T, error)
}

type EnvConfig struct {
	Env app.Env `env:"ENV, default=local"`
}

type LoggerConfig struct {
	LogLevel string `env:"LOG_LEVEL, default=debug"`
}
