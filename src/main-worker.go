package main

import (
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/geometry-labs/go-service-template/core"
	"github.com/geometry-labs/go-service-template/kafka"

	"github.com/geometry-labs/go-service-template/worker/healthcheck"
	"github.com/geometry-labs/go-service-template/worker/loader"
	"github.com/geometry-labs/go-service-template/worker/transformers"
)

const VersionWorker = "v0.1.0"

func main() {

	core.GetEnvironment()

	core.StartLoggingInit()
	zap.S().Debug("Main: Starting logging with level ", core.Config.LogLevel)

	// Start Prometheus client
	core.MetricsWorkerStart()

	// Start Health server
	healthcheck.Start()

	// Start kafka consumer
	kafka.StartWorkerConsumers()

	// Start kafka Producer
	kafka.StartProducers()
	// Wait for Kafka
	time.Sleep(1 * time.Second)

	// Start Postgres loader
	loader.StartBlockRawsLoader()

	// Start transformers
	transformers.StartBlocksTransformer()

	// Listen for close sig
	// Register for interupt (Ctrl+C) and SIGTERM (docker)
	shutdown := make(chan int)

	//create a notification channel to shutdown
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		zap.S().Info("Shutting down...")
		shutdown <- 1
		core.GetGlobal().ShutdownChan <- 1
	}()

	<-shutdown
	<-core.GetGlobal().ShutdownChan
}
