package scenarios

import (
	"github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/app/pkg/logger"
	"github.com/talbs1986/simplego/server/pkg/server"
)

type ServiceConfig[T interface{}] struct {
	app.AppConfig
	app.EnvConfig
	logger.LoggerConfig
	server.ServerConfig
	UserConfig T
}
