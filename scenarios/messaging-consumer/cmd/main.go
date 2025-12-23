package main

import (
	"fmt"

	"github.com/talbs1986/simplego/app/pkg/app"
	configs "github.com/talbs1986/simplego/configs/pkg/app"
	"github.com/talbs1986/simplego/messaging/pkg/messaging"

	"github.com/talbs1986/simplego/scenarios/messaging-consumer/pkg/scenarios"
)

type MyConfig struct{}

func main() {
	scenarios.StartConsumer[MyConfig](proc, buildMsgHandlers)
}

func proc(appObj *app.App) error {
	appObj.Logger.Log().Info("consumer: finnished running user code")
	return nil
}

func buildMsgHandlers(appObj *app.App) ([]scenarios.MsgHandler, error) {
	appObj.Logger.Log().Info("consumer: building msg handlers")
	cfg, err := configs.GetConfig[scenarios.ConsumerConfig[MyConfig]](appObj)
	if err != nil {
		return nil, fmt.Errorf("consumer: failed to get config, due to: %w", err)
	}

	res := []scenarios.MsgHandler{}
	for _, topic := range cfg.ConsumerConfig.Topics {
		res = append(res, scenarios.MsgHandler{
			Topic: topic,
			Proc: func(mw *messaging.MessageWrapper) error {
				appObj.Logger.Log().Info(fmt.Sprintf("consumer: handling msg %s", mw.ID))
				return nil
			},
		})
	}
	return res, nil
}
