package core

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func init() {
	//core.GetEnvironment()
	Vars.ConfigFile = "config.api.test"
	Vars.ConfigType = "yaml"
	Vars.ConfigPath = "../../envfiles"
	ConfigInit()
}

func TestInit(t *testing.T) {

	// Test file logging
	os.Setenv("LOG_LEVEL", "Info")
	os.Setenv("LOG_TO_FILE", "true")
	GetEnvironment()
	StartLoggingInit()
	log.Info("File log")

	// Test levels
	os.Setenv("LOG_LEVEL", "Panic")
	GetEnvironment()
	StartLoggingInit()
	log.Info("Should not log")

	os.Setenv("LOG_LEVEL", "FATAL")
	GetEnvironment()
	StartLoggingInit()
	log.Info("Should not log")

	os.Setenv("LOG_LEVEL", "ERROR")
	GetEnvironment()
	StartLoggingInit()
	log.Info("Should not log")

	os.Setenv("LOG_LEVEL", "WARN")
	GetEnvironment()
	StartLoggingInit()
	log.Warn("Warning")

	os.Setenv("LOG_LEVEL", "INFO")
	GetEnvironment()
	StartLoggingInit()
	log.Info("Info")

	os.Setenv("LOG_LEVEL", "DEBUG")
	GetEnvironment()
	StartLoggingInit()
	log.Debug("Debug")

	os.Setenv("LOG_LEVEL", "TRACE")
	GetEnvironment()
	StartLoggingInit()
	log.Trace("Trace")
}
