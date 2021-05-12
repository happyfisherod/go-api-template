package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/geometry-labs/worker/config"
	"github.com/geometry-labs/worker/healthcheck"
	"github.com/geometry-labs/worker/kafka"
	"github.com/geometry-labs/worker/logging"
	"github.com/geometry-labs/worker/metrics"
	"github.com/geometry-labs/worker/workers"

	log "github.com/sirupsen/logrus"
)

func main() {
	config.GetEnvironment()

	logging.Init()
	log.Debug("Main: Starting logging with level ", config.Vars.LogLevel)

	// Start Prometheus client
	metrics.Start()

	// Start Health server
	healthcheck.Start()

	// Start kafka consumer
	kafka.StartConsumers()

	// Start kafka consumer
	kafka.StartProducers()

	// Start workers
	workers.StartBlocksWorker()

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
