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
		"HEALTH_PORT":                  "health_port",
		"METRICS_PORT":                 "metrics_port",
		"HEALTH_PREFIX":                "health_prefix",
		"METRICS_PREFIX":               "metrics_prefix",
		"LOG_LEVEL":                    "log_level",
		"LOG_TO_FILE":                  "true",
		"NETWORK_NAME":                 "network_name",
		"KAFKA_BROKER_URL":             "kafka_broker_url",
		"KAFKA_GROUP_ID":               "kafka_group_id",
		"CONSUMER_TOPICS":              "[consumer_topics,consumer_topics,consumer_topics]",
		"PRODUCER_TOPICS":              "[producer_topics,producer_topics,producer_topics]",
		"BLOCKS_WORKER_CONSUMER_TOPIC": "blocks_worker_consumer_topic",
		"BLOCKS_WORKER_PRODUCER_TOPIC": "blocks_worker_producer_topic",
	}

	for k, v := range env_map {
		os.Setenv(k, v)
	}

	// Load env
	GetEnvironment()

	// Check env
	assert.Equal(env_map["HEALTH_PORT"], Vars.HealthPort)
	assert.Equal(env_map["METRICS_PORT"], Vars.MetricsPort)
	assert.Equal(env_map["HEALTH_PREFIX"], Vars.HealthPrefix)
	assert.Equal(env_map["METRICS_PREFIX"], Vars.MetricsPrefix)
	assert.Equal(env_map["LOG_LEVEL"], Vars.LogLevel)
	assert.Equal(true, Vars.LogToFile)
	assert.Equal(env_map["NETWORK_NAME"], Vars.NetworkName)
	assert.Equal(env_map["KAFKA_BROKER_URL"], Vars.KafkaBrokerURL)
	assert.Equal(env_map["KAFKA_GROUP_ID"], Vars.KafkaGroupID)
	assert.Equal(3, len(Vars.ConsumerTopics))
	assert.Equal(3, len(Vars.ProducerTopics))
	assert.Equal(env_map["BLOCKS_WORKER_CONSUMER_TOPIC"], Vars.BlocksWorkerConsumerTopic)
	assert.Equal(env_map["BLOCKS_WORKER_PRODUCER_TOPIC"], Vars.BlocksWorkerProducerTopic)
}
