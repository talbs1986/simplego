package metrics

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/talbs1986/simplego/app/pkg/logger"
	"github.com/talbs1986/simplego/metrics/pkg/metrics"
	simplego "github.com/talbs1986/simplego/metrics/pkg/metrics"
)

var (
	testPushInterval     = time.Second * 10
	testHistogramBuckets = []float64{0.5, 0.8, 0.99}
)

func Test_PromMetricsCTORShouldBeSingleton(t *testing.T) {
	instance = nil
	expected, err := NewPromMetrics(logger.NewFMTLogger(nil), &PromMetricsConfig{PushInterval: &testPushInterval}, WithPusher(&mockPusher{}))
	assert.Nil(t, err)
	actual, err := NewPromMetrics(logger.NewFMTLogger(nil), &PromMetricsConfig{PushInterval: &testPushInterval}, WithPusher(&mockPusher{}))
	assert.Nil(t, err)
	assert.Equal(t, actual, expected)
}

func Test_NegativePromMetricsCTORValidations(t *testing.T) {
	instance = nil
	_, err := NewPromMetrics(logger.NewFMTLogger(nil), &PromMetricsConfig{})
	assert.Error(t, err)
	assert.Equal(t, simplego.ErrMissingMetricsPusher, err)
}

func Test_RegisterMetricsAsSingleton(t *testing.T) {
	instance = nil
	service, err := NewPromMetrics(logger.NewFMTLogger(nil), &PromMetricsConfig{}, WithPusher(&mockPusher{}))
	assert.Nil(t, err)
	assert.Nil(t, service.RegisterCounter("test_counter", "test desc", nil))
	expectedCounter, err := service.GetCounter("test_counter")
	assert.Nil(t, err)
	assert.Equal(t, simplego.ErrMetricExists, service.RegisterCounter("test_counter", "test desc", nil))
	actualCounter, err := service.GetCounter("test_counter")
	assert.Nil(t, err)
	assert.Equal(t, expectedCounter, actualCounter)

	assert.Nil(t, service.RegisterGauge("test_gauge", "test desc", nil))
	expectedGauge, err := service.GetGauge("test_gauge")
	assert.Nil(t, err)
	assert.Equal(t, simplego.ErrMetricExists, service.RegisterGauge("test_gauge", "test desc", nil))
	actualGauge, err := service.GetGauge("test_gauge")
	assert.Nil(t, err)
	assert.Equal(t, expectedGauge, actualGauge)

	assert.Nil(t, service.RegisterHistogram("test_histo", "test desc", testHistogramBuckets, nil))
	expectedHisto, err := service.GetHistogram("test_histo")
	assert.Nil(t, err)
	assert.Equal(t, simplego.ErrMetricExists, service.RegisterHistogram("test_histo", "test desc", testHistogramBuckets, nil))
	actualHisto, err := service.GetHistogram("test_histo")
	assert.Nil(t, err)
	assert.Equal(t, expectedHisto, actualHisto)
}

func Test_NegativeGetHTTPMetricsHandlerMissingBuilder(t *testing.T) {
	instance = nil
	service, err := NewPromMetrics(logger.NewFMTLogger(nil), &PromMetricsConfig{}, WithPusher(&mockPusher{}))
	assert.Nil(t, err)
	_, err = service.GetMetricsHandler(context.Background())
	assert.Equal(t, metrics.ErrMissingMetricsMiddlewareBuilder, err)
}

func Test_NegativeGetMissingMetrics(t *testing.T) {
	instance = nil
	service, err := NewPromMetrics(logger.NewFMTLogger(nil), &PromMetricsConfig{}, WithPusher(&mockPusher{}))
	assert.Nil(t, err)
	_, err = service.GetCounter("test_counter")
	assert.Equal(t, simplego.ErrMissingMetric, err)

	_, err = service.GetGauge("test_gauge")
	assert.Equal(t, simplego.ErrMissingMetric, err)

	_, err = service.GetHistogram("test_histo")
	assert.Equal(t, simplego.ErrMissingMetric, err)
}
