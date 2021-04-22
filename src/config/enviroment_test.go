package config

import (
	"os"

	"testing"
)

func TestEnvironment(t *testing.T) {

	// Set env
	env_map := map[string]string{
		"PORT":         "port",
		"PREFIX":       "prefix",
		"METRICS_PORT": "metrics_port",
		"NETWORK_NAME": "network_name",
	}

	for k, v := range env_map {
		os.Setenv(k, v)
	}

	// Load env
	GetEnvironment()

	// Check env
	if Vars.Port != env_map["PORT"] {
		t.Errorf("Invalid value for env variable: PORT")
	}
	if Vars.Prefix != env_map["PREFIX"] {
		t.Errorf("Invalid value for env variable: PREFIX")
	}
	if Vars.MetricsPort != env_map["METRICS_PORT"] {
		t.Errorf("Invalid value for env variable: METRICS_PORT")
	}
	if Vars.NetworkName != env_map["NETWORK_NAME"] {
		t.Errorf("Invalid value for env variable: NETWORK_NAME")
	}
}
