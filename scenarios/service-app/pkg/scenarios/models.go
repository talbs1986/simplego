package scenarios

import (
	"github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/configs/pkg/configs"
	"github.com/talbs1986/simplego/server/pkg/server"
)

type ServiceConfig[T interface{}] struct {
	app.AppConfig
	configs.EnvConfig
	configs.LoggerConfig
	server.ServerConfig
	UserConfig T
}
