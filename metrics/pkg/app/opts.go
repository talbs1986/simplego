package app

import (
	simplego "github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/metrics/pkg/metrics"
)

const (
	appKeyServiceMetrics = "metrics"
)

// WithMetrics registers the metric service with the app
func WithMetrics(m metrics.IMetrics) simplego.AppOpt {
	return func(s *simplego.App) {
		s.RegisterAppService(appKeyServiceMetrics, m)
	}
}
