package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/talbs1986/simplego/app/pkg/logger"
	simplego "github.com/talbs1986/simplego/server/pkg/server"
)

// NewLoggerMiddleware builds a default logger server middleware for request logging
func NewLoggerMiddleware(log logger.ILogger, healthRouteValidator simplego.HealthRouteValidator) simplego.ServerMiddleware {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			defer func() {
				t2 := time.Now()
				if rec := recover(); rec != nil {
					log.Log().Error(nil, "simplego chi server: recovering from handle request error")
					http.Error(ww, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}

				// log end request
				fields := &logger.LogFields{
					"remote_ip":          r.RemoteAddr,
					"url":                r.URL.Path,
					"method":             r.Method,
					"user_agent":         r.Header.Get("User-Agent"),
					"latency_ms":         time.Since(t2).Milliseconds(),
					"req_content_length": r.Header.Get("Content-Length"),
					"res_status":         ww.Status(),
					"res_bytes_written":  ww.BytesWritten(),
				}
				if healthRouteValidator(r.URL) {
					return
				}
				log.Log().With(fields).Info("simplego chi server: finished handling request")
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
