package metrics

import (
	"context"
	"errors"
	"net/http"
	"time"

	simplego "github.com/talbs1986/simplego/app/pkg/app"
)

// IMetrics defines the api for metrics
type IMetrics interface {
	// RegisterCounter register a new counter by name, description and labels
	RegisterCounter(string, string, *[]string) error
	// RegisterHistogram register a new histogram by name, description, buckets and labels
	RegisterHistogram(string, string, []float64, *[]string) error
	// RegisterGauge register a new gauge by name, description and labels
	RegisterGauge(string, string, *[]string) error
	// GetMetricsHandler gets a metric http handler
	GetMetricsHandler(context.Context) (http.Handler, error)
	// PushCollectedMetrics manually push the current metrics
	PushCollectedMetrics(context.Context) error
	// GetCounter gets a counter metric by name
	GetCounter(string) (ICounter, error)
	// GetGauge gets a gauge metric by name
	GetGauge(string) (IGauge, error)
	// GetHistogram gets a histogram metric by name
	GetHistogram(string) (IHistogram, error)
	// CloseableService clean and close resources api
	simplego.CloseableService
}

// ICounter defines an for a counter metric
type ICounter interface {
	// IncBy increments by the amount
	IncBy(float64, MetricLabels) error
	// Inc increments by 1
	Inc(MetricLabels) error
}

// ICounter defines an api for a gauge metric
type IGauge interface {
	// Set sets to the amount
	Set(float64, MetricLabels) error
	// IncBy increments by the amount
	IncBy(float64, MetricLabels) error
	// Inc increments by 1
	Inc(MetricLabels) error
}

// IHistogram defines an api for an histogram metric
type IHistogram interface {
	// Record records the amount
	Record(float64, MetricLabels) error
}

// MetricsMiddlewareBuilder defines the api for metric http middleware builder
type MetricsMiddlewareBuilder interface {
	// BuildServerHandler builds a metric http handler
	BuildServerHandler() (http.Handler, error)
}

// MetricsPusher defines the api for a metric push
type MetricsPusher interface {
	// Push pushes the collected metrics
	Push(context.Context) error
	// Start starts the automatic push by the interval
	Start(time.Duration) error
	// CloseableService clean and close resources api
	simplego.CloseableService
}

// MetricLabels defines the metric labels type
type MetricLabels map[string]string

// ErrMissingMetric missing metric error
var ErrMissingMetric = errors.New("simplego metrics: metric is missing")

// ErrMissingMetricsMiddlewareBuilder missing metric middleware builder error
var ErrMissingMetricsMiddlewareBuilder = errors.New("simplego metrics: metrics service server middleware builder isnt configured")

// ErrMissingMetricsPusher missing metric pusher error
var ErrMissingMetricsPusher = errors.New("simplego metrics: metrics service pusher isnt configured")

// ErrMetricExists metric already exists errors
var ErrMetricExists = errors.New("simplego metrics: metrics already registered")
