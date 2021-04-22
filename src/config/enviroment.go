package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	Port        string `envconfig:"PORT" required:"false" default:"8080"`
	Prefix      string `envconfig:"PREFIX" required:"false" default:""`
	MetricsPort string `envconfig:"METRICS_PORT" required:"false" default:"9400"`
	NetworkName string `envconfig:"NETWORK_NAME" required:"false" default:"mainnet"`
}

var Vars Environment

// Run once on main.go
func GetEnvironment() {
	_ = godotenv.Load()
	err := envconfig.Process("", &Vars)
	if err != nil {
		log.Fatalf("ERROR: envconfig - %s\n", err.Error())
	}
}
