package messaging

import (
	"github.com/nats-io/nats.go"
)

// NATSConsumerOpt defines the nats consumer option function
type NATSConsumerOpt func(s *natsConsumerImpl)

// NATSPublisherOpt defines the nats publisher option function
type NATSPublisherOpt func(s *natsPublisherImpl)

// WithPublisherUpsertStream publisher option for upsert the nats stream config
func WithPublisherUpsertStream() NATSPublisherOpt {
	return func(s *natsPublisherImpl) {
		s.upsertStreamOpt = true
	}
}

// WithPublisherNATSConnection publisher option for setting the nats connection
func WithPublisherNATSConnection(c *nats.Conn) NATSPublisherOpt {
	return func(s *natsPublisherImpl) {
		s.conn = c
	}
}

// WithPublisherNATSStream publisher option for setting the nats jet stream context
func WithPublisherNATSStream(js nats.JetStreamContext) NATSPublisherOpt {
	return func(s *natsPublisherImpl) {
		s.stream = js
	}
}

// WitConsumerUpsertStream consumer option for upsert the nats stream config
func WitConsumerUpsertStream() NATSConsumerOpt {
	return func(s *natsConsumerImpl) {
		s.upsertStreamOpt = true
	}
}

// WithConsumerNATSConnection consumer option for setting the nats connection
func WithConsumerNATSConnection(c *nats.Conn) NATSConsumerOpt {
	return func(s *natsConsumerImpl) {
		s.conn = c
	}
}

// WithConsumerNATSStream consumer option for setting the nats jet stream context
func WithConsumerNATSStream(js nats.JetStreamContext) NATSConsumerOpt {
	return func(s *natsConsumerImpl) {
		s.stream = js
	}
}

// WithConsumerNATSSubscription consumer option for setting the nats subscription
func WithConsumerNATSSubscription(subject string, ss *nats.Subscription) NATSConsumerOpt {
	return func(s *natsConsumerImpl) {
		s.sub[subject] = ss
	}
}
