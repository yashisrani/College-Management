package middleware

import (
	"strconv"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Total number of requests processed by the MyApp web server.",
		},
		[]string{"path", "method"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "myapp_request_duration_seconds",
			Help:    "Duration of requests in seconds.",
			Buckets: prometheus.DefBuckets, // default buckets for histogram
		},
		[]string{"path", "method"},
	)

	StatusCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_response_status_total",
			Help: "Total number of responses by status code.",
		},
		[]string{"path", "method", "status_code"},
	)

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_requests_errors_total",
			Help: "Total number of error requests processed by the MyApp web server.",
		},
		[]string{"path", "status"},
	)
)

func PrometheusInit() {
	// Register all metrics with Prometheus
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(StatusCounter)
	prometheus.MustRegister(ErrorCount)
}

// TrackMetrics is the middleware function to track metrics for each request
func TrackMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		duration := time.Since(start).Seconds()
		path := c.FullPath() // FullPath gives you route pattern like /users/:id
		if path == "" {
			path = c.Request.URL.Path // fallback for non-matched routes
		}

		method := c.Request.Method
		status := c.Writer.Status()
		statusCodeStr := strconv.Itoa(status)

		// Increment request counter
		RequestCount.WithLabelValues(path, method).Inc()

		// Observe request duration
		RequestDuration.WithLabelValues(path, method).Observe(duration)

		// Increment status code counter
		StatusCounter.WithLabelValues(path, method, statusCodeStr).Inc()

		// If it's an error (>= 400), increment error count
		if status >= http.StatusBadRequest {
			ErrorCount.WithLabelValues(path, statusCodeStr).Inc()
		}
	}
}

// PrometheusHandler returns an HTTP handler for metrics endpoint
func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
