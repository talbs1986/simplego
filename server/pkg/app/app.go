package app

import (
	simplego "github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/server/pkg/server"
)

// GetServerService gets the registered application server service
func GetServerService(s *simplego.App) server.IServer {
	return s.GetAppService(appKeyServiceServer).(server.IServer)
}
