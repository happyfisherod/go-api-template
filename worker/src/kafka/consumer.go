package kafka

import (
	"strings"

	"github.com/geometry-labs/worker/config"
	"github.com/geometry-labs/worker/metrics"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

func StartConsumers() {
	kafka_broker := config.Vars.KafkaBrokerURL
	consumer_topics := strings.Split(config.Vars.ConsumerTopics, ",")

	log.Debug("Start Consumer: kafka_broker=", kafka_broker, " consumer_topics=", consumer_topics)

	for _, t := range consumer_topics {
		// Broadcaster indexed in Broadcasters map
		// Starts go routine
		newBroadcaster(t, make(chan *sarama.ConsumerMessage))

		topic_consumer := &KafkaTopicConsumer{
			kafka_broker,
			t,
			Broadcasters[t],
		}

		// One routine per topic
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

	consumer, err := sarama.NewConsumer([]string{k.BrokerURL}, nil)
	if err != nil {
		log.Panic("KAFKA CONSUMER PANIC: ", err.Error())
	}
	defer consumer.Close()

	offset := sarama.OffsetOldest
	partitions, err := consumer.Partitions(k.TopicName)
	if err != nil {
		log.Panic("KAFKA CONSUMER PANIC: ", err.Error())
	}

	log.Debug("Consumer ", k.TopicName, ": started consuming")
	for _, p := range partitions {
		pc, _ := consumer.ConsumePartition(k.TopicName, p, offset)

		// One routine per partition
		go func(pc sarama.PartitionConsumer) {
			for {
				topic_msg := <-pc.Messages()

				log.Debug("Consumer ", k.TopicName, ": consumed message key=", string(topic_msg.Key))
				metrics.Metrics["kafka_messages_consumed"].Inc()
				k.Broadcaster.ConsumerChan <- topic_msg
				log.Debug("Consumer ", k.TopicName, ": broadcasted message key=", string(topic_msg.Key))
			}
		}(pc)
	}
}
