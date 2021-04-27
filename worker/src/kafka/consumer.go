package kafka

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	"github.com/geometry-labs/worker/metrics"
)

func StartConsumer() {
	// TODO
}

type KafkaTopicConsumer struct {
	TopicName string
	TopicChan chan *kafka.Message

	BrokerURL string
}

func (k *KafkaTopicConsumer) consumeAndBroadcastTopics() {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": k.BrokerURL,
		"group.id":          "websocket-api-group",
		"auto.offset.reset": "latest",
	})

	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	consumer.SubscribeTopics([]string{k.TopicName}, nil)

	newBroadcaster(k.TopicName, k.TopicChan)

	for {
		msg, err := consumer.ReadMessage(-1)
		metrics.Metrics["kafka_messages_consumed"].Inc()

		if err == nil {

			// NOTE: use select statement for non-blocking channels
			select {
			case k.TopicChan <- msg:
			default:
			}
		}
	}
}
