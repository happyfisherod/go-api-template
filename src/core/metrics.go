package core

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Metrics map[string]prometheus.Counter

func MetricsApiStart() {
	Metrics = make(map[string]prometheus.Counter)

	createApiGauges()

	// Start server
	http.Handle(Vars.MetricsPrefix, promhttp.Handler())
	go http.ListenAndServe(":"+Vars.MetricsPort, nil)
	log.Println("Started Metrics:", Vars.MetricsPort)
}

func MetricsWorkerStart() {
	Metrics = make(map[string]prometheus.Counter)

	// Create gauges
	createWorkerGauges()

	// Start server
	http.Handle(Vars.MetricsPrefix, promhttp.Handler())
	go http.ListenAndServe(":"+Vars.MetricsPort, nil)
	log.Println("Started Metrics:", Vars.MetricsPort)
}

func createApiGauges() {
	Metrics["requests_amount"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "requests_amount",
		Help:        "amount of requests",
		ConstLabels: prometheus.Labels{"network_name": Vars.NetworkName},
	})
	Metrics["kafka_messages_consumed"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "kafka_messages_consumed",
		Help:        "amount of messageds from kafka consumed",
		ConstLabels: prometheus.Labels{"network_name": Vars.NetworkName},
	})
	Metrics["websockets_connected"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "websockets_connected",
		Help:        "amount of websockets that have connected to the server",
		ConstLabels: prometheus.Labels{"network_name": Vars.NetworkName},
	})
	Metrics["websockets_bytes_written"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "websockets_bytes_written",
		Help:        "amount of bytes written through websockets",
		ConstLabels: prometheus.Labels{"network_name": Vars.NetworkName},
	})
}

func createWorkerGauges() {
	Metrics["kafka_messages_consumed"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "kafka_messages_consumed",
		Help:        "amount of messages from kafka consumed",
		ConstLabels: prometheus.Labels{"network_name": Vars.NetworkName},
	})
	Metrics["kafka_messages_produced"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "kafka_messages_produced",
		Help:        "amount of messages from kafka produced",
		ConstLabels: prometheus.Labels{"network_name": Vars.NetworkName},
	})
}
