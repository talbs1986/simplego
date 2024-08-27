package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/talbs1986/simplego/app/pkg/logger"
	simplego_metrics "github.com/talbs1986/simplego/metrics/pkg/metrics"
	simplego_server "github.com/talbs1986/simplego/server/pkg/server"
)

func NewMetricsMiddleware(service simplego_metrics.IMetrics, silentRouteValidator simplego_server.SilentRoutesValidator, routeExtractor ReqMetricsMiddlewareRouteExtractor, reqCounter simplego_metrics.ICounter, reqLatency simplego_metrics.IHistogram) simplego_server.ServerMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if silentRouteValidator(r.URL) {
				next.ServeHTTP(w, r)
				return
			}

			wrapped := &simplego_server.InternalServerResponse{ResponseWriter: w}
			start := time.Now()
			next.ServeHTTP(wrapped, r)

			duration := float64(time.Since(start).Milliseconds())
			status := strconv.Itoa(wrapped.Status)
			routePath := routeExtractor(r.URL)
			reqCounter.Inc(simplego_metrics.MetricLabels{"path": routePath, "method": r.Method, "status": status})
			reqLatency.Record(duration, simplego_metrics.MetricLabels{"path": routePath, "method": r.Method, "status": status})
		})
	}
}

// NewLoggerMiddleware builds a default logger server middleware for request logging
func NewLoggerMiddleware(log logger.ILogger, silentRouteValidator simplego_server.SilentRoutesValidator) simplego_server.ServerMiddleware {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			wrapped := &simplego_server.InternalServerResponse{ResponseWriter: w}
			if !silentRouteValidator(r.URL) {
				defer func() {
					t2 := time.Now()
					if rec := recover(); rec != nil {
						log.Log().Error(nil, "simplego chi server: recovering from handle request error")
						http.Error(wrapped, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					}

					// log end request
					fields := &logger.LogFields{
						"remote_ip":          r.RemoteAddr,
						"url":                r.URL.Path,
						"method":             r.Method,
						"user_agent":         r.Header.Get("User-Agent"),
						"latency_ms":         time.Since(t2).Milliseconds(),
						"req_content_length": r.Header.Get("Content-Length"),
						"res_status":         wrapped.Status,
						"res_bytes_written":  wrapped.BytesWritten,
					}
					log.Log().With(fields).Info("simplego chi server: finished handling request")
				}()
			}

			next.ServeHTTP(wrapped, r)
		}
		return http.HandlerFunc(fn)
	}
}
