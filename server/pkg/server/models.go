package server

import (
	"errors"
	"net/http"
	"net/url"

	simplego "github.com/talbs1986/simplego/app/pkg/app"
)

// IMetrics defines the api for metrics
type IServer interface {
	// UseMiddleware register middleware to be used next in line
	UseMiddleware(ServerMiddleware) error
	// GetMiddlewares gets the current registered middlewares in order
	GetMiddlewares() []ServerMiddleware
	// RegisterRoute registers the handler to the full route and method
	RegisterRoute(ServerRoute) error
	// Start starts listening for requests
	Start() error
	// CloseableService clean and close resources api
	simplego.CloseableService
}

// ServerRoute defines the server route object
type ServerRoute struct {
	Method  string
	Route   string
	Handler ServerRouteHandle
}

type ServerMiddleware = func(http.Handler) http.Handler
type ServerRouteHandle = http.HandlerFunc
type HealthRouteValidator = func(*url.URL) bool

// ErrInvalidMethod method is invalid error
var ErrInvalidMethod = errors.New("simplego server: method is invalid")

// ErrInvalidRoute route is invalid error
var ErrInvalidRoute = errors.New("simplego server: route is invalid")

// ErrMissingMiddleware missing middleware error
var ErrMissingMiddleware = errors.New("simplego server: server middleware is missing")
