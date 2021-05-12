package kafka

import (
	"strings"

	"github.com/geometry-labs/worker/config"
	"github.com/geometry-labs/worker/metrics"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

type KafkaTopicProducer struct {
	BrokerURL string
	TopicName string
	TopicChan chan *sarama.ProducerMessage
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
			make(chan *sarama.ProducerMessage),
		}

		go KafkaTopicProducers[t].produceTopic()
	}
}

func (k *KafkaTopicProducer) produceTopic() {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{k.BrokerURL}, config)
	if err != nil {
		log.Panic("KAFKA PRODUCER NEWSYNCPRODUCER PANIC: ", err.Error())
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Panic("KAFKA PRODUCER CLOSE PANIC: ", err.Error())
		}
	}()

	log.Debug("Producer ", k.TopicName, ": Started producing")
	for {
		topic_msg := <-k.TopicChan

		partition, offset, err := producer.SendMessage(topic_msg)
		if err != nil {
			log.Warn("Producer ", k.TopicName, ": Err sending message=", err.Error())
		}

		log.Debug("Producer ", k.TopicName, ": Producing message partition=", partition, " offset=", offset)
		metrics.Metrics["kafka_messages_produced"].Inc()
	}
}
