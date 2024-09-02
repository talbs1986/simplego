package messaging

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/talbs1986/simplego/app/pkg/logger"
	simplego "github.com/talbs1986/simplego/messaging/pkg/messaging"
)

// NATSConsumerOpt defines the nats consumer option function
type NATSConsumerOpt func(s *natsConsumerImpl)

// NATSConsumerConfig - eventbus nats consumer config object
type NATSConsumerConfig struct {
	// ServiceName - unique service name
	ServiceName string `json:"serviceName"`
	// Destination - stream name to consume from
	Destination string `json:"destination"`
	// NATSClusterHost - nats cluster host
	NATSClusterHost string `json:"natsCluserHost"`
	// NATSClusterUser - nats cluster user
	NATSClusterUser string `json:"natsClusterUser"`
	// NATSClusterPassword - nats cluster password
	NATSClusterPassword string `json:"natsClusterPassword"`
	// MaxPendingMsgs - nats consumer flow control for max pending msgs for all consumers
	MaxPendingMsgs int `json:"maxPendingMsgs"`
	// NATSStreamConfig - nats stream config shared with publisher
	*NATSStreamConfig `json:"streamConfig"`
}

type natsConsumerImpl struct {
	log             logger.ILogger
	conn            *nats.Conn
	stream          nats.JetStreamContext
	subject         string
	consumerGroup   string
	streamName      string
	maxPendingMsgs  int
	sub             *nats.Subscription
	upsertStreamOpt bool
	dest            string
}

// NewNATSConsumer - creates a new eventbus consumer using nats
func NewNATSConsumer(log logger.ILogger, cfg *NATSConsumerConfig, opts ...NATSConsumerOpt) (simplego.IConsumer, error) {
	if len(cfg.ServiceName) < 1 {
		return nil, errors.New("simplego nats consumer: failed to init due to missing service name in config")
	}
	if len(cfg.Destination) < 1 {
		return nil, errors.New("simplego nats consumer: failed to init due to missing destination")
	}
	s := &natsConsumerImpl{
		log:            log,
		consumerGroup:  cfg.ServiceName + "-consumer",
		maxPendingMsgs: cfg.MaxPendingMsgs,
		dest:           cfg.Destination,
	}

	for _, opt := range opts {
		opt(s)
	}
	if s.conn == nil {
		natsConn, err := nats.Connect(cfg.NATSClusterHost, nats.UserInfo(cfg.NATSClusterUser, cfg.NATSClusterPassword),
			nats.Name(s.consumerGroup))
		if err != nil {
			return nil, fmt.Errorf("simplego nats consumer: failed to connect to nats cluster, due to: %w", err)
		}
		s.conn = natsConn
	}
	if s.stream == nil {
		js, err := s.conn.JetStream()
		if err != nil {
			return nil, fmt.Errorf("simplego nats consumer: failed to connect to nats jet stream, due to: %w", err)
		}
		s.stream = js
	}
	if s.upsertStreamOpt {
		s.log.Log().Info("simplego nats consumer: upserting stream config")
		if err := BuildDefaultNATSStream(cfg.NATSStreamConfig, s.stream); err != nil {
			return nil, err
		}
	}

	if cfg.MaxPendingMsgs < 1 {
		s.log.Log().Info("simplego nats consumer: setting default max pending messages")
		s.maxPendingMsgs = DefaultConsumerMaxPendingMsgs
	}
	return s, nil
}

func (s *natsConsumerImpl) Consume(subject string, proc simplego.MsgProcessor) error {
	qname := fmt.Sprintf("%s.%s.%s", s.dest, s.consumerGroup, s.subject)
	sub, err := s.stream.QueueSubscribe(subject, qname, s.buildHandleMsg(proc), nats.Durable(s.consumerGroup),
		nats.BindStream(s.streamName), nats.MaxAckPending(s.maxPendingMsgs))
	if err != nil {
		return fmt.Errorf("simplego nats consumer: '%s' failed to consume from stream '%s' and subject '%s' ,due to: %w", s.consumerGroup, s.dest, s.subject, err)
	}
	s.sub = sub
	return nil
}

func (s *natsConsumerImpl) Pull(subject string, maxMsgBatch int, proc simplego.MsgProcessor) ([]*simplego.MessageWrapper, error) {
	// TODO create map of subjects
	if s.sub == nil {
		sub, err := s.stream.PullSubscribe(subject, s.consumerGroup, nats.BindStream(s.streamName), nats.MaxAckPending(s.maxPendingMsgs))
		if err != nil {
			return nil, fmt.Errorf("simplego nats consumer: '%s' failed to pull from stream '%s' and subject '%s' ,due to: %w", s.consumerGroup, s.dest, s.subject, err)
		}
		s.sub = sub
	}
	natsBatch, err := s.sub.FetchBatch(maxMsgBatch)
	if err != nil {
		return nil, fmt.Errorf("simplego nats consumer: '%s' failed to fetch (%d) messages from stream '%s' and subject '%s',due to: %w", s.consumerGroup, maxMsgBatch, s.dest, subject, err)
	}
	batch := natsBatch.Messages()
	msgs := make([]*simplego.MessageWrapper, len(batch))
	i := 0
	for m := range batch {
		res := s.handleMessage(proc, m)
		if res == nil {
			return nil, fmt.Errorf("simplego nats consumer: '%s' failed to handle message index (%d) from pull batch of stream '%s' and subject '%s'", s.consumerGroup, i, s.dest, subject)
		}
		msgs = append(msgs, res)
	}
	return msgs, nil
}

// Status - returns current unacked count
func (s *natsConsumerImpl) Status(ctx context.Context) (uint64, error) {
	info, err := s.stream.ConsumerInfo(s.streamName, s.consumerGroup)
	if err != nil {
		return 0, fmt.Errorf("simplego nats consumer: '%s' failed to return status from stream '%s' ,due to: %w", s.consumerGroup, s.streamName, err)
	}
	return info.NumPending, nil
}

func (s *natsConsumerImpl) buildHandleMsg(proc simplego.MsgProcessor) func(msg *nats.Msg) {
	return func(natsMsg *nats.Msg) {
		s.handleMessage(proc, natsMsg)
	}
}

func (s *natsConsumerImpl) handleMessage(proc simplego.MsgProcessor, natsMsg *nats.Msg) *simplego.MessageWrapper {
	msg, err := s.unmarshalMsg(natsMsg)
	if err != nil {
		s.log.Log().Error(err, "simplego nats consumer: failed to unmarshal msg")
		if err := natsMsg.Nak(); err != nil {
			s.log.Log().Error(err, "simplego nats consumer: failed to nack msg after unmarshal")
			return nil
		}
		return nil
	}
	if err := proc(msg); err != nil {
		s.log.Log().Error(err, "simplego nats consumer: failed to proc msg")
		if err := natsMsg.Nak(); err != nil {
			s.log.Log().Error(err, "simplego nats consumer: failed to nack msg after proc")
			return nil
		}
		return nil
	}
	if err := natsMsg.Ack(); err != nil {
		s.log.Log().Error(err, "simplego nats consumer: failed to ack msg")
		if err := natsMsg.Nak(); err != nil {
			s.log.Log().Error(err, "simplego nats consumer: failed to nack msg after ack")
			return nil
		}
		return nil
	}
	return msg
}

func (s *natsConsumerImpl) Close(ctx context.Context) error {
	s.log.Log().Info("simplego nats consumer: closing")
	s.conn.Close()
	return nil
}

func (s *natsConsumerImpl) unmarshalMsg(natsMsg *nats.Msg) (*simplego.MessageWrapper, error) {
	var msg simplego.MessageWrapper
	if err := json.Unmarshal(natsMsg.Data, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
