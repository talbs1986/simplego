package messaging

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	simplego "github.com/talbs1986/simplego/messaging/pkg/messaging"
)

type testkitConsumer struct {
	done   chan bool
	stream map[string]chan *simplego.MessageWrapper
}

// NewTestKitConsumer - creates a new consumer using in memory testkit
func NewTestKitConsumer() (simplego.IConsumer, error) {
	s := testkitConsumer{
		stream: map[string]chan *simplego.MessageWrapper{},
		done:   make(chan bool),
	}
	return &s, nil
}

func (s *testkitConsumer) Consume(dest string, proc simplego.MsgProcessor) error {
	msgStream, exists := s.stream[dest]
	if !exists {
		return fmt.Errorf("simplego teskit consumer: failed to find msg stream for - '%s'", dest)
	}
	go func() {
		for {
			select {
			case <-s.done:
				return
			case msg := <-msgStream:
				if msg == nil {
					continue
				}
			SIM_ACK:
				for {
					if err := proc(msg); err != nil {
						<-time.After(time.Millisecond * 10)
						continue
					}
					break SIM_ACK
				}
			}
		}
	}()
	return nil
}

func (s *testkitConsumer) Pull(dest string, maxMsgBatch int, proc simplego.MsgProcessor) ([]*simplego.MessageWrapper, error) {
	return nil, errors.New("simplego testkit publisher: pull messages isnt implemented")
}

// Status - returns current unacked count
func (s *testkitConsumer) Status(ctx context.Context) (int64, error) {
	count := int64(0)
	for _, stream := range s.stream {
		count += int64(len(stream))
	}
	return count, nil
}

func (s *testkitConsumer) Close(ctx context.Context) error {
	fmt.Fprintf(os.Stdout, "simplego testkit consumer: closing")
	s.done <- true
	return nil
}
