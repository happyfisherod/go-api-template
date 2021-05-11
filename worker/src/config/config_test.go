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
		"HEALTH_PORT":      "health_port",
		"METRICS_PORT":     "metrics_port",
		"HEALTH_PREFIX":    "health_prefix",
		"METRICS_PREFIX":   "metrics_prefix",
		"LOG_LEVEL":        "log_level",
		"LOG_TO_FILE":      "true",
		"NETWORK_NAME":     "network_name",
		"KAFKA_BROKER_URL": "kafka_broker_url",
		"KAFKA_GROUP_ID":   "kafka_group_id",
		"INPUT_TOPICS":     "input_topics",
		"OUTPUT_TOPICS":    "output_topics",
	}

	for k, v := range env_map {
		os.Setenv(k, v)
	}

	// Load env
	GetEnvironment()

	// Check env
	assert.Equal(Vars.HealthPort, env_map["HEALTH_PORT"])
	assert.Equal(Vars.MetricsPort, env_map["METRICS_PORT"])
	assert.Equal(Vars.HealthPrefix, env_map["HEALTH_PREFIX"])
	assert.Equal(Vars.MetricsPrefix, env_map["METRICS_PREFIX"])
	assert.Equal(Vars.LogLevel, env_map["LOG_LEVEL"])
	assert.Equal(Vars.LogToFile, true)
	assert.Equal(Vars.NetworkName, env_map["NETWORK_NAME"])
	assert.Equal(Vars.KafkaBrokerURL, env_map["KAFKA_BROKER_URL"])
	assert.Equal(Vars.KafkaGroupID, env_map["KAFKA_GROUP_ID"])
	assert.Equal(Vars.InputTopics, env_map["INPUT_TOPICS"])
	assert.Equal(Vars.OutputTopics, env_map["OUTPUT_TOPICS"])
}
