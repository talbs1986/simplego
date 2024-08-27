package server

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/talbs1986/simplego/app/pkg/logger"
	simplego_metrics "github.com/talbs1986/simplego/metrics/pkg/metrics"
	simplego_server "github.com/talbs1986/simplego/server/pkg/server"
)

var (
	defaultServerMiddlewares = []simplego_server.ServerMiddleware{
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
	}
)

const (
	defaultRouteExtractPathLvl = 3
)

func BuildDefaultSilentRoutesValidator() simplego_server.SilentRoutesValidator {
	return func(u *url.URL) bool {
		path := u.Path
		return strings.HasPrefix(path, "/health") || strings.HasPrefix(path, "/ready") || strings.HasPrefix(path, "/metrics")
	}
}
func BuildDefaultReqRouteExtractor(ignoreFromPathLevel int) ReqMetricsMiddlewareRouteExtractor {
	return func(u *url.URL) string {
		if ignoreFromPathLevel < 1 {
			return u.Path
		}
		runes := strings.Split(u.Path, "/")
		res := strings.Join(runes[:int(math.Min(float64(len(runes)), float64(ignoreFromPathLevel)))], "/")
		return res
	}
}
func BuildDefaultServerRoutes(metricsService simplego_metrics.IMetrics) []simplego_server.ServerRoute {
	metricsHandler, err := metricsService.GetMetricsHandler(context.Background())
	if err != nil {
		panic(fmt.Errorf("simplego service: failed to get metrics handle, due to: %w", err))
	}
	return []simplego_server.ServerRoute{
		{
			Method: http.MethodGet,
			Route:  "/health",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
		},
		{
			Method: http.MethodGet,
			Route:  "/ready",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
		},
		{
			Method: http.MethodGet,
			Route:  "/metrics",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				metricsHandler.ServeHTTP(w, r)
			},
		},
	}
}

func BuildDefaultServerMiddlewares(reqTimeout time.Duration, logger logger.ILogger, metricsService simplego_metrics.IMetrics, userMiddlewares ...simplego_server.ServerMiddleware) []simplego_server.ServerMiddleware {
	silentRouteValidator := BuildDefaultSilentRoutesValidator()
	if err := metricsService.RegisterHistogram("reqs_latency", "server requests latency histograms",
		[]float64{0.1, 0.5, 0.9}, &[]string{"path", "method", "status"}); err != nil {
		panic(fmt.Errorf("simplego service: failed to register reqs_latency metrics histogram, due to: %w", err))
	}
	if err := metricsService.RegisterCounter("reqs_counter", "server requests counter", &[]string{"path", "method", "status"}); err != nil {
		panic(fmt.Errorf("simplego service: failed to register reqs_counter metrics counter, due to: %w", err))
	}
	reqsLatencyHist, err := metricsService.GetHistogram("reqs_latency")
	if err != nil {
		panic(fmt.Errorf("simplego service: failed to get reqs_latency metrics histogram, due to: %w", err))
	}
	reqsCounter, err := metricsService.GetCounter("reqs_counter")
	if err != nil {
		panic(fmt.Errorf("simplego service: failed to get reqs_counter metrics counter, due to: %w", err))
	}
	tmp := []simplego_server.ServerMiddleware{}
	tmp = append(tmp, defaultServerMiddlewares...)
	tmp = append(tmp,
		middleware.Timeout(reqTimeout),
		NewLoggerMiddleware(logger, silentRouteValidator),
		NewMetricsMiddleware(metricsService, silentRouteValidator, BuildDefaultReqRouteExtractor(defaultRouteExtractPathLvl), reqsCounter, reqsLatencyHist))
	if len(userMiddlewares) > 0 {
		tmp = append(tmp, userMiddlewares...)
	}
	return tmp
}
