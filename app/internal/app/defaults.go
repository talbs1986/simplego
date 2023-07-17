package app

import (
	"fmt"
	"os"

	"github.com/talbs1986/simplego/logger/pkg/logger"
	zerolog "github.com/talbs1986/simplego/zerolog-logger/pkg/logger"
)

func DefaultLogger(cfg *logger.Config) logger.ILogger {
	l, err := zerolog.NewSimpleZerolog(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "simplego app: failed to initialize default logger, due to: %s", err.Error())
		os.Exit(1)
	}
	return l
}

func DefaultConfigurations(cfg *logger.Config) logger.ILogger {
	l, err := zerolog.NewSimpleZerolog(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "simplego app: failed to initialize default logger, due to: %s", err.Error())
		os.Exit(1)
	}
	return l
}
