package app

import (
	"github.com/talbs1986/simplego/logger/pkg/logger"
)

type AppOpt func(*App)

type App struct {
	Logger logger.ILogger
}

func NewApp(opts ...AppOpt) *App {
	s := &App{
		// Logger: zerolog.New,
	}
	return s
}
