package kafka

import (
	"strings"

	"github.com/geometry-labs/worker/config"
	"github.com/geometry-labs/worker/metrics"

	log "github.com/sirupsen/logrus"
	confluent "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func StartConsumers() {
	kafka_broker := config.Vars.KafkaBrokerURL
	consumer_topics := strings.Split(config.Vars.ConsumerTopics, ",")

	log.Debug("Start Consumer: kafka_broker=", kafka_broker, " consumer_topics=", consumer_topics)

	for _, t := range consumer_topics {
		// Broadcaster indexed in Broadcasters map
		// Starts go routine
		newBroadcaster(t)

		topic_consumer := &KafkaTopicConsumer{
			kafka_broker,
			t,
			Broadcasters[t],
		}

		log.Debug("Start Consumers: starting ", t, " consumer...")
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

	log.Debug(k.TopicName, " Consumer: started consuming")

	for {
		topic_msg, err := consumer.ReadMessage(-1)
		metrics.Metrics["kafka_messages_consumed"].Inc()

		if err == nil {
			log.Debug(k.TopicName, " Consumer: consuming message - ", string(topic_msg.Key))
			k.Broadcaster.InputChan <- topic_msg
		}
	}
}
