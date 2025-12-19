package scenarios

import (
	"fmt"

	"github.com/talbs1986/simplego/app/pkg/app"
	simplego_messaging_app "github.com/talbs1986/simplego/messaging/pkg/app"
	simplego_messaging "github.com/talbs1986/simplego/messaging/pkg/messaging"
	"github.com/talbs1986/simplego/nats-messaging/pkg/messaging"
	scenario_server "github.com/talbs1986/simplego/scenarios/service-app/pkg/scenarios"
)

type ExecutionFunc func(*app.App) error

func StartConsumer[T interface{}](cfg *ConsumerConfig[T], f ExecutionFunc, msgHandlers map[string]simplego_messaging.MsgProcessor) {

	var consumeFunc scenario_server.ExecutionFunc = func(a *app.App) error {
		// init consumers
		consumerNameBuilder := func(topic string) string {
			return fmt.Sprintf("%s-%s", cfg.AppConfig.Name, topic)
		}

		for topic := range msgHandlers {
			consumerService, err := messaging.NewNATSConsumer(a.Logger, &messaging.NATSConsumerConfig{
				ServiceName:         cfg.AppConfig.Name,
				NATSClusterHost:     cfg.ConsumerConfig.Host,
				NATSClusterPassword: cfg.ConsumerConfig.Password,
				NATSClusterUser:     cfg.ConsumerConfig.User,
				MaxPendingMsgs:      cfg.ConsumerConfig.MaxPendingMsgs,
				Destination:         fmt.Sprintf("%s-queue", cfg.AppConfig.Name),
				NATSStreamConfig: &messaging.NATSStreamConfig{
					Name:               topic,
					AllowedSubjects:    []string{"*"},
					RetentionMaxMsgAge: messaging.DefaultNATSStreamMsgRetention,
					DeduplicateWindow:  messaging.DefaultNATSStreamDeduplicateWindow,
					Replicas:           messaging.DefaultNATSStreamReplicas,
					CompressionEnabled: false,
				},
			})
			if err != nil {
				panic(fmt.Errorf("simplego consumer: failed to init consumer, due to: %w", err))
			}
			simplego_messaging_app.RegisterConsumer(consumerNameBuilder(topic), a, consumerService)
		}

		// Execute user code
		err := f(a)
		if err != nil {
			a.Logger.Log().Error(err, "simplego consumer: finished running with error")
		} else {
			a.Logger.Log().Info("simplego consumer: finished running successfully")
		}
		return nil
	}

	scenario_server.StartService[T](&cfg.ServiceConfig, consumeFunc)

}
