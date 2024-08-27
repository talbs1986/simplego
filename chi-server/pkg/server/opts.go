package server

import (
	simplego "github.com/talbs1986/simplego/server/pkg/server"
)

// WithMiddlewares uses the provided middleware
func WithMiddlewares(ms []simplego.ServerMiddleware) ChiServerOpt {
	return func(s *chiServerImpl) {
		for _, m := range ms {
			middleware := m
			if err := s.UseMiddleware(middleware); err != nil {
				s.logger.Log().Fatal(err, "simplego chi server: failed to set app middlewares")
			}
		}
	}
}

// WithRoutes registers the provided middleware
func WithRoutes(rs []simplego.ServerRoute) ChiServerOpt {
	return func(s *chiServerImpl) {
		for _, r := range rs {
			route := r
			if err := s.RegisterRoute(route); err != nil {
				s.logger.Log().Fatal(err, "simplego chi server: failed to set app middlewares")
			}
		}
	}
}
