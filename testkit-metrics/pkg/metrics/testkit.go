package metrics

import (
	"context"
	"net/http"

	simplego "github.com/talbs1986/simplego/metrics/pkg/metrics"
)

// TestkitMetrics defines a testkit metric service
type TestkitMetrics struct {
	counters   map[string]simplego.ICounter
	gauges     map[string]simplego.IGauge
	histograms map[string]simplego.IHistogram
}

// NewTestkitMetrics creates a new testkit metric service
func NewTestkitMetrics() (simplego.IMetrics, error) {
	s := &TestkitMetrics{}
	return s, nil
}
func (s *TestkitMetrics) Close(ctx context.Context) error {
	*s = TestkitMetrics{}
	return nil
}
func (s *TestkitMetrics) BuildServerHandler() (http.Handler, error) {
	return nil, nil
}
func (s *TestkitMetrics) Push(ctx context.Context) error {
	return nil
}
func (s *TestkitMetrics) RegisterCounter(name string, description string, labels *[]string) error {
	if _, exists := s.counters[name]; exists {
		return simplego.ErrMetricExists
	}
	c := NewTestkitCounter(name, labels)
	s.counters[name] = c
	return nil
}
func (s *TestkitMetrics) RegisterHistogram(name string, description string, buckets []float64, labels *[]string) error {
	if _, exists := s.histograms[name]; exists {
		return simplego.ErrMetricExists
	}
	h := NewTestkitHistogram(name, buckets, labels)
	s.histograms[name] = h
	return nil
}
func (s *TestkitMetrics) RegisterGauge(name string, description string, labels *[]string) error {
	if _, exists := s.gauges[name]; exists {
		return simplego.ErrMetricExists
	}
	g := NewTestkitGauge(name, labels)
	s.counters[name] = g
	return nil
}
func (s *TestkitMetrics) GetMetricsHandler(context.Context) (http.Handler, error) {
	return nil, nil
}
func (s *TestkitMetrics) PushCollectedMetrics(ctx context.Context) error {
	return nil
}
func (s *TestkitMetrics) GetCounter(name string) (simplego.ICounter, error) {
	counter, exists := s.counters[name]
	if !exists {
		return nil, simplego.ErrMissingMetric
	}
	return counter, nil
}
func (s *TestkitMetrics) GetGauge(name string) (simplego.IGauge, error) {
	gauge, exists := s.gauges[name]
	if !exists {
		return nil, simplego.ErrMissingMetric
	}
	return gauge, nil
}
func (s *TestkitMetrics) GetHistogram(name string) (simplego.IHistogram, error) {
	histogram, exists := s.histograms[name]
	if !exists {
		return nil, simplego.ErrMissingMetric
	}
	return histogram, nil
}
