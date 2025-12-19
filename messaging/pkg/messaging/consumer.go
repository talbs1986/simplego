package messaging

import (
	"context"
	"errors"

	simplego "github.com/talbs1986/simplego/app/pkg/app"
)

var ErrConsumerAlreadyExists = errors.New("simplego consumer: already exists")

// IConsumer - message consumer api
type IConsumer interface {
	// Consume - start consume messages
	Consume(string, MsgProcessor) error
	// Status - returns current unacked count
	Status(context.Context) (uint64, error)
	// Pull - pulls sync messages
	Pull(string, int, MsgProcessor) ([]*MessageWrapper, error)
	// CloseableService clean and close resources api
	simplego.CloseableService
}

// MsgProcessor - consumer msg processor function, note: should be thread safe
type MsgProcessor func(*MessageWrapper) error
