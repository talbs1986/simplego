package main

import (
	"github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/messaging/pkg/messaging"
	"github.com/talbs1986/simplego/scenarios/messaging-consumer/pkg/scenarios"
)

type MyConfig struct{}

func main() {
	scenarios.StartConsumer[MyConfig](
		&scenarios.ConsumerConfig[MyConfig]{},
		proc,
		map[string]messaging.MsgProcessor{})
}

func proc(appObj *app.App) error {
	appObj.Logger.Log().Info("consumer: finnished running user code")
	return nil
}
