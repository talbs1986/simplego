package metrics

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/talbs1986/simplego/app/pkg/logger"
	"github.com/talbs1986/simplego/metrics/pkg/metrics"
)

type PromMetricsConfig struct {
	PushInterval *time.Duration
}
type PromMetricsOpt func(s *promMetricsImpl)

type promMetricsImpl struct {
	logger            logger.ILogger
	registerer        *prometheus.Registry
	pusher            metrics.MetricsPusher
	middlewareBuilder metrics.MetricsMiddlewareBuilder
	counters          map[string]metrics.ICounter
	gauges            map[string]metrics.IGauge
	histograms        map[string]metrics.IHistogram
}

func NewPromMetrics(l logger.ILogger, cfg *PromMetricsConfig, opts ...PromMetricsOpt) (metrics.IMetrics, error) {
	s := &promMetricsImpl{
		registerer: prometheus.NewRegistry(),
		logger:     l,
	}

	for _, opt := range opts {
		opt(s)
	}
	if s.pusher == nil && s.middlewareBuilder == nil {
		return nil, metrics.ErrMissingMetricsPusher
	}
	if s.counters == nil {
		s.counters = map[string]metrics.ICounter{}
	}
	if s.histograms == nil {
		s.histograms = map[string]metrics.IHistogram{}
	}
	if s.gauges == nil {
		s.gauges = map[string]metrics.IGauge{}
	}
	if s.pusher != nil {
		if cfg.PushInterval != nil {
			return nil, errors.New("simplego metrics: failed to start prom metrics pusher , missing interval from config")
		}
		go func() {
			if err := s.pusher.Start(*cfg.PushInterval); err != nil {
				l.Log().Fatal(err, "simplego metrics: failed to start prom metrics pusher")
			}
		}()
	}

	return s, nil
}
func (s *promMetricsImpl) Close(ctx context.Context) error {
	if s.pusher != nil {
		return s.pusher.Close(ctx)
	}
	return nil
}
func (s *promMetricsImpl) BuildServerHandler() (http.Handler, error) {
	return promhttp.Handler(), nil
}
func (s *promMetricsImpl) Push(ctx context.Context) error {
	if s.pusher == nil {
		return metrics.ErrMissingMetricsPusher
	}
	return s.pusher.Push(ctx)
}

func (s *promMetricsImpl) RegisterCounter(name string, description string, labels *[]string) error {
	c := newPromCounter(name, description, labels)
	return s.registerer.Register(c.underlying)
}
func (s *promMetricsImpl) RegisterHistogram(name string, description string, buckets []float64, labels *[]string) error {
	h := newPromHistogram(name, description, buckets, labels)
	return s.registerer.Register(h.underlying)
}
func (s *promMetricsImpl) RegisterGauge(name string, description string, labels *[]string) error {
	g := newPromGauge(name, description, labels)
	return s.registerer.Register(g.underlying)
}

func (s *promMetricsImpl) GetMetricsHandler(context.Context) (http.Handler, error) {
	if s.middlewareBuilder == nil {
		return nil, metrics.ErrMissingMetricsMiddlewareBuilder
	}
	return s.middlewareBuilder.BuildServerHandler()
}

func (s *promMetricsImpl) PushCollectedMetrics(ctx context.Context) error {
	if s.pusher == nil {
		return metrics.ErrMissingMetricsPusher
	}
	return s.pusher.Push(ctx)
}
func (s *promMetricsImpl) GetCounter(name string) (metrics.ICounter, error) {
	counter, exists := s.counters[name]
	if !exists {
		return nil, metrics.ErrMissingMetric
	}
	return counter, nil

}
func (s *promMetricsImpl) GetGauge(name string) (metrics.IGauge, error) {
	gauge, exists := s.gauges[name]
	if !exists {
		return nil, metrics.ErrMissingMetric
	}
	return gauge, nil

}
func (s *promMetricsImpl) GetHistogram(name string) (metrics.IHistogram, error) {
	histogram, exists := s.histograms[name]
	if !exists {
		return nil, metrics.ErrMissingMetric
	}
	return histogram, nil

}
