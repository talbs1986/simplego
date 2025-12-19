package scenarios

import (
	"github.com/talbs1986/simplego/messaging/pkg/messaging"
	scenario_server "github.com/talbs1986/simplego/scenarios/service-app/pkg/scenarios"
)

type ConsumerConfig[T interface{}] struct {
	scenario_server.ServiceConfig[T]
	messaging.ConsumerConfig
}
