package logs

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var TotalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of incoming requests",
	},
	[]string{"path"},
)

var TotalHTTPMethods = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_methods_total",
		Help: "Number of requests per HTTP method",
	},
	[]string{"method"},
)

var HTTPDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_in_seconds",
		Help: "Duration of HTTP requests",
	},
	[]string{"path"},
)