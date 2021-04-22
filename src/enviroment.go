package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	Port        string `envconfig:"REST_API_PORT" required:"false" default:"8080"`
	Prefix      string `envconfig:"REST_API_PREFIX" required:"false" default:""`
	MetricsPort string `envconfig:"REST_API_METRICS_PORT" required:"false" default:"9400"`
	NetworkName string `envconfig:"REST_API_NETWORK_NAME" required:"false" default:"mainnet"`
}

func getEnvironment() Environment {
	_ = godotenv.Load()

	var env Environment
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatalf("ERROR: envconfig - %s\n", err.Error())
	}
	return env
}
