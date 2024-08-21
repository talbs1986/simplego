package configs

// EnvConfig defines the basic app environment config supporting env vars
type EnvConfig struct {
	Env Env `env:"ENV, default=local"`
}

// LoggerConfig defines the basic app logger config supporting env vars
type LoggerConfig struct {
	LogLevel string `env:"LOG_LEVEL, default=debug"`
}
