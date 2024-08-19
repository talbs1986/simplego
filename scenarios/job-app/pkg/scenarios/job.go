package scenarios

import (
	"fmt"

	"github.com/talbs1986/simplego/app/pkg/app"
	simplego_config "github.com/talbs1986/simplego/configs/pkg/app"
	"github.com/talbs1986/simplego/goenv-configs/pkg/configs"
	"github.com/talbs1986/simplego/zerolog-logger/pkg/logger"
)

type ExecutionFunc func(*app.App) error

func StartJob[T interface{}](cfg *app.AppConfig, f ExecutionFunc) {
	cfgParser, err := configs.NewGoEnvConfigParser[T]()
	if err != nil {
		panic(fmt.Errorf("simplego job: failed to init config parser, due to: %w", err))
	}
	appObj := app.NewApp(cfg,
		simplego_config.WithConfigParser(cfgParser),
		app.WithLogger(logger.DefaultLoggerWithDefaultConfig()))
	appObj.Logger.Log().Info("simplego job: app started")

	if err := f(appObj); err != nil {
		appObj.Logger.Log().Error(err, "simplego job: finished running with error")
	} else {
		appObj.Logger.Log().Info("simplego job: finished running successfully")
	}
	appObj.Stop()
	appObj.Logger.Log().Info("simplego job: gracefully shutting down")
}
