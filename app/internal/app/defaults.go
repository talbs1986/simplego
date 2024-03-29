package app

import (
	"context"
	"fmt"
	"os"

	"github.com/talbs1986/simplego/configs/pkg/configs"
	goenvconfig "github.com/talbs1986/simplego/goenv-configs/pkg/configs"
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

func DefaultConfigurations[T interface{}](ctx context.Context) configs.IConfigs[T] {
	p, err := goenvconfig.NewGoEnvConfigParser[T]()
	if err != nil {
		fmt.Fprintf(os.Stderr, "simplego app: failed to initialize default configuration parser, due to: %s", err.Error())
		os.Exit(1)
	}
	parsers := []configs.ConfigParser[T]{p}
	c, err := configs.NewConfigs[T](ctx, configs.WithConfigParsers[T](parsers))
	if err != nil {
		fmt.Fprintf(os.Stderr, "simplego app: failed to initialize default configuration, due to: %s", err.Error())
		os.Exit(1)
	}
	return c
}
