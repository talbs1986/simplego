package app

import (
	simplego "github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/metrics/pkg/metrics"
)

// GetMetricsService gets the registered application metrics service
func GetMetricsService(s *simplego.App) metrics.IMetrics {
	return s.GetAppService(appKeyServiceMetrics).(metrics.IMetrics)
}
