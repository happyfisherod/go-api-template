package main

import (
	"os"

	"testing"
)

func TestEnvironment(t *testing.T) {

	// Set env
	env_map := map[string]string{
		"KAFKA_WEBSOCKET_API_TOPICS":       "topics",
		"KAFKA_WEBSOCKET_API_BROKER_URL":   "broker_url_env",
		"KAFKA_WEBSOCKET_API_PORT":         "port_env",
		"KAFKA_WEBSOCKET_API_PREFIX":       "prefix_env",
		"KAFKA_WEBSOCKET_API_HEALTH_PORT":  "health_port_env",
		"KAFKA_WEBSOCKET_API_METRICS_PORT": "metrics_port_env",
		"KAFKA_WEBSOCKET_API_NETWORK_NAME": "network_name",
	}

	for k, v := range env_map {
		os.Setenv(k, v)
	}

	// Check env
	env := getEnvironment()

	if env.Topics != env_map["KAFKA_WEBSOCKET_API_TOPICS"] {
		t.Errorf("Invalid value for env variable: KAFKA_WEBSOCKET_API_TOPICS")
	}
	if env.BrokerURL != env_map["KAFKA_WEBSOCKET_API_BROKER_URL"] {
		t.Errorf("Invalid value for env variable: KAFKA_WEBSOCKET_API_BROKER_URL")
	}
	if env.Port != env_map["KAFKA_WEBSOCKET_API_PORT"] {
		t.Errorf("Invalid value for env variable: KAFKA_WEBSOCKET_API_PORT")
	}
	if env.Prefix != env_map["KAFKA_WEBSOCKET_API_PREFIX"] {
		t.Errorf("Invalid value for env variable: KAFKA_WEBSOCKET_API_PREFIX")
	}
	if env.HealthPort != env_map["KAFKA_WEBSOCKET_API_HEALTH_PORT"] {
		t.Errorf("Invalid value for env variable: KAFKA_WEBSOCKET_API_HEALTH_PORT")
	}
}
