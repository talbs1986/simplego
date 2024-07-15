package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	simplego "github.com/talbs1986/simplego/metrics/pkg/metrics"
)

type promCounter struct {
	underlying *prometheus.CounterVec
}

func newPromCounter(name string, description string, labels *[]string) *promCounter {
	l := []string{}
	if labels != nil {
		l = *labels
	}
	c := &promCounter{
		underlying: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: name,
				Help: description,
			},
			l,
		),
	}
	return c
}

// IncBy increments by the amount
func (c *promCounter) IncBy(count float64, labels simplego.MetricLabels) error {
	c.underlying.With(prometheus.Labels(labels)).Add(count)
	return nil
}

// Inc increments by 1
func (c *promCounter) Inc(labels simplego.MetricLabels) error {
	c.underlying.With(prometheus.Labels(labels)).Inc()
	return nil
}

type promHistogram struct {
	underlying *prometheus.HistogramVec
}

func newPromHistogram(name string, description string, buckets []float64, labels *[]string) *promHistogram {
	l := []string{}
	if labels != nil {
		l = *labels
	}
	h := &promHistogram{
		underlying: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    name,
				Help:    description,
				Buckets: buckets,
			},
			l,
		),
	}
	return h
}

// Record records the amount
func (c *promHistogram) Record(count float64, labels simplego.MetricLabels) error {
	c.underlying.With(prometheus.Labels(labels)).Observe(count)
	return nil
}

type promGauge struct {
	underlying *prometheus.GaugeVec
}

func newPromGauge(name string, description string, labels *[]string) *promGauge {
	l := []string{}
	if labels != nil {
		l = *labels
	}
	g := &promGauge{
		underlying: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: name,
				Help: description,
			},
			l,
		),
	}
	return g
}

// Set sets to the amount
func (c *promGauge) Set(count float64, labels simplego.MetricLabels) error {
	c.underlying.With(prometheus.Labels(labels)).Set(count)
	return nil
}

// IncBy increments by the amount
func (c *promGauge) IncBy(count float64, labels simplego.MetricLabels) error {
	c.underlying.With(prometheus.Labels(labels)).Add(count)
	return nil
}

// Inc increments by 1
func (c *promGauge) Inc(labels simplego.MetricLabels) error {
	c.underlying.With(prometheus.Labels(labels)).Inc()
	return nil
}
