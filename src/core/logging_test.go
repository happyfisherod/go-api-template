package core

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestInit(t *testing.T) {

	// Test file logging
	os.Setenv("LOG_LEVEL", "Info")
	os.Setenv("LOG_TO_FILE", "true")
	GetEnvironment()
	LoggingInit()
	log.Info("File log")

	// Test levels
	os.Setenv("LOG_LEVEL", "Panic")
	GetEnvironment()
	LoggingInit()
	log.Info("Should not log")

	os.Setenv("LOG_LEVEL", "FATAL")
	GetEnvironment()
	LoggingInit()
	log.Info("Should not log")

	os.Setenv("LOG_LEVEL", "ERROR")
	GetEnvironment()
	LoggingInit()
	log.Info("Should not log")

	os.Setenv("LOG_LEVEL", "WARN")
	GetEnvironment()
	LoggingInit()
	log.Warn("Warning")

	os.Setenv("LOG_LEVEL", "INFO")
	GetEnvironment()
	LoggingInit()
	log.Info("Info")

	os.Setenv("LOG_LEVEL", "DEBUG")
	GetEnvironment()
	LoggingInit()
	log.Debug("Debug")

	os.Setenv("LOG_LEVEL", "TRACE")
	GetEnvironment()
	LoggingInit()
	log.Trace("Trace")
}
