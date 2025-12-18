package main

import (
	"github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/app/pkg/logger"
	simplego_config_app "github.com/talbs1986/simplego/configs/pkg/app"

	"github.com/talbs1986/simplego/scenarios/job-app/pkg/scenarios"
)

type MyConfig struct{}

func main() {
	scenarios.StartJob[MyConfig](
		&scenarios.JobConfig[MyConfig]{},
		proc)
}

func proc(appObj *app.App) error {
	cfg, err := simplego_config_app.GetConfig[scenarios.JobConfig[MyConfig]](appObj)
	if err != nil {
		return err
	}
	appObj.Logger.Log().With(&logger.LogFields{"config": cfg}).Info("job: doing smth...")
	return nil
}
