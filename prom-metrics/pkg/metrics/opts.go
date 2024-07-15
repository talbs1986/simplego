package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	simplego "github.com/talbs1986/simplego/metrics/pkg/metrics"
)

// WithDefaultPusher creates a new prometheus gateway pusher
func WithDefaultPusher(host string, serviceName string) PromMetricsOpt {
	return func(s *promMetricsImpl) {
		s.pusher = newPusher(s.logger, host, serviceName)
	}
}

// WithDefaultPusher creates a new prometheus http handler
func WithDefaultMetricsMiddleware() PromMetricsOpt {
	return func(s *promMetricsImpl) {
		s.metricsHTTPHandler = promhttp.Handler()
	}
}

// WithPusher sets the metrics pusher
func WithPusher(p simplego.MetricsPusher) PromMetricsOpt {
	return func(s *promMetricsImpl) {
		s.pusher = p
	}
}

// WithHTTPMetricsHandler sets the http metrics handle builder
func WithHTTPMetricsHandler(b simplego.MetricsMiddlewareBuilder) PromMetricsOpt {
	return func(s *promMetricsImpl) {
		h, err := b.BuildServerHandler()
		if err != nil {
			s.logger.Log().Fatal(err, "simplego metrics: failed to build metrics server handler")
			return
		}
		s.metricsHTTPHandler = h
	}
}
