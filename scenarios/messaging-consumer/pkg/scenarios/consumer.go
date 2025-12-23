package scenarios

import (
	"fmt"

	"github.com/talbs1986/simplego/app/pkg/app"
	simplego_config "github.com/talbs1986/simplego/configs/pkg/app"
	simplego_messaging_app "github.com/talbs1986/simplego/messaging/pkg/app"
	simplego_messaging "github.com/talbs1986/simplego/messaging/pkg/messaging"
	"github.com/talbs1986/simplego/nats-messaging/pkg/messaging"
	scenario_server "github.com/talbs1986/simplego/scenarios/service-app/pkg/scenarios"
)

type ExecutionFunc func(*app.App) error
type MsgHandlersBuilderFunc func(*app.App) ([]MsgHandler, error)
type MsgHandler struct {
	Proc  simplego_messaging.MsgProcessor
	Topic string
}

func StartConsumer[T interface{}](f ExecutionFunc, msgHandlersBuilder MsgHandlersBuilderFunc) {
	var consumeFunc scenario_server.ExecutionFunc = func(appObj *app.App) error {

		cfg, err := simplego_config.GetConfig[ConsumerConfig[T]](appObj)
		if err != nil {
			return err
		}
		consumerNameBuilder := func(topic string) string {
			return fmt.Sprintf("%s-%s", cfg.AppConfig.Name, topic)
		}

		msgHandlers, err := msgHandlersBuilder(appObj)
		if err != nil {
			return fmt.Errorf("simplego consumer: failed to build msg handlers, due to: %w", err)
		}

		for _, msgHandler := range msgHandlers {
			consumerService, err := messaging.NewNATSConsumer(appObj.Logger, &messaging.NATSConsumerConfig{
				ServiceName:         cfg.AppConfig.Name,
				NATSClusterHost:     cfg.ConsumerConfig.Host,
				NATSClusterPassword: cfg.ConsumerConfig.Password,
				NATSClusterUser:     cfg.ConsumerConfig.User,
				MaxPendingMsgs:      cfg.ConsumerConfig.MaxPendingMsgs,
				Destination:         fmt.Sprintf("%s-queue", cfg.AppConfig.Name),
				NATSStreamConfig: &messaging.NATSStreamConfig{
					Name:               msgHandler.Topic,
					AllowedSubjects:    []string{"*"},
					RetentionMaxMsgAge: messaging.DefaultNATSStreamMsgRetention,
					DeduplicateWindow:  messaging.DefaultNATSStreamDeduplicateWindow,
					Replicas:           messaging.DefaultNATSStreamReplicas,
					CompressionEnabled: false,
				},
			})
			if err != nil {
				return fmt.Errorf("simplego consumer: failed to init consumer, due to: %w", err)
			}
			simplego_messaging_app.RegisterConsumer(consumerNameBuilder(msgHandler.Topic), appObj, consumerService)

			err = consumerService.Consume(msgHandler.Topic, msgHandler.Proc)
			if err != nil {
				return fmt.Errorf("simplego consumer: failed to start consumer, due to: %w", err)
			}
		}

		// Execute user code
		err = f(appObj)
		if err != nil {
			appObj.Logger.Log().Error(err, "simplego consumer: finished running with error")
		} else {
			appObj.Logger.Log().Info("simplego consumer: finished running successfully")
		}
		return nil
	}

	scenario_server.StartService[T](consumeFunc)
}
