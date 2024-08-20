package main

import (
	"time"

	"github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/app/pkg/logger"
	simplego_config "github.com/talbs1986/simplego/configs/pkg/app"
	"github.com/talbs1986/simplego/scenarios/job-app/pkg/scenarios"
)

type SomeConfig struct {
}

func main() {
	scenarios.StartJob[SomeConfig](
		&scenarios.JobConfig[SomeConfig]{
			AppConfig: app.AppConfig{
				Name:                "job-exmaple",
				Version:             "1.0.0",
				ServiceCloseTimeout: time.Second * 30,
			},
		},
		proc)
}

func proc(appObj *app.App) error {
	cfg, err := simplego_config.GetConfig[SomeConfig](appObj)
	if err != nil {
		return err
	}
	appObj.Logger.Log().With(&logger.LogFields{"config": cfg}).Info("simplego job: doing smth...")
	return nil
}
