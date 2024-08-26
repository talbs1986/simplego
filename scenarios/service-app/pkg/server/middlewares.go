package server

import (
	"context"
	"fmt"
	"net/http"

	simplego_metrics "github.com/talbs1986/simplego/metrics/pkg/metrics"
	simplego "github.com/talbs1986/simplego/server/pkg/server"
)

func NewMetricsMiddleware(service simplego_metrics.IMetrics) (simplego.ServerMiddleware, error) {
	metricsHandler, err := service.GetMetricsHandler(context.Background())
	if err != nil {
		return nil, fmt.Errorf("simplego service: failed to get metrics handler")
	}

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			metricsHandler.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}, nil
}
