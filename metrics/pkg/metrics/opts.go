package metrics

func WithCounters(o map[string]ICounter) MetricsOpt {
	return func(s *metricsImpl) {
		s.counters = o
	}
}

func WithHistograms(o map[string]IHistogram) MetricsOpt {
	return func(s *metricsImpl) {
		s.histograms = o
	}
}

func WithGauges(o map[string]IGauge) MetricsOpt {
	return func(s *metricsImpl) {
		s.gauges = o
	}
}
