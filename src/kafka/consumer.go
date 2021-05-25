package kafka

import (
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"

	"github.com/geometry-labs/go-service-template/core"
)

func StartConsumers() {
	kafka_broker := core.Vars.KafkaBrokerURL
	consumer_topics := core.Vars.ConsumerTopics

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

	offset := sarama.OffsetOldest
	partitions, err := consumer.Partitions(k.TopicName)
	if err != nil {
		log.Panic("KAFKA CONSUMER PARTITIONS PANIC: ", err.Error())
	}

	log.Debug("Consumer ", k.TopicName, ": Started consuming")
	for _, p := range partitions {
		pc, err := consumer.ConsumePartition(k.TopicName, p, offset)

		if err != nil {
			log.Panic("KAFKA CONSUMER PARTITIONS PANIC: ", err.Error())
		}
		if pc == nil {
			log.Panic("KAFKA CONSUMER PARTITIONS PANIC: Failed to create PartitionConsumer")
		}

		// Watch errors
		// go func() {
		// 		for err := range pc.Errors() {
		// log.Warn("KAFKA CONSUMER WARN: ", err.Error())
		// 	}
		// }()

		// One routine per partition
		go func(pc sarama.PartitionConsumer) {
			for {
				topic_msg := <-pc.Messages()
				log.Debug("Consumer ", k.TopicName, ": Consumed message key=", string(topic_msg.Key))

				// Broadcast
				k.Broadcaster.ConsumerChan <- topic_msg

				log.Debug("Consumer ", k.TopicName, ": Broadcasted message key=", string(topic_msg.Key))
			}
		}(pc)
	}
}
