package messaging

import (
	"context"

	simplego "github.com/talbs1986/simplego/app/pkg/app"
)

// IConsumer - message consumer api
type IConsumer interface {
	//Consume - start consume messages
	Consume(eventType MessageType, proc MsgProcessor) error
	//Status - returns current unacked count
	Status(ctx context.Context) (int64, error)
	//Pull - pulls sync messages
	Pull(eventType MessageType, maxMsgBatch int) ([]MessageWrapper, error)
	// CloseableService clean and close resources api
	simplego.CloseableService
}

// MsgProcessor - consumer msg processor function, note: should be thread safe
type MsgProcessor func(msg *MessageWrapper) error
