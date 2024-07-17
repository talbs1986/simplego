package logger

import (
	"fmt"
	"os"

	"github.com/talbs1986/simplego/app/pkg/logger"
)

// DefaultLogger default zerolog logger builder with config
func DefaultLogger(cfg *logger.Config) logger.ILogger {
	l, err := NewSimpleZerolog(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "simplego logger: failed to initialize default zerolog logger, due to: %s", err.Error())
		os.Exit(1)
	}
	return l
}

// DefaultLoggerWithDefaultConfig default zerolog logger builder with default config
func DefaultLoggerWithDefaultConfig() logger.ILogger {
	l, err := NewSimpleZerolog(logger.DefaultConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "simplego logger: failed to initialize default zerolog logger, due to: %s", err.Error())
		os.Exit(1)
	}
	return l
}
