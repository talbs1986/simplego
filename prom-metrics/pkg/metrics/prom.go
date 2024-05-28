package metrics

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	simplego "github.com/talbs1986/simplego/metrics/pkg/metrics"
)

type PromMetricsOpt func(s *promMetricsImpl)

type promMetricsImpl struct {
	pushInterval *time.Duration
	registerer   *prometheus.Registry
}

func NewPromMetrics(opts ...PromMetricsOpt) (simplego.IMetrics, error) {
	instance := &promMetricsImpl{
		registerer: prometheus.NewRegistry(),
	}
	for _, opt := range opts {
		opt(instance)
	}
	if instance.pushInterval != nil {
		//TODO start ticker
	}

	return instance, nil
}

func (s *promMetricsImpl) BuildServerHandler() (http.Handler, error) {
	return promhttp.Handler(), nil
}
func (s *promMetricsImpl) Push(context.Context) error {
	if s.pushInterval == nil {
		return simplego.ErrMissingMetricsPusher
	}
	//TODO push gateway
	return nil
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

//TODO impl

// GetMetricsHandler(context.Context) (http.Handler, error)
// 	PushCollectedMetrics(context.Context) error
// 	GetCounter(string) (ICounter, error)
// 	GetGauge(string) (IGauge, error)
// 	GetHistogram(string) (IHistogram, error)
