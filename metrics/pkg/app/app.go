package app

import (
	simplego "github.com/talbs1986/simplego/app/pkg/app"
)

// GetMetricsService gets the registered application metrics service
func GetMetricsService(s *simplego.App) interface{} {
	return s.GetAppService(appKeyServiceMetrics)
}
