package server

import (
	"net/url"
)

type ReqMetricsMiddlewareRouteExtractor = func(*url.URL) string
