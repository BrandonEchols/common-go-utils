package routing

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

/*
	Below we define Prometheus-Middleware Factory methods. The methods take a name (that must be unique) and generate
	a middleware function that can be added to a CustomRouter.
*/

/*
	This middleware will report metrics on the durations of requests
*/
func MakePrometheusAPIRequestsDuration(name string) func(http.HandlerFunc) http.HandlerFunc {
	histogramOpts := prometheus.HistogramOpts{
		Name:        name,
		Help:        "A histogram of latencies for requests.",
		Buckets:     []float64{.25, .5, 1, 2.5, 5, 10},
		ConstLabels: nil,
	}
	durationVec := prometheus.NewHistogramVec(
		histogramOpts,
		[]string{"code", "method"},
	)

	// Register all of the metrics in the standard registry.
	prometheus.MustRegister(durationVec)

	return func(next http.HandlerFunc) http.HandlerFunc {
		return promhttp.InstrumentHandlerDuration(durationVec, next)
	}
}

/*
	This middleware will report metrics on the total number of requests
*/
func MakePrometheusAPIRequestsCounter(name string, methods []string) func(http.HandlerFunc) http.HandlerFunc {
	api_counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: "A counter for requests to the wrapped handler.",
		},
		methods,
	)

	// Register all of the metrics in the standard registry.
	prometheus.MustRegister(api_counter)

	return func(next http.HandlerFunc) http.HandlerFunc {
		return promhttp.InstrumentHandlerCounter(api_counter, next)
	}
}

/*
	This middleware will report metrics on the number of requests currently being handled
*/
func MakePrometheusInFlightCounter(name string) func(http.HandlerFunc) http.HandlerFunc {
	inFlightGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: name,
		Help: "A gauge of requests currently being served by the wrapped handler.",
	})

	// Register all of the metrics in the standard registry.
	prometheus.MustRegister(inFlightGauge)

	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			inFlightGauge.Inc()
			defer inFlightGauge.Dec()
			next.ServeHTTP(w, r)
		})
	}
}
