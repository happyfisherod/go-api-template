package kafka

import (
	"strings"

	"github.com/geometry-labs/worker/config"
	"github.com/geometry-labs/worker/metrics"

	log "github.com/sirupsen/logrus"
	confluent "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func StartConsumer() {
	kafka_broker := config.Vars.KafkaBrokerURL
	input_topics := strings.Split(config.Vars.InputTopics, ",")

	for _, t := range input_topics {
		// Broadcaster indexed in Broadcasters map
		newBroadcaster(t, make(chan *confluent.Message))

		topic_consumer := &KafkaTopicConsumer{
			kafka_broker,
			t,
			Broadcasters[t],
		}

		go topic_consumer.consumeTopic()
	}
}

type KafkaTopicConsumer struct {
	BrokerURL   string
	TopicName   string
	Broadcaster *TopicBroadcaster
}

func (k *KafkaTopicConsumer) consumeTopic() {

	consumer, err := confluent.NewConsumer(&confluent.ConfigMap{
		"bootstrap.servers": k.BrokerURL,
		"group.id":          "websocket-api-group",
		"auto.offset.reset": "latest",
	})

	if err != nil {
		log.Panic("KAFKA CONSUMER PANIC: ", err.Error())
	}
	defer consumer.Close()

	consumer.SubscribeTopics([]string{k.TopicName}, nil)

	for {
		topic_msg, err := consumer.ReadMessage(-1)
		metrics.Metrics["kafka_messages_consumed"].Inc()

		if err == nil {
			k.Broadcaster.InputChan <- topic_msg
		}
	}
}
