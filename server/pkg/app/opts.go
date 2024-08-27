package app

import (
	simplego "github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/server/pkg/server"
)

const (
	appKeyServiceServer = "server"
)

// WithServer registers the metric service with the app
func WithServer(m server.IServer) simplego.AppOpt {
	return func(s *simplego.App) {
		s.RegisterAppService(appKeyServiceServer, m)
	}
}
