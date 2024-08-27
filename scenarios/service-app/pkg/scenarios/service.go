package scenarios

import (
	"context"
	"fmt"

	"github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/chi-server/pkg/server"
	simplego_config "github.com/talbs1986/simplego/configs/pkg/app"
	"github.com/talbs1986/simplego/goenv-configs/pkg/configs"
	simplego_metrics "github.com/talbs1986/simplego/metrics/pkg/app"
	"github.com/talbs1986/simplego/prom-metrics/pkg/metrics"
	scenario_server "github.com/talbs1986/simplego/scenarios/service-app/pkg/server"
	simplego_server "github.com/talbs1986/simplego/server/pkg/app"
	"github.com/talbs1986/simplego/zerolog-logger/pkg/logger"
)

type ExecutionFunc func(*app.App) error

func StartService[T interface{}](cfg *ServiceConfig[T], f ExecutionFunc) {
	// parsing config from env
	cfgParser, err := configs.NewGoEnvConfigParser[ServiceConfig[T]]()
	if err != nil {
		panic(fmt.Errorf("simplego service: failed to init config parser, due to: %w", err))
	}
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ServiceCloseTimeout)
	parsedEnvConfig, err := cfgParser.Parse(ctx)
	if err != nil {
		panic(fmt.Errorf("simplego service: failed to parse env config, due to: %w", err))
	}
	cancel()

	// init logger
	logger := logger.DefaultLoggerWithDefaultConfig()

	// init metrics
	metricService, err := metrics.NewPromMetrics(logger, &metrics.PromMetricsConfig{},
		metrics.WithDefaultMetricsMiddleware())
	if err != nil {
		panic(fmt.Errorf("simplego service: failed to init metrics service, due to: %w", err))
	}

	// init server
	serverService, err := server.NewChiServer(logger, &parsedEnvConfig.ServerConfig,
		server.WithMiddlewares(scenario_server.BuildDefaultServerMiddlewares(parsedEnvConfig.RequestTimeout, logger, metricService)),
		server.WithRoutes(scenario_server.BuildDefaultServerRoutes(metricService)))
	if err != nil {
		panic(fmt.Errorf("simplego service: failed to init server, due to: %w", err))
	}
	// init app
	appObj := app.NewApp(&parsedEnvConfig.AppConfig,
		simplego_config.WithConfigParser(cfgParser),
		app.WithLogger(logger),
		simplego_metrics.WithMetrics(metricService),
		simplego_server.WithServer(serverService),
	)

	// process stuff
	if err := f(appObj); err != nil {
		appObj.Logger.Log().Error(err, "simplego job: finished running with error")
	} else {
		appObj.Logger.Log().Info("simplego job: finished running successfully")
	}

	// start server
	if err := serverService.Start(); err != nil {
		panic(fmt.Errorf("simplego service: failed to start server, due to: %w", err))
	}
	appObj.Logger.Log().Info("simplego service: app started")
	// gracefully shutdown
	appObj.WaitForShutodwn()
	appObj.Stop()
	appObj.Logger.Log().Info("simplego service: gracefully shutting down")
}
