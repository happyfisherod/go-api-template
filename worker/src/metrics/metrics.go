package metrics

import (
	"net/http"

	"github.com/geometry-labs/worker/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Metrics map[string]prometheus.Counter = nil

func Start() {
	if Metrics != nil {
		return
	}

	Metrics = make(map[string]prometheus.Counter)

	// Create gauges
	Metrics["kafka_messages_produced"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "kafka_messages_produced",
		Help:        "amount of messages from kafka produced",
		ConstLabels: prometheus.Labels{"network_name": config.Vars.NetworkName},
	})
	Metrics["kafka_messages_consumed"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "kafka_messages_consumed",
		Help:        "amount of messages from kafka consumed",
		ConstLabels: prometheus.Labels{"network_name": config.Vars.NetworkName},
	})

	// Start server
	http.Handle(config.Vars.MetricsPrefix, promhttp.Handler())
	go http.ListenAndServe(":"+config.Vars.MetricsPort, nil)
}
