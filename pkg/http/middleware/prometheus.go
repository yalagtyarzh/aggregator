package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func (m *Middleware) PrometheusMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.RequestURI != "/metrics" {
			requestsHandler := promhttp.InstrumentHandlerCounter(m.appPrometheus.ReqCounterVec.MustCurryWith(prometheus.Labels{"handler": request.URL.Path}), handler)
			latencyHandler := promhttp.InstrumentHandlerDuration(m.appPrometheus.LatencyHistogram.MustCurryWith(prometheus.Labels{"handler": request.URL.Path}), requestsHandler)
			inFlightHandler := promhttp.InstrumentHandlerInFlight(m.appPrometheus.ReqsInFlightGauge, latencyHandler)
			inFlightHandler.ServeHTTP(writer, request)
		} else {
			handler.ServeHTTP(writer, request)
		}
	})
}
