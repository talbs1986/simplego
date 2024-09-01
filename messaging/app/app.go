package app

import (
	simplego "github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/messaging/pkg/messaging"
)

// GetConsumer gets the registered application consumer
func GetConsumer(consumerName string, s *simplego.App) messaging.IConsumer {
	return s.GetAppService(consumerName).(messaging.IConsumer)
}

// GetPublisher gets the registered application publisher
func GetPublisher(publisherName string, s *simplego.App) messaging.IPublisher {
	return s.GetAppService(publisherName).(messaging.IPublisher)
}
