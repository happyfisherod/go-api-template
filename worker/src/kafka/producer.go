package kafka

import (
	"strings"

	"github.com/geometry-labs/worker/config"
	"github.com/geometry-labs/worker/metrics"

	log "github.com/sirupsen/logrus"
	confluent "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KafkaTopicProducer struct {
	BrokerURL string
	TopicName string
	TopicChan chan *confluent.Message
}

// map[Topic_Name] -> Producer
var KafkaTopicProducers = map[string]*KafkaTopicProducer{}

func StartProducers() {
	kafka_broker := config.Vars.KafkaBrokerURL
	producer_topics := strings.Split(config.Vars.ProducerTopics, ",")

	log.Debug("Start Producer: kafka_broker=", kafka_broker, " producer_topics=", producer_topics)

	for _, t := range producer_topics {
		KafkaTopicProducers[t] = &KafkaTopicProducer{
			kafka_broker,
			t,
			make(chan *confluent.Message),
		}

		go KafkaTopicProducers[t].produceTopic()
	}
}

func (k *KafkaTopicProducer) produceTopic() {

	producer, err := confluent.NewProducer(&confluent.ConfigMap{
		"bootstrap.servers": k.BrokerURL,
	})

	if err != nil {
		log.Panic("KAFKA PRODUCER PANIC: ", err.Error())
	}
	defer producer.Close()

	log.Debug(k.TopicName, " Producer: started producing")

	for {
		topic_msg := <-k.TopicChan

		producer.Produce(&confluent.Message{
			TopicPartition: confluent.TopicPartition{Topic: &k.TopicName, Partition: confluent.PartitionAny},
			Key:            topic_msg.Key,
			Value:          topic_msg.Value,
		}, nil)

		log.Debug(k.TopicName, " Producer: producing message - ", string(topic_msg.Key))
		metrics.Metrics["kafka_messages_produced"].Inc()
	}
}
