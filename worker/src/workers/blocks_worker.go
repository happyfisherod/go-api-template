package workers

import (
	"strings"

	"github.com/geometry-labs/worker/config"
	"github.com/geometry-labs/worker/kafka"
	"github.com/geometry-labs/worker/utils"

	log "github.com/sirupsen/logrus"
	confluent "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
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

	consumer_topic_chan := make(chan *confluent.Message)
	producer_topic_chan := kafka.KafkaTopicProducers[producer_topic_name].TopicChan

	// Register consumer channel
	broadcaster_output_chan_id := kafka.Broadcasters[consumer_topic_name].AddOutputChannel(consumer_topic_chan)
	defer func() {
		kafka.Broadcasters[consumer_topic_name].RemoveOutputChannel(broadcaster_output_chan_id)
	}()

	log.Debug("Blocks Worker: started working")
	for {
		topic_msg := <-consumer_topic_chan

		producer_topic_chan <- topic_msg

		log.Debug("Blocks worker: last seen block #", string(topic_msg.Key))
	}
}
