package server

import (
	"bytes"
	"context"
	"errors"
	"net/http"

	simplego "github.com/talbs1986/simplego/server/pkg/server"
)

// TestkitMetrics defines a testkit metric service
type TestkitServer struct {
	Middlewares []simplego.ServerMiddleware
	Routes      map[string]simplego.ServerRoute
}

// NewTestkitServer creates a new testkit server
func NewTestkitServer() (simplego.IServer, error) {
	s := &TestkitServer{Middlewares: []simplego.ServerMiddleware{}, Routes: map[string]simplego.ServerRoute{}}
	return s, nil
}
func (s *TestkitServer) Close(ctx context.Context) error {
	*s = TestkitServer{Middlewares: []simplego.ServerMiddleware{}, Routes: map[string]simplego.ServerRoute{}}
	return nil
}

// UseMiddleware register middleware to be used next in line
func (s *TestkitServer) UseMiddleware(m simplego.ServerMiddleware) error {
	s.Middlewares = append(s.Middlewares, m)
	return nil
}

// GetMiddlewares gets the current registered middlewares in order
func (s *TestkitServer) GetMiddlewares() []simplego.ServerMiddleware {
	return s.Middlewares
}

// RegisterRoute registers the handler to the full route and method
func (s *TestkitServer) RegisterRoute(r simplego.ServerRoute) error {
	s.Routes[r.Method+"-"+r.Route] = r
	return nil
}

// Start starts listening for requests
func (s *TestkitServer) Start() error {
	return nil
}

// TestRoute runs the middlewares and then the route
func (s *TestkitServer) TestRoute(method string, path string, body []byte) error {
	testMiddleware := TestkitMiddleware{}
	for _, m := range s.Middlewares {
		m(testMiddleware)
	}
	route, exists := s.Routes[method+"-"+path]
	if !exists {
		return errors.New("simplego server testkit: failed to find route")
	}

	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		return errors.New("simplego server testkit: failed to create req")
	}
	route.Handler.ServeHTTP(&simplego.InternalServerResponse{ResponseWriter: &simplego.InternalServerResponse{}}, req)
	return nil
}
