package metrics

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/talbs1986/simplego/app/pkg/app"
)

type IMetrics interface {
	RegisterCounter(string, string, *[]string) error
	RegisterHistogram(string, string, []float64, *[]string) error
	RegisterGauge(string, string, *[]string) error
	GetMetricsHandler(context.Context) (http.Handler, error)
	PushCollectedMetrics(context.Context) error
	GetCounter(string) (ICounter, error)
	GetGauge(string) (IGauge, error)
	GetHistogram(string) (IHistogram, error)
	app.CloseableService
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
	Start(time.Duration) error
	app.CloseableService
}
type MetricsFetcher interface {
	RegisterCounter(string, string, *[]string) error
	RegisterHistogram(string, string, []float64, *[]string) error
	RegisterGauge(string, string, *[]string) error
}
type MetricLabels map[string]string

var ErrMissingMetric = errors.New("simplego metrics: metric is missing")
var ErrMissingMetricsMiddlewareBuilder = errors.New("simplego metrics: metrics service server middleware builder isnt configured")
var ErrMissingMetricsPusher = errors.New("simplego metrics: metrics service pusher isnt configured")
