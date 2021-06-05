package metrics

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/geometry-labs/go-service-template/core"
)

var Metrics map[string]prometheus.Counter

// TODO Split API and WORKER starts
func Start() {
	Metrics = make(map[string]prometheus.Counter)

	// Create gauges
	Metrics["requests_amount"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "requests_amount",
		Help:        "amount of requests",
		ConstLabels: prometheus.Labels{"network_name": core.Vars.NetworkName},
	})
	Metrics["kafka_messages_consumed"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "kafka_messages_consumed",
		Help:        "amount of messageds from kafka consumed",
		ConstLabels: prometheus.Labels{"network_name": core.Vars.NetworkName},
	})
	Metrics["websockets_connected"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "websockets_connected",
		Help:        "amount of websockets that have connected to the server",
		ConstLabels: prometheus.Labels{"network_name": core.Vars.NetworkName},
	})
	Metrics["websockets_bytes_written"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "websockets_bytes_written",
		Help:        "amount of bytes written through websockets",
		ConstLabels: prometheus.Labels{"network_name": core.Vars.NetworkName},
	})

	// Start server
	http.Handle(core.Vars.MetricsPrefix, promhttp.Handler())
	go http.ListenAndServe(":"+core.Vars.MetricsPort, nil)
	zap.S().Info("Started Metrics:", core.Vars.MetricsPort)
}
