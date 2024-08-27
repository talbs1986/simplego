package server

import (
	"context"
	"errors"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/talbs1986/simplego/app/pkg/logger"
	simplego "github.com/talbs1986/simplego/server/pkg/server"
)

// ChiServerOpt defines the chi server option function
type ChiServerOpt func(s *chiServerImpl)

type chiServerImpl struct {
	logger      logger.ILogger
	middlewares []simplego.ServerMiddleware
	router      *chi.Mux
	srvr        *http.Server
}

// NewChiServer creates a new server implemented by chi
func NewChiServer(l logger.ILogger, cfg *simplego.ServerConfig, opts ...ChiServerOpt) (simplego.IServer, error) {
	s := &chiServerImpl{
		logger:      l,
		middlewares: []simplego.ServerMiddleware{},
		router:      chi.NewRouter(),
	}

	for _, opt := range opts {
		opt(s)
	}

	s.srvr = &http.Server{
		Handler:           s.router,
		Addr:              cfg.Addr,
		ReadTimeout:       cfg.ReadTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}

	return s, nil
}

// Close cleans and closes service resources
func (s *chiServerImpl) Close(ctx context.Context) error {
	s.logger.Log().Info("simplego chi server: stopping server")
	return s.srvr.Shutdown(ctx)
}

// UseMiddleware register middleware to be used next in line
func (s *chiServerImpl) UseMiddleware(m simplego.ServerMiddleware) error {
	if m == nil {
		return simplego.ErrMissingMiddleware
	}
	s.router.Use(m)
	s.middlewares = append(s.middlewares, m)
	return nil
}

// GetMiddlewares gets the current middlewares in order
func (s *chiServerImpl) GetMiddlewares() []simplego.ServerMiddleware {
	return s.middlewares
}

// RegisterRoute registers the handler to the full route and method
func (s *chiServerImpl) RegisterRoute(r simplego.ServerRoute) error {
	if len(r.Route) < 1 {
		return simplego.ErrInvalidRoute
	}
	switch r.Method {
	case http.MethodGet:
		s.router.Get(r.Route, r.Handler)
	case http.MethodPost:
		s.router.Post(r.Route, r.Handler)
	case http.MethodPut:
		s.router.Put(r.Route, r.Handler)
	case http.MethodDelete:
		s.router.Delete(r.Route, r.Handler)
	case http.MethodHead:
		s.router.Head(r.Route, r.Handler)
	default:
		return simplego.ErrInvalidMethod
	}
	return nil
}

// Start starts listening for requests
func (s *chiServerImpl) Start() error {
	go func() {
		if err := s.srvr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Log().Fatal(err, "simplego chi server: failed to start server")
		}
		s.logger.Log().Info("simplego chi server: server closed successfully")
	}()
	return nil
}
