package logging

import (
	"fmt"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/geometry-labs/api/config"
)

func TestInit(t *testing.T) {

	// Test file logging
	os.Setenv("LOG_TO_FILE", "true")
	config.GetEnvironment()
	Init()
	log.Info("File log")

	// Test levels
	os.Setenv("LOG_LEVEL", "Panic")
	config.GetEnvironment()
	Init()
	log.Info("Should not log")

	os.Setenv("LOG_LEVEL", "FATAL")
	config.GetEnvironment()
	Init()
	log.Info("Should not log")

	os.Setenv("LOG_LEVEL", "ERROR")
	config.GetEnvironment()
	Init()
	log.Info("Should not log")

	os.Setenv("LOG_LEVEL", "WARN")
	config.GetEnvironment()
	Init()
	log.Warn("Warning")

	os.Setenv("LOG_LEVEL", "INFO")
	config.GetEnvironment()
	Init()
	log.Info("Info")

	os.Setenv("LOG_LEVEL", "DEBUG")
	config.GetEnvironment()
	Init()
	log.Debug("Debug")

	os.Setenv("LOG_LEVEL", "TRACE")
	config.GetEnvironment()
	Init()
	log.Trace("Trace")
}
