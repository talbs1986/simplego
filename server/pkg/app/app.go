package app

import (
	simplego "github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/server/pkg/server"
)

// GetMetricsService gets the registered application metrics service
func GetMetricsService(s *simplego.App) server.IServer {
	return s.GetAppService(appKeyServiceServer).(server.IServer)
}
