package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strings"
)

type Prometheus struct {
	ReqCounterVec     *prometheus.CounterVec
	ReqsInFlightGauge prometheus.Gauge
	LatencyHistogram  *prometheus.HistogramVec
}

func New(namespace, prefix, appName string) *Prometheus {
	namespace = strings.Replace(namespace, "-", "_", -1)
	appName = strings.Replace(appName, "-", "_", -1)

	p := &Prometheus{}
	subsystem := prefix + "_" + appName

	p.ReqCounterVec = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "requests_total",
			Help:      "Total HTTP requests processed",
		},
		[]string{"method", "code", "handler"},
	)

	p.ReqsInFlightGauge = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "pending_requests",
			Help:      "Requests at this moment",
		},
	)

	p.LatencyHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "request_duration_milliseconds",
			Help:      "Request processing duration time ms",
			Buckets:   []float64{0.0001, 0.001, 0.01, 0.1, 0.5, 1, 1.5, 2},
		},
		[]string{"method", "code", "handler"},
	)

	return p
}
