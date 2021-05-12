package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Environment struct {

	// Monitoring
	HealthPort    string `envconfig:"HEALTH_PORT" required:"false" default:"8080"`
	MetricsPort   string `envconfig:"METRICS_PORT" required:"false" default:"9400"`
	HealthPrefix  string `envconfig:"HEALTH_PREFIX" required:"false" default:"/health"`
	MetricsPrefix string `envconfig:"METRICS_PREFIX" required:"false" default:"/metrics"`

	// Logging
	LogLevel    string `envconfig:"LOG_LEVEL" required:"false" default:"INFO"`
	LogToFile   bool   `envconfig:"LOG_TO_FILE" required:"false" default:"false"`
	NetworkName string `envconfig:"NETWORK_NAME" required:"false" default:"mainnet"`

	// Kafka
	KafkaBrokerURL string `envconfig:"KAFKA_BROKER_URL" required:"false" default:""`
	KafkaGroupID   string `envconfig:"KAFKA_GROUP_ID" required:"false" default:"worker-group-id"`
	ConsumerTopics string `envconfig:"CONSUMER_TOPICS" required:"false" default:"input-topic-1,input-topic-2"`
	ProducerTopics string `envconfig:"PRODUCER_TOPICS" required:"false" default:"output-topic-1,output-topic-2"`

	// Workers
	BlocksWorkerConsumerTopic string `envconfig:"BLOCKS_WORKER_CONSUMER_TOPIC" required:"false" default:"input-topic-1"`
	BlocksWorkerProducerTopic string `envconfig:"BLOCKS_WORKER_PRODUCER_TOPIC" required:"false" default:"output-topic-1"`
}

var Vars Environment

// Run once on main.go
func GetEnvironment() {
	//Get environment variable file
	env_file := os.Getenv("ENV_FILE")
	if env_file != "" {
		_ = godotenv.Load(env_file)
	} else {
		_ = godotenv.Load()
	}

	err := envconfig.Process("", &Vars)
	if err != nil {
		log.Fatalf("ERROR: envconfig - %s\n", err.Error())
	}
}
