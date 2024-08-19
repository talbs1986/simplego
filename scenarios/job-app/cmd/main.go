package main

import (
	"time"

	"github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/configs/pkg/configs"
	"github.com/talbs1986/simplego/scenarios/job-app/pkg/scenarios"
)

type AppEnvConfig struct {
	configs.EnvConfig
	configs.LoggerConfig
}

func main() {
	scenarios.StartJob[AppEnvConfig](
		&app.AppConfig{
			Name:                "job-exmaple",
			Version:             "1.0.0",
			ServiceCloseTimeout: time.Second * 30,
		},
		proc)
}

func proc(appObj *app.App) error {
	appObj.Logger.Log().Info("simplego job: doing smth...")
	return nil
}
