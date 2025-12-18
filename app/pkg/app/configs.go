package app

import "time"

// AppConfig defines the configurations object for the App
type AppConfig struct {
	Name                string        `env:"APP_NAME, default=serviceA"`
	Version             string        `env:"APP_VERSION, default=1.0.0"`
	ServiceCloseTimeout time.Duration `env:"APP_SERVICE_CLOSE_TIMEOUT, default=30s"`
}

// EnvConfig defines the basic app environment config supporting env vars
type EnvConfig struct {
	Env Env `env:"ENV, default=local"`
}
