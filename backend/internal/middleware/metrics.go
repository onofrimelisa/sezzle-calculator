package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/sezzle/calculator-backend/internal/platform/metrics"
)

func Metrics(m *metrics.Metrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			sw := newStatusResponseWriter(w)
			next.ServeHTTP(sw, r)
			duration := time.Since(start).Seconds()
			status := strconv.Itoa(sw.statusCode)
			path := r.URL.Path
			m.RequestsTotal.WithLabelValues(path, status).Inc()
			m.RequestDuration.WithLabelValues(path).Observe(duration)
		})
	}
}
