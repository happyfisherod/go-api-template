package config

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {
	assert := assert.New(t)

	// Set env
	env_map := map[string]string{
		"VERSION":                 "version",
		"PORT":                    "port",
		"HEALTH_PORT":             "health_port",
		"METRICS_PORT":            "metrics_port",
		"REST_PREFIX":             "rest_prefix",
		"WEBSOCKET_PREFIX":        "websocket_prefix",
		"HEALTH_PREFIX":           "health_prefix",
		"METRICS_PREFIX":          "metrics_prefix",
		"HEALTH_POLLING_INTERVAL": "5",
		"LOG_LEVEL":               "log_level",
		"LOG_TO_FILE":             "true",
		"NETWORK_NAME":            "network_name",
		"KAFKA_BROKER_URL":        "kafka_broker_url",
		"TOPIC_NAMES":             "topic_names",
	}

	for k, v := range env_map {
		os.Setenv(k, v)
	}

	// Load env
	GetEnvironment()

	// Check env
	assert.Equal(Vars.Version, env_map["VERSION"])
	assert.Equal(Vars.Port, env_map["PORT"])
	assert.Equal(Vars.HealthPort, env_map["HEALTH_PORT"])
	assert.Equal(Vars.MetricsPort, env_map["METRICS_PORT"])
	assert.Equal(Vars.RestPrefix, env_map["REST_PREFIX"])
	assert.Equal(Vars.WebsocketPrefix, env_map["WEBSOCKET_PREFIX"])
	assert.Equal(Vars.HealthPrefix, env_map["HEALTH_PREFIX"])
	assert.Equal(Vars.MetricsPrefix, env_map["METRICS_PREFIX"])
	assert.Equal(Vars.HealthPollingInterval, 5)
	assert.Equal(Vars.LogLevel, env_map["LOG_LEVEL"])
	assert.Equal(Vars.LogToFile, true)
	assert.Equal(Vars.NetworkName, env_map["NETWORK_NAME"])
	assert.Equal(Vars.KafkaBrokerURL, env_map["KAFKA_BROKER_URL"])
	assert.Equal(Vars.TopicNames, env_map["TOPIC_NAMES"])
}
