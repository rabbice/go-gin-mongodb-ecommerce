package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rabbice/ecommerce/logs"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logs.TotalRequests.WithLabelValues(
			c.Request.URL.Path).Inc()
		logs.TotalHTTPMethods.WithLabelValues(
			c.Request.Method).Inc()
		timer := prometheus.NewTimer(logs.HTTPDuration.WithLabelValues(c.Request.URL.Path))
		c.Next()
		timer.ObserveDuration()
	}

}
