package configs

import (
	"context"
)

// ConfigParser defines an api for the configuration parser
type ConfigParser[T interface{}] interface {
	// Parse parses the configuration into an object
	Parse(context.Context) (*T, error)
	// Get gets the current configuration object
	Get(context.Context) (*T, error)
}
