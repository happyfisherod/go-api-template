package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Metrics map[string]prometheus.Counter

func StartPrometheusHttpServer(metrics_port string, network_name string) {
	Metrics = make(map[string]prometheus.Counter)

	// Create gauges
	Metrics["requests_amount"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "requests_amount",
		Help:        "amount of requests",
		ConstLabels: prometheus.Labels{"network_name": network_name},
	})

	// Start server
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+metrics_port, nil)
}
