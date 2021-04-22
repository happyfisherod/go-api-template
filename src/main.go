package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/geometry-labs/rest-api/metrics"
	"github.com/geometry-labs/rest-api/rest"
)

func main() {
	env := getEnvironment()

	// Start Prometheus client
	go metrics.StartPrometheusHttpServer(env.MetricsPort, env.NetworkName)

	// Start REST Api server
	go rest.StartHttpServer(env.Port, env.Prefix)

	// Listen for close sig
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	// Keep main thread alive
	for sig := range sigCh {
		log.Printf("Stopping websocket server...%s", sig.String())
	}
}
