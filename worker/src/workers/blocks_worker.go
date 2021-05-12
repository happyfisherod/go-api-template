package workers

import (
	"strings"

	"github.com/geometry-labs/worker/config"
	"github.com/geometry-labs/worker/kafka"
	"github.com/geometry-labs/worker/utils"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

func StartBlocksWorker() {
	go blocksWorker()
}

func blocksWorker() {
	consumer_topic_name := config.Vars.BlocksWorkerConsumerTopic
	producer_topic_name := config.Vars.BlocksWorkerProducerTopic

	// Check topic names
	if utils.StringInSlice(consumer_topic_name, strings.Split(config.Vars.ConsumerTopics, ",")) == false {
		log.Panic("Blocks Worker: invalid BLOCKS_WORKER_CONSUMER_TOPIC value. MUST be a topic in CONSUMER_TOPICS")
	}
	if utils.StringInSlice(producer_topic_name, strings.Split(config.Vars.ProducerTopics, ",")) == false {
		log.Panic("Blocks Worker: invalid BLOCKS_WORKER_PRODUCER_TOPIC value. MUST be a topic in PRODUCER_TOPICS")
	}

	consumer_topic_chan := make(chan *sarama.ConsumerMessage)
	producer_topic_chan := kafka.KafkaTopicProducers[producer_topic_name].TopicChan

	// Register consumer channel
	broadcaster_output_chan_id := kafka.Broadcasters[consumer_topic_name].AddWorkerChannel(consumer_topic_chan)
	defer func() {
		kafka.Broadcasters[consumer_topic_name].RemoveWorkerChannel(broadcaster_output_chan_id)
	}()

	log.Debug("Blocks Worker: started working")
	for {
		consumer_topic_msg := <-consumer_topic_chan

		producer_topic_msg := &sarama.ProducerMessage{
			Topic:     producer_topic_name,
			Partition: -1,
			Key:       sarama.StringEncoder(string(consumer_topic_msg.Key)),
			Value:     sarama.StringEncoder(string(consumer_topic_msg.Value)),
		}

		producer_topic_chan <- producer_topic_msg

		log.Debug("Blocks worker: last seen block #", string(consumer_topic_msg.Key))
	}
}
