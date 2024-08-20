package scenarios

import (
	"github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/configs/pkg/configs"
	"github.com/talbs1986/simplego/metrics/pkg/metrics"
)

type JobConfig[T interface{}] struct {
	app.AppConfig
	configs.EnvConfig
	configs.LoggerConfig
	metrics.MetricsPushConfig
	UserConfig T
}
