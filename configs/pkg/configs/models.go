package configs

import (
	"context"
)

// Env defines the environment type the app is running on
type Env string

const (
	// EnvLocal env local
	EnvLocal Env = "local"
	// EnvDev env dev
	EnvDev Env = "dev"
	// EnvStg env staging
	EnvStg Env = "stg"
	// EnvProd env production
	EnvProd Env = "prd"
)

// ConfigParser defines an api for the configuration parser
type ConfigParser[T interface{}] interface {
	// Parse parses the configuration into an object
	Parse(context.Context) (*T, error)
	// Get gets the current configuration object
	Get(context.Context) (*T, error)
}
