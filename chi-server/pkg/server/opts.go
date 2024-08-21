package server

import (
	"net/http"

	simplego "github.com/talbs1986/simplego/server/pkg/server"
)

// WithMiddlewares uses the provided middleware
func WithMiddlewares(ms []simplego.ServerMiddleware) ChiServerOpt {
	return func(s *chiServerImpl) {
		for _, m := range ms {
			s.UseMiddleware(m)
		}
	}
}

// WithRoutes registeres the provided middleware
func WithRoutes(rs []simplego.ServerRoute) ChiServerOpt {
	return func(s *chiServerImpl) {
		for _, r := range rs {
			s.RegisterRoute(r)
		}
	}
}

// WithHTTPServer sets the http server
func WithHTTPServer(h *http.Server) ChiServerOpt {
	return func(s *chiServerImpl) {
		s.srvr = h
	}
}
