package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/geometry-labs/api/config"
	"github.com/geometry-labs/api/healthcheck"
	"github.com/geometry-labs/api/metrics"
	"github.com/geometry-labs/api/server"
)

func main() {
	config.GetEnvironment()

	// Start Prometheus client
	go metrics.StartPrometheusHttpServer(config.Vars.MetricsPort, config.Vars.NetworkName)

	// Start API server
	go server.Start()

	// Start Health server
	go healthcheck.Start()

	// Listen for close sig
	// Register for interupt (Ctrl+C) and SIGTERM (docker)
	shutdown := make(chan int)

	//create a notification channel to shutdown
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Shutting down...")
		shutdown <- 1
	}()

	<-shutdown
}
