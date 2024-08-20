package metrics

import "time"

// MetricsPushConfig defines the basic metrics pusher config
type MetricsPushConfig struct {
	PushGatewayHost string        `env:"METRIC_PUSH_GATEWAY_HOST, default=localhost"`
	PushInterval    time.Duration `env:"METRIC_PUSH_INTERVAL, default=5s"`
}
