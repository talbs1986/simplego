package app

import (
	"context"
	"fmt"
	"os"

	goenvconfig "github.com/talbs1986/simplego/configs-goenv/pkg/configs"
	"github.com/talbs1986/simplego/configs/pkg/configs"
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
	c, err := configs.NewConfigs[T](ctx, configs.WithConfigParsers([]configs.ConfigParser{goenvconfig.NewGoEnvConfigParser[T]()}))
	if err != nil {
		fmt.Fprintf(os.Stderr, "simplego app: failed to initialize default configuration parser, due to: %s", err.Error())
		os.Exit(1)
	}
	return c
}
