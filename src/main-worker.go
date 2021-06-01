package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/geometry-labs/go-service-template/core"
	"github.com/geometry-labs/go-service-template/kafka"

	"github.com/geometry-labs/go-service-template/worker/healthcheck"
	"github.com/geometry-labs/go-service-template/worker/transformers"
)

func main() {
	os.Setenv("CONSUMER_TOPICS", "blocks")
	os.Setenv("PRODUCER_TOPICS", "blocks-ws")

	core.GetEnvironment()

	core.LoggingInit()
	log.Debug("Main: Starting logging with level ", core.Vars.LogLevel)

	// Start Prometheus client
	core.MetricsWorkerStart()

	// Start Health server
	healthcheck.Start()

	//// Start kafka consumer
	kafka.StartWorkerConsumers()

	//// Start kafka consumer
	kafka.StartProducers()

	time.Sleep(1 * time.Second)

	//// Start transformers
	transformers.StartBlocksTransformer()

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
