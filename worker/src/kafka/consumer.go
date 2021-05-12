package kafka

import (
	"strings"
	"time"

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
		log.Debug("Start Consumers: Starting ", t, " consumer...")
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
		log.Panic("KAFKA CONSUMER NEWCONSUMER PANIC: ", err.Error())
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Panic("KAFKA CONSUMER CLOSE PANIC: ", err.Error())
		}
	}()

	// offset := sarama.OffsetOldest
	partitions, err := consumer.Partitions(k.TopicName)
	if err != nil {
		log.Panic("KAFKA CONSUMER PARTITIONS PANIC: ", err.Error())
	}

	log.Debug("Consumer ", k.TopicName, ": Started consuming")
	for _, p := range partitions {
		pc, _ := consumer.ConsumePartition(k.TopicName, p, 0)

		// Watch errors
		go func() {
			for err := range pc.Errors() {
				log.Warn("KAFKA CONSUMER WARN: ", err.Error())
			}
		}()

		// One routine per partition
		go func(pc sarama.PartitionConsumer) {
			for {
				select {
				case topic_msg := <-pc.Messages():
					log.Debug("Consumer ", k.TopicName, ": Consumed message key=", string(topic_msg.Key))
					metrics.Metrics["kafka_messages_consumed"].Inc()

					// Broadcast
					k.Broadcaster.ConsumerChan <- topic_msg

					log.Debug("Consumer ", k.TopicName, ": Broadcasted message key=", string(topic_msg.Key))
				case <-time.After(3 * time.Second):
					log.Debug("Consumer ", k.TopicName, ": No new messages...sleeping 3 seconds")
					continue
				}
			}
		}(pc)
	}
}
