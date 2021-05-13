package kafka

import (
	"os"
	"testing"
	"time"

	"github.com/geometry-labs/worker/config"
	"github.com/geometry-labs/worker/metrics"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

func init() {
	os.Setenv("LOG_LEVEL", "DEBUG")

	config.GetEnvironment()
	metrics.Start()
}

func TestKafkaTopicConsumer(t *testing.T) {
	topic_name := "mock-topic"

	// Mock broker
	mock_broker_id := int32(1)
	mock_broker := sarama.NewMockBroker(t, mock_broker_id)
	defer mock_broker.Close()

	mock_broker.SetHandlerByMap(map[string]sarama.MockResponse{
		"MetadataRequest": sarama.NewMockMetadataResponse(t).
			SetBroker(mock_broker.Addr(), mock_broker.BrokerID()).
			SetLeader(topic_name, 0, mock_broker.BrokerID()),
		"ProduceRequest": sarama.NewMockProduceResponse(t),
	})
}
