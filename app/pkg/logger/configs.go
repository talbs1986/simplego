package logger

// LoggerConfig defines the basic app logger config supporting env vars
type LoggerConfig struct {
	LogLevel string `env:"LOG_LEVEL, default=debug"`
}
