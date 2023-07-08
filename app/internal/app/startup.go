package app

import (
	"fmt"
	"os"

	"github.com/talbs1986/simplego/logger/pkg/logger"
	zerolog "github.com/talbs1986/simplego/zerolog-logger/pkg/logger"
)

type AppOpt func(*App)

type App struct {
	Logger logger.ILogger
}

func NewApp(opts ...AppOpt) *App {
	s := &App{}
	for _, opt := range opts {
		opt(s)
	}
	if s.Logger == nil {
		l, err := zerolog.NewSimpleZerolog(nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "simplego app: failed to initialize default logger, due to: %s", err.Error())
			os.Exit(1)
		}
		s.Logger = l
	}
	return s
}

func defaultLogger(cfg *logger.Config) logger.ILogger {
	l, err := zerolog.NewSimpleZerolog(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "simplego app: failed to initialize default logger, due to: %s", err.Error())
		os.Exit(1)
	}
	return l
}

func WithLoggerConfig(cfg *logger.Config) AppOpt {
	return func(s *App) {
		s.Logger = defaultLogger(cfg)
	}
}

func WithLogger(l logger.ILogger) AppOpt {
	return func(s *App) {
		s.Logger = l
	}
}
