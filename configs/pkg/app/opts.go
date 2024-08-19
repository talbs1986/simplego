package app

import (
	simplego "github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/configs/pkg/configs"
)

const (
	appKeyServiceConfigs = "config"
)

// WithConfigParser registers the config parser service with the app
func WithConfigParser[T interface{}](m configs.ConfigParser[T]) simplego.AppOpt {
	return func(s *simplego.App) {
		s.RegisterAppService(appKeyServiceConfigs, m)
	}
}
