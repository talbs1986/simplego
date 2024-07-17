package metrics

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/talbs1986/simplego/app/pkg/logger"
	simplego "github.com/talbs1986/simplego/metrics/pkg/metrics"
)

var instance simplego.IMetrics

// PromMetricsConfig defines the prometheus metrics configuration object
type PromMetricsConfig struct {
	PushInterval *time.Duration
}

// PromMetricsOpt defines the prometheus metrics option function
type PromMetricsOpt func(s *promMetricsImpl)

type promMetricsImpl struct {
	logger             logger.ILogger
	registerer         *prometheus.Registry
	pusher             simplego.MetricsPusher
	metricsHTTPHandler http.Handler
	counters           map[string]simplego.ICounter
	gauges             map[string]simplego.IGauge
	histograms         map[string]simplego.IHistogram
}

// NewPromMetrics creates a new prometheus metrics service
func NewPromMetrics(l logger.ILogger, cfg *PromMetricsConfig, opts ...PromMetricsOpt) (simplego.IMetrics, error) {
	if instance != nil {
		return instance, nil
	}
	s := &promMetricsImpl{
		registerer: prometheus.NewRegistry(),
		logger:     l,
	}

	for _, opt := range opts {
		opt(s)
	}
	if s.pusher == nil && s.metricsHTTPHandler == nil {
		return nil, simplego.ErrMissingMetricsPusher
	}
	if s.counters == nil {
		s.counters = map[string]simplego.ICounter{}
	}
	if s.histograms == nil {
		s.histograms = map[string]simplego.IHistogram{}
	}
	if s.gauges == nil {
		s.gauges = map[string]simplego.IGauge{}
	}
	if s.pusher != nil {
		if cfg.PushInterval == nil {
			s.logger.Log().Warn(nil, "simplego metrics: failed to start auto prom metrics pusher , missing interval from config")
		} else {
			go func() {
				if err := s.pusher.Start(*cfg.PushInterval); err != nil {
					s.logger.Log().Fatal(err, "simplego metrics: failed to start prom metrics pusher")
				}
			}()
		}
	}

	instance = s
	return s, nil
}

// Close cleans and closes service resources
func (s *promMetricsImpl) Close(ctx context.Context) error {
	if s.pusher != nil {
		return s.pusher.Close(ctx)
	}
	return nil
}

// RegisterCounter register a new counter by name, description and labels
func (s *promMetricsImpl) RegisterCounter(name string, description string, labels *[]string) error {
	if _, exists := s.counters[name]; exists {
		return simplego.ErrMetricExists
	}
	c := newPromCounter(name, description, labels)
	s.counters[name] = c
	return s.registerer.Register(c.underlying)
}

// RegisterHistogram register a new histogram by name, description, buckets and labels
func (s *promMetricsImpl) RegisterHistogram(name string, description string, buckets []float64, labels *[]string) error {
	if _, exists := s.histograms[name]; exists {
		return simplego.ErrMetricExists
	}
	h := newPromHistogram(name, description, buckets, labels)
	s.histograms[name] = h
	return s.registerer.Register(h.underlying)
}

// RegisterGauge register a new gauge by name, description and labels
func (s *promMetricsImpl) RegisterGauge(name string, description string, labels *[]string) error {
	if _, exists := s.gauges[name]; exists {
		return simplego.ErrMetricExists
	}
	g := newPromGauge(name, description, labels)
	s.gauges[name] = g
	return s.registerer.Register(g.underlying)
}

// GetMetricsHandler gets a metric http handler
func (s *promMetricsImpl) GetMetricsHandler(context.Context) (http.Handler, error) {
	if s.metricsHTTPHandler == nil {
		return nil, simplego.ErrMissingMetricsMiddlewareBuilder
	}
	return s.metricsHTTPHandler, nil
}

// PushCollectedMetrics manually push the current metrics
func (s *promMetricsImpl) PushCollectedMetrics(ctx context.Context) error {
	if s.pusher == nil {
		return simplego.ErrMissingMetricsPusher
	}
	return s.pusher.Push(ctx)
}

// GetCounter gets a counter metric by name
func (s *promMetricsImpl) GetCounter(name string) (simplego.ICounter, error) {
	counter, exists := s.counters[name]
	if !exists {
		return nil, simplego.ErrMissingMetric
	}
	return counter, nil
}

// GetGauge gets a gauge metric by name
func (s *promMetricsImpl) GetGauge(name string) (simplego.IGauge, error) {
	gauge, exists := s.gauges[name]
	if !exists {
		return nil, simplego.ErrMissingMetric
	}
	return gauge, nil
}

// GetHistogram gets a histogram metric by name
func (s *promMetricsImpl) GetHistogram(name string) (simplego.IHistogram, error) {
	histogram, exists := s.histograms[name]
	if !exists {
		return nil, simplego.ErrMissingMetric
	}
	return histogram, nil
}
