package metrics

import (
	"context"
	"errors"
	"net/http"
)

type IMetrics interface {
	GetMetricsHandler(context.Context) (http.Handler, error)
	PushCollectedMetrics(context.Context) error
	GetCounter(string) (ICounter, error)
	GetGauge(string) (IGauge, error)
	GetHistogram(string) (IHistogram, error)
}

type ICounter interface {
	IncBy(float64, MetricLabels) error
	Inc(MetricLabels) error
}
type IGauge interface {
	Set(float64, MetricLabels) error
	IncBy(float64, MetricLabels) error
	Inc(MetricLabels) error
}
type IHistogram interface {
	Record(float64, MetricLabels) error
}

type MetricsMiddlewareBuilder interface {
	BuildServerHandler() (http.Handler, error)
}
type MetricsPusher interface {
	Push(context.Context) error
}
type MetricsRegisterer interface {
	RegisterCounter(string, string, *[]string) error
	RegisterHistogram(string, string, []float64, *[]string) error
	RegisterGauge(string, string, *[]string) error
}
type MetricLabels map[string]string

type MetricsOpt func(*metricsImpl)

type metricsImpl struct {
	registerer        MetricsRegisterer
	pusher            MetricsPusher
	middlewareBuilder MetricsMiddlewareBuilder
	counters          map[string]ICounter
	gauges            map[string]IGauge
	histograms        map[string]IHistogram
}

var ErrMissingMetric = errors.New("simplego metrics: metric is missing")
var ErrMissingMetricsMiddlewareBuilder = errors.New("simplego metrics: metrics service server middleware builder isnt configured")
var ErrMissingMetricsPusher = errors.New("simplego metrics: metrics service pusher isnt configured")

func NewMetrics(ctx context.Context, registerer MetricsRegisterer, opts ...MetricsOpt) (IMetrics, error) {
	s := &metricsImpl{registerer: registerer}
	for _, opt := range opts {
		opt(s)
	}
	if s.registerer == nil {
		return nil, errors.New("simplego metrics: metrics registerer cant be nil")
	}
	if s.pusher == nil && s.middlewareBuilder == nil {
		return nil, errors.New("simplego metrics: metrics pusher or server middleware builder cant be nil")
	}
	if s.counters == nil {
		s.counters = map[string]ICounter{}
	}
	if s.histograms == nil {
		s.histograms = map[string]IHistogram{}
	}
	if s.gauges == nil {
		s.gauges = map[string]IGauge{}
	}
	return s, nil
}

func (s *metricsImpl) GetMetricsHandler(context.Context) (http.Handler, error) {
	if s.middlewareBuilder == nil {
		return nil, ErrMissingMetricsMiddlewareBuilder
	}
	return s.middlewareBuilder.BuildServerHandler()
}
func (s *metricsImpl) PushCollectedMetrics(ctx context.Context) error {
	if s.pusher == nil {
		return ErrMissingMetricsPusher
	}
	return s.pusher.Push(ctx)
}
func (s *metricsImpl) GetCounter(name string) (ICounter, error) {
	counter, exists := s.counters[name]
	if !exists {
		return nil, ErrMissingMetric
	}
	return counter, nil

}
func (s *metricsImpl) GetGauge(name string) (IGauge, error) {
	gauge, exists := s.gauges[name]
	if !exists {
		return nil, ErrMissingMetric
	}
	return gauge, nil

}
func (s *metricsImpl) GetHistogram(name string) (IHistogram, error) {
	histogram, exists := s.histograms[name]
	if !exists {
		return nil, ErrMissingMetric
	}
	return histogram, nil

}
