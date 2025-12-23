package messaging

import (
	simplego "github.com/talbs1986/simplego/app/pkg/app"
)

// IPublisher - message publisher api
type IPublisher interface {
	// Publish - creates a new message from the payload and publish it to dest
	Publish(string, MessageType, interface{}) (*MessageWrapper, error)
	// CloseableService clean and close resources api
	simplego.CloseableService
}
