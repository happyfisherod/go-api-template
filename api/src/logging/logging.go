package logging

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/geometry-labs/api/config"
)

func Init() {
	if config.Vars.LogToFile == true {
		f, err := os.OpenFile("./api.log", os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			panic("Error opening log file: " + err.Error())
		}
		log.SetOutput(f)
	}

	switch config.Vars.LogLevel {
	case "PANIC", "panic", "Panic":
		log.SetLevel(log.PanicLevel)
		break
	case "FATAL", "fatal", "Fatal":
		log.SetLevel(log.FatalLevel)
		break
	case "ERROR", "error", "Error":
		log.SetLevel(log.ErrorLevel)
		break
	case "WARN", "warn", "Warn":
		log.SetLevel(log.WarnLevel)
		break
	case "INFO", "info", "Info":
		log.SetLevel(log.InfoLevel)
		break
	case "DEBUG", "debug", "Debug":
		log.SetLevel(log.DebugLevel)
		break
	case "TRACE", "trace", "Trace":
		log.SetLevel(log.TraceLevel)
		break
	default:
		panic("Error invalid log level: " + config.Vars.LogLevel)
	}
}
