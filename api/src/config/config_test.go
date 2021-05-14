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
		"KAFKA_GROUP_ID":          "kafka_group_id",
		"TOPIC_NAMES":             "[topic_names,topic_names,topic_names]",
		"SCHEMA_NAMES":            "[schema_names,schema_names,schema_names]",
	}

	for k, v := range env_map {
		os.Setenv(k, v)
	}

	// Load env
	GetEnvironment()

	// Check env
	assert.Equal(env_map["VERSION"], Vars.Version)
	assert.Equal(env_map["PORT"], Vars.Port)
	assert.Equal(env_map["HEALTH_PORT"], Vars.HealthPort)
	assert.Equal(env_map["METRICS_PORT"], Vars.MetricsPort)
	assert.Equal(env_map["REST_PREFIX"], Vars.RestPrefix)
	assert.Equal(env_map["WEBSOCKET_PREFIX"], Vars.WebsocketPrefix)
	assert.Equal(env_map["HEALTH_PREFIX"], Vars.HealthPrefix)
	assert.Equal(env_map["METRICS_PREFIX"], Vars.MetricsPrefix)
	assert.Equal(5, Vars.HealthPollingInterval)
	assert.Equal(env_map["LOG_LEVEL"], Vars.LogLevel)
	assert.Equal(true, Vars.LogToFile)
	assert.Equal(env_map["NETWORK_NAME"], Vars.NetworkName)
	assert.Equal(env_map["KAFKA_BROKER_URL"], Vars.KafkaBrokerURL)
	assert.Equal(env_map["KAFKA_GROUP_ID"], Vars.KafkaGroupID)
	assert.Equal(3, len(Vars.TopicNames))
	assert.Equal(3, len(Vars.SchemaNames))
}
