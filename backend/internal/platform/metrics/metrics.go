package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	RequestsTotal   *prometheus.CounterVec
	RequestDuration *prometheus.HistogramVec
	ErrorsTotal     *prometheus.CounterVec
}

func New() *Metrics {
	return &Metrics{
		RequestsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "calculator_requests_total",
				Help: "Total number of calculator requests",
			},
			[]string{"operation", "status"},
		),
		RequestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "calculator_request_duration_seconds",
				Help:    "Duration of calculator requests in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"operation"},
		),
		ErrorsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "calculator_errors_total",
				Help: "Total number of calculator errors",
			},
			[]string{"operation", "error_code"},
		),
	}
}
