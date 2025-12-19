package app

import (
	"fmt"

	simplego "github.com/talbs1986/simplego/app/pkg/app"
	"github.com/talbs1986/simplego/messaging/pkg/messaging"
)

const (
	appKeyServiceConsumerPrefix  = "consumer"
	appKeyServicePublisherPrefix = "publisher"
)

// WithConsumer registers a consumer by name with the app
func WithConsumer(consumerName string, c messaging.IConsumer) simplego.AppOpt {
	return func(s *simplego.App) {
		s.RegisterAppService(fmt.Sprintf("%s-%s", appKeyServiceConsumerPrefix, consumerName), c)
	}
}

// WithPublisher registers a publisher by name with the app
func WithPublisher(publisherName string, p messaging.IPublisher) simplego.AppOpt {
	return func(s *simplego.App) {
		s.RegisterAppService(fmt.Sprintf("%s-%s", appKeyServicePublisherPrefix, publisherName), p)
	}
}
