package messaging

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/talbs1986/simplego/app/pkg/logger"
	simplego "github.com/talbs1986/simplego/messaging/pkg/messaging"
)

// NATSPublisherConfig - nats publisher config object
type NATSPublisherConfig struct {
	// ServiceName - unique service name
	ServiceName string `json:"serviceName"`
	// NATSClusterHost - nats cluster host
	NATSClusterHost string `json:"natsCluserHost"`
	// NATSClusterUser - nats cluster user
	NATSClusterUser string `json:"natsClusterUser"`
	// NATSClusterPassword - nats cluster password
	NATSClusterPassword string `json:"natsClusterPassword"`
	// NATSStreamConfig - nats stream config to upsert stream settings
	*NATSStreamConfig `json:"streamConfig"`
}

type natsPublisherImpl struct {
	conn            *nats.Conn
	stream          nats.JetStreamContext
	sender          string
	upsertStreamOpt bool
	log             logger.ILogger
}

// NewNATSPublisher - creates a new messaging publisher using nats
func NewNATSPublisher(log logger.ILogger, cfg *NATSPublisherConfig, opts ...NATSPublisherOpt) (simplego.IPublisher, error) {
	if len(cfg.ServiceName) < 1 {
		return nil, errors.New("simplego nats publisher: failed to init due to missing service name in config")
	}
	s := &natsPublisherImpl{
		log:    log,
		sender: cfg.ServiceName,
	}
	for _, opt := range opts {
		opt(s)
	}
	if s.conn == nil {
		natsConn, err := nats.Connect(cfg.NATSClusterHost, nats.UserInfo(cfg.NATSClusterUser, cfg.NATSClusterPassword),
			nats.Name(cfg.ServiceName+"-publisher"))
		if err != nil {
			return nil, fmt.Errorf("simplego nats publisher: failed to connect to nats cluster, due to: %w", err)
		}
		s.conn = natsConn
	}
	if s.stream == nil {
		js, err := s.conn.JetStream()
		if err != nil {
			return nil, fmt.Errorf("simplego nats publisher: failed to connect to nats jet stream, due to: %w", err)
		}
		s.stream = js
	}
	if s.upsertStreamOpt {
		s.log.Log().Info("simplego nats publisher: upserting stream config")
		if err := BuildDefaultNATSStream(cfg.NATSStreamConfig, s.stream); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (s *natsPublisherImpl) Publish(dest string, typ simplego.MessageType, payload interface{}) (*simplego.MessageWrapper, error) {
	msg := &simplego.MessageWrapper{
		ID:        uuid.NewString(),
		Sender:    s.sender,
		Timestamp: time.Now().UnixMilli(),
		Type:      typ,
		Payload:   payload,
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("simplego nats publisher: failed to marshal message, due to: %w", err)
	}
	_, err = s.stream.Publish(dest, data)
	if err != nil {
		return nil, fmt.Errorf("simplego nats publisher: failed to publish message, due to: %w", err)
	}
	return msg, nil
}

func (s *natsPublisherImpl) Close(ctx context.Context) error {
	fmt.Fprintf(os.Stdout, "simplego nats publisher: closing")
	s.conn.Close()
	return nil
}
