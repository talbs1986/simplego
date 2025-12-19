package messaging

import (
	"errors"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	DefaultNATSStreamMsgRetention      = time.Hour * 24 * 3
	DefaultNATSStreamDeduplicateWindow = time.Minute * 10
	DefaultNATSStreamReplicas          = 1
	DefaultConsumerMaxPendingMsgs      = 1000
	DefaultNATSStreamName              = "simplego-messages"
)

var DefaultNATSStreamConfig = NATSStreamConfig{
	RetentionMaxMsgAge: DefaultNATSStreamMsgRetention,
	DeduplicateWindow:  DefaultNATSStreamDeduplicateWindow,
	Replicas:           DefaultNATSStreamReplicas,
	AllowedSubjects:    []string{"*"},
	Name:               DefaultNATSStreamName,
	CompressionEnabled: false,
}

// NATSStreamConfig - nats stream configuration object
type NATSStreamConfig struct {
	// RetentionMaxMsgAge - nats stream max msg age to keep in stream
	RetentionMaxMsgAge time.Duration `json:"retentionMaxMsgAge"`
	// DeduplicateWindow - nats stream dedupe window time
	DeduplicateWindow time.Duration `json:"deduplicateWindow"`
	// Replicas - nats stream msg replicas
	Replicas int `json:"replicas"`
	// CompressionEnabled - nats stream message compression
	CompressionEnabled bool `json:"compressionEnabled"`
	// AllowedSubjects - nats stream allowed subjects config
	AllowedSubjects []string `json:"allowedSubjects"`
	// Name - nats stream name
	Name string `json:"name"`
}

func BuildDefaultNATSStream(cfg *NATSStreamConfig, js nats.JetStreamContext) error {
	if cfg == nil {
		cfg = &DefaultNATSStreamConfig
	}
	config := &nats.StreamConfig{
		Name:        cfg.Name,
		Description: cfg.Name,
		Subjects:    cfg.AllowedSubjects,
		Storage:     nats.FileStorage,
		Retention:   nats.WorkQueuePolicy,
		Duplicates:  cfg.DeduplicateWindow,
		Replicas:    cfg.Replicas,
		MaxAge:      cfg.RetentionMaxMsgAge,
	}

	if cfg.CompressionEnabled {
		config.Compression = nats.S2Compression
	}

	// try to create the stream
	_, err := js.AddStream(config)
	// The stream already exists - so that means we
	// probably changed the settings
	if err != nil {
		if errors.Is(err, nats.ErrStreamNameAlreadyInUse) {
			_, err = js.UpdateStream(config)
			if err != nil {
				return fmt.Errorf("simplego nats publisher: failed to update stream config, due to: %w", err)
			}
			return nil
		}
		return fmt.Errorf("simplego nats publisher: failed to create stream, due to: %w", err)
	}
	return nil
}
