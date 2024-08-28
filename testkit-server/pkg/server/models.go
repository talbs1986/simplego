package server

import "net/http"

type TestkitMiddleware struct {
	TriggerCount int
}

func (s *TestkitMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.TriggerCount++
}
