package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/geometry-labs/worker/config"
	"github.com/geometry-labs/worker/healthcheck"
	"github.com/geometry-labs/worker/kafka"
	"github.com/geometry-labs/worker/logging"
	"github.com/geometry-labs/worker/metrics"
)

func main() {
	config.GetEnvironment()

	logging.Init()

	// Start Prometheus client
	go metrics.StartPrometheusHttpServer(config.Vars.MetricsPort, config.Vars.NetworkName)

	// Start kafka consumer and broadcaster
	go kafka.StartConsumer()

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
