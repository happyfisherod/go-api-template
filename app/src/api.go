package main

import (
	"github.com/geometry-labs/app/config"
	"github.com/geometry-labs/app/healthcheck"
	"github.com/geometry-labs/app/kafka"
	"github.com/geometry-labs/app/logging"
	"github.com/geometry-labs/app/metrics"
	"github.com/geometry-labs/app/server"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	//"github.com/geometry-labs/api/healthcheck"
	//"github.com/geometry-labs/api/kafka"
	//"github.com/geometry-labs/api/logging"
	//"github.com/geometry-labs/api/metrics"
	//"github.com/geometry-labs/api/server"
)

func main() {
	config.GetEnvironment()

	logging.Init()

	// Start kafka consumers
	// Go routines start in function
	kafka.Start()

	// Start Prometheus client
	// Go routine starts in function
	metrics.Start()

	// Start API server
	// Go routine starts in function
	server.Start()

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
