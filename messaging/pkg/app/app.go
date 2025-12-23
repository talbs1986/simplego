package app

import (
	"fmt"

	simplego "github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/messaging/pkg/messaging"
)

// GetConsumer gets the registered application consumer
func GetConsumer(consumerName string, s *simplego.App) messaging.IConsumer {
	return s.GetAppService(fmt.Sprintf("%s-%s", appKeyServiceConsumerPrefix, consumerName)).(messaging.IConsumer)
}

// RegisterConsumer register application consumer
func RegisterConsumer(consumerName string, s *simplego.App, c messaging.IConsumer) {
	s.RegisterAppService(fmt.Sprintf("%s-%s", appKeyServiceConsumerPrefix, consumerName), c)
}

// GetPublisher gets the registered application publisher
func GetPublisher(publisherName string, s *simplego.App) messaging.IPublisher {
	return s.GetAppService(fmt.Sprintf("%s-%s", appKeyServicePublisherPrefix, publisherName)).(messaging.IPublisher)
}

// RegisterPublisher register application publisher
func RegisterPublisher(publisherName string, s *simplego.App, p messaging.IPublisher) {
	s.RegisterAppService(fmt.Sprintf("%s-%s", appKeyServicePublisherPrefix, publisherName), p)
}
