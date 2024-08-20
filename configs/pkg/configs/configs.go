package configs

// EnvConfig defines the basic app environment config
type EnvConfig struct {
	Env Env `env:"ENV, default=local"`
}

// LoggerConfig defines the basic app logger config
type LoggerConfig struct {
	LogLevel string `env:"LOG_LEVEL, default=debug"`
}
