package logging

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/geometry-labs/app/config"
)

func Init() {
	if config.Vars.LogToFile == true {
		f, err := os.OpenFile("./api.log", os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			panic("Error opening log file: " + err.Error())
		}
		log.SetOutput(f)
	}

	switch strings.ToUpper(config.Vars.LogLevel) {
	case "PANIC":
		log.SetLevel(log.PanicLevel)
		break
	case "FATAL":
		log.SetLevel(log.FatalLevel)
		break
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
		break
	case "WARN":
		log.SetLevel(log.WarnLevel)
		break
	case "INFO":
		log.SetLevel(log.InfoLevel)
		break
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
		break
	case "TRACE":
		log.SetLevel(log.TraceLevel)
		break
	default:
		panic("Error invalid log level: " + config.Vars.LogLevel)
	}
	log.Println("Log Level: " + log.GetLevel().String())
}
