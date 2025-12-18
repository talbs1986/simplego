package scenarios

import (
	"github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/app/pkg/logger"
	"github.com/talbs1986/simplego/metrics/pkg/metrics"
)

type JobConfig[T interface{}] struct {
	app.AppConfig
	app.EnvConfig
	logger.LoggerConfig
	metrics.MetricsPushConfig
	UserConfig T
}
