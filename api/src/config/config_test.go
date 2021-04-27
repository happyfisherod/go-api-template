package config

import (
	"os"

	"testing"
)

func TestEnvironment(t *testing.T) {

	// Set env
	env_map := map[string]string{
		"PORT":             "port",
		"HEALTH_PORT":      "health_port",
		"REST_PREFIX":      "rest_prefix",
		"WEBSOCKET_PREFIX": "websocket_prefix",
		"METRICS_PORT":     "metrics_port",
		"LOG_LEVEL":        "log_level",
		"LOG_TO_FILE":      "true",
		"NETWORK_NAME":     "network_name",
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
	if Vars.HealthPort != env_map["HEALTH_PORT"] {
		t.Errorf("Invalid value for env variable: HEALTH_PORT")
	}
	if Vars.RestPrefix != env_map["REST_PREFIX"] {
		t.Errorf("Invalid value for env variable: REST_PREFIX")
	}
	if Vars.WebsocketPrefix != env_map["WEBSOCKET_PREFIX"] {
		t.Errorf("Invalid value for env variable: WEBSOCKET_PREFIX")
	}
	if Vars.MetricsPort != env_map["METRICS_PORT"] {
		t.Errorf("Invalid value for env variable: METRICS_PORT")
	}
	if Vars.LogLevel != env_map["LOG_LEVEL"] {
		t.Errorf("Invalid value for env variable: LOG_LEVEL")
	}
	if Vars.LogToFile != true {
		t.Errorf("Invalid value for env variable: LOG_TO_FILE")
	}
	if Vars.NetworkName != env_map["NETWORK_NAME"] {
		t.Errorf("Invalid value for env variable: NETWORK_NAME")
	}
}
