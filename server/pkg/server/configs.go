package server

import "time"

// ServerConfig defines the basic server config supporting env vars
type ServerConfig struct {
	Addr              string        `env:"SERVER_ADDR, default=:8081"`
	RequestTimeout    time.Duration `env:"SERVER_REQUEST_TIMEOUT, default=5m"`
	ReadTimeout       time.Duration `env:"SERVER_READ_TIMEOUT, default=1m"`
	ReadHeaderTimeout time.Duration `env:"SERVER_READ_HEADER_TIMEOUT, default=1m"`
	WriteTimeout      time.Duration `env:"SERVER_WRITE_TIMEOUT, default=1m"`
	IdleTimeout       time.Duration `env:"SERVER_IDLE_TIMEOUT, default=1m"`
}
