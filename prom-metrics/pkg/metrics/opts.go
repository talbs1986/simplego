package metrics

import "github.com/talbs1986/simplego/metrics/pkg/metrics"

func WithCounters(o map[string]metrics.ICounter) PromMetricsOpt {
	return func(s *promMetricsImpl) {
		s.counters = o
	}
}

func WithHistograms(o map[string]metrics.IHistogram) PromMetricsOpt {
	return func(s *promMetricsImpl) {
		s.histograms = o
	}
}

func WithGauges(o map[string]metrics.IGauge) PromMetricsOpt {
	return func(s *promMetricsImpl) {
		s.gauges = o
	}
}

func WithPusher(p metrics.MetricsPusher) PromMetricsOpt {
	return func(s *promMetricsImpl) {
		s.pusher = p
	}
}
