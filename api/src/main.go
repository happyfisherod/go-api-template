package main

import (
	"github.com/geometry-labs/api/kafka"
	"github.com/geometry-labs/api/logging"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/geometry-labs/api/config"
	"github.com/geometry-labs/api/healthcheck"
	"github.com/geometry-labs/api/metrics"
	"github.com/geometry-labs/api/server"
)

func main() {
	config.GetEnvironment()

	logging.Init()

	// Start kafka consumers
	// Go routines start in function
	kafka.Start()
	//returnConsumedBlockChan := make(chan string)
	//defer close(returnConsumedBlockChan)
	//go service.StartConsumeBlocks(returnConsumedBlockChan)
	//go service.StartTransformBlocks(returnConsumedBlockChan)

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
