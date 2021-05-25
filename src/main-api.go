package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/geometry-labs/go-service-template/core"
	"github.com/geometry-labs/go-service-template/kafka"

	"github.com/geometry-labs/go-service-template/api/healthcheck"
	"github.com/geometry-labs/go-service-template/api/metrics"
	"github.com/geometry-labs/go-service-template/api/routes"
)

func main() {
	core.GetEnvironment()

	core.LoggingInit()

	// Start kafka consumers
	// Go routines start in function
	kafka.StartConsumers()

	// Start Prometheus client
	// Go routine starts in function
	metrics.Start()

	// Start API server
	// Go routine starts in function
	routes.Start()

	// Start Health server
	// Go routine starts in function
	healthcheck.Start()

	// Listen for close sig
	// Register for interupt (Ctrl+C) and SIGTERM (docker)
	shutdown := make(chan int)

	//create a notification channel to shutdown
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Info("Shutting down...")
		shutdown <- 1
	}()

	<-shutdown
}
