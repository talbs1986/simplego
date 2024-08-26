package server

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/talbs1986/simplego/app/pkg/logger"
	simplego "github.com/talbs1986/simplego/server/pkg/server"
)

var (
	defaultServerMiddlewares = []simplego.ServerMiddleware{
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
	}
)

func BuildDefaultServeHealthRoutesValidator() simplego.HealthRouteValidator {
	return func(u *url.URL) bool {
		ustr := u.String()
		return strings.HasPrefix(ustr, "/health") || strings.HasPrefix(ustr, "/ready") || strings.HasPrefix(ustr, "/metrics")
	}
}
func BuildDefaultServeHealthRoutes() []simplego.ServerRoute {
	return []simplego.ServerRoute{
		{
			Method: http.MethodGet,
			Route:  "/health",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				return
			},
		},
		{
			Method: http.MethodGet,
			Route:  "/ready",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				return
			},
		},
	}
}

func BuildDefaultServerMiddlewares(reqTimeout time.Duration, logger logger.ILogger, userMiddlewares ...simplego.ServerMiddleware) []simplego.ServerMiddleware {
	tmp := []simplego.ServerMiddleware{}
	tmp = append(tmp, defaultServerMiddlewares...)
	tmp = append(tmp,
		middleware.Timeout(reqTimeout),
		NewLoggerMiddleware(logger, BuildDefaultServeHealthRoutesValidator()))
	if len(userMiddlewares) > 0 {
		tmp = append(tmp, userMiddlewares...)
	}
	return tmp
}
