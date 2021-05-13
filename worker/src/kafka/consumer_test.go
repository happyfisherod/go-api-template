package kafka

import (
	"os"
	"testing"
	"time"

	"github.com/geometry-labs/worker/config"
	"github.com/geometry-labs/worker/metrics"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("LOG_LEVEL", "DEBUG")

	config.GetEnvironment()
	metrics.Start()
}

// ISSUE: only passes sometimes?
// Consumer does not read messages everytime (more times not)
func TestKafkaTopicConsumer(t *testing.T) {
	assert := assert.New(t)

	topic_name := "mock-topic"

	// Mock broker
	mock_broker_id := int32(1)
	mock_broker := sarama.NewMockBroker(t, mock_broker_id)
	defer mock_broker.Close()

	mock_broker.SetHandlerByMap(map[string]sarama.MockResponse{
		"MetadataRequest": sarama.NewMockMetadataResponse(t).
			SetBroker(mock_broker.Addr(), mock_broker.BrokerID()).
			SetLeader(topic_name, 0, mock_broker.BrokerID()),
		"OffsetRequest": sarama.NewMockOffsetResponse(t).
			SetOffset(topic_name, 0, sarama.OffsetOldest, 0).
			SetOffset(topic_name, 0, sarama.OffsetNewest, 2),
		"FetchRequest": sarama.NewMockFetchResponse(t, 1).
			SetMessage(topic_name, 0, 0, sarama.ByteEncoder([]byte{0x41, 0x42})).
			SetMessage(topic_name, 0, 1, sarama.ByteEncoder([]byte{0x41, 0x43})).
			SetMessage(topic_name, 0, 2, sarama.ByteEncoder([]byte{0x41, 0x44})),
	})

	mock_broker.SetNotifier(func(bytes_read int, bytes_written int) {
		log.Debug("MOCK NOTIFIER: bytes_read=", bytes_read, " bytes_written=", bytes_written)
	})

	// Start broadcaster
	newBroadcaster(topic_name)
	log.Info("TestKafkaTopicConsumer: Starting broadcaster...")
	go Broadcasters[topic_name].Start()

	// Start consumer
	topic_consumer := &KafkaTopicConsumer{
		mock_broker.Addr(),
		topic_name,
		Broadcasters[topic_name],
	}
	log.Info("TestKafkaTopicConsumer: Starting consumer...")
	go topic_consumer.consumeTopic()

	// Add broadcasting channeld
	topic_chan := make(chan *sarama.ConsumerMessage)
	broadcaster_output_chan_id := Broadcasters[topic_name].AddWorkerChannel(topic_chan)
	defer func() {
		Broadcasters[topic_name].RemoveWorkerChannel(broadcaster_output_chan_id)
	}()

	for {
		select {
		case topic_msg := <-topic_chan:
			// Pass
			assert.NotEqual(len(topic_msg.Value), 0)
			return
		case <-time.After(10 * time.Second):
			t.Fatal("No messages from consumer")
		}
	}
}
