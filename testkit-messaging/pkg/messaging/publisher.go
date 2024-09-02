package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	simplego "github.com/talbs1986/simplego/messaging/pkg/messaging"
)

// TestKitPublisherConfig - testkit publisher config object
type TestKitPublisherConfig struct {
	// ServiceName - unique service name
	ServiceName string `json:"serviceName"`
}

type TestkitPublisher struct {
	sender         string
	DestMessageMap map[string][]*simplego.MessageWrapper
}

// NewTestKitPublisher - creates a new publisher using in memory test kit
func NewTestKitPublisher(cfg *TestKitPublisherConfig) (simplego.IPublisher, error) {
	s := TestkitPublisher{
		sender:         cfg.ServiceName,
		DestMessageMap: map[string][]*simplego.MessageWrapper{},
	}

	return &s, nil
}
func (s *TestkitPublisher) Close(ctx context.Context) error {
	return nil
}
func (s *TestkitPublisher) Publish(dest string, typ simplego.MessageType, payload interface{}) (*simplego.MessageWrapper, error) {
	msg := &simplego.MessageWrapper{
		ID:        uuid.NewString(),
		Sender:    s.sender,
		Timestamp: time.Now().UnixMilli(),
		Type:      typ,
		Payload:   payload,
	}
	_, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("simplego testkit publisher: failed to marshal message, due to: %w", err)
	}
	if _, exists := s.DestMessageMap[dest]; !exists {
		s.DestMessageMap[dest] = []*simplego.MessageWrapper{}
	}
	s.DestMessageMap[dest] = append(s.DestMessageMap[dest], msg)
	return msg, nil
}
