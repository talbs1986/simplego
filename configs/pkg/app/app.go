package app

import (
	simplego "github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/configs/pkg/configs"
)

// GetConfig gets the application config from the parser
func GetConfig[T interface{}](s *simplego.App) (*T, error) {
	return s.GetAppService(appKeyServiceConfigs).(configs.ConfigParser[T]).Parse(s.CTX)
}
