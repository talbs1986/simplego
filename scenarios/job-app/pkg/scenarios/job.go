package scenarios

import (
	"context"
	"fmt"

	"github.com/talbs1986/simplego/app/pkg/app"
	simplego_config "github.com/talbs1986/simplego/configs/pkg/app"
	"github.com/talbs1986/simplego/goenv-configs/pkg/configs"
	simplego_metrics "github.com/talbs1986/simplego/metrics/pkg/app"
	"github.com/talbs1986/simplego/prom-metrics/pkg/metrics"
	"github.com/talbs1986/simplego/zerolog-logger/pkg/logger"
)

type ExecutionFunc func(*app.App) error

func StartJob[T interface{}](cfg *JobConfig[T], f ExecutionFunc) {
	// parsing config from env
	cfgParser, err := configs.NewGoEnvConfigParser[JobConfig[T]]()
	if err != nil {
		panic(fmt.Errorf("simplego job: failed to init config parser, due to: %w", err))
	}
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ServiceCloseTimeout)
	parsedEnvConfig, err := cfgParser.Parse(ctx)
	if err != nil {
		panic(fmt.Errorf("simplego job: failed to parse env config, due to: %w", err))
	}
	cancel()

	// init logger
	logger := logger.DefaultLoggerWithDefaultConfig()

	// init metrics
	metricService, err := metrics.NewPromMetrics(logger, &metrics.PromMetricsConfig{
		PushInterval: &parsedEnvConfig.PushInterval,
	}, metrics.WithDefaultPusher(parsedEnvConfig.PushGatewayHost, parsedEnvConfig.Name))
	if err != nil {
		panic(fmt.Errorf("simplego job: failed to init metrics service, due to: %w", err))
	}

	// init app
	appObj := app.NewApp(&parsedEnvConfig.AppConfig,
		simplego_config.WithConfigParser(cfgParser),
		app.WithLogger(logger),
		simplego_metrics.WithMetrics(metricService),
	)
	appObj.Logger.Log().Info("simplego job: app started")

	// process stuff
	if err := f(appObj); err != nil {
		appObj.Logger.Log().Error(err, "simplego job: finished running with error")
	} else {
		appObj.Logger.Log().Info("simplego job: finished running successfully")
	}

	// gracefully shutdown
	appObj.Stop()
	appObj.Logger.Log().Info("simplego job: gracefully shutting down")
}
