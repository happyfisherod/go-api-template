package transformers

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/Shopify/sarama.v1"

	"github.com/geometry-labs/go-service-template/core"
	"github.com/geometry-labs/go-service-template/kafka"

	"github.com/geometry-labs/go-service-template/worker/utils"
)

func StartBlocksTransformer() {
	go blocksTransformer()
}

func blocksTransformer() {
	consumer_topic_name := "blocks"
	producer_topic_name := "blocks-ws"

	// Check topic names
	if utils.StringInSlice(consumer_topic_name, core.Vars.ConsumerTopics) == false {
		log.Panic("Blocks Worker: invalid BLOCKS_WORKER_CONSUMER_TOPIC value. MUST be a topic in CONSUMER_TOPICS")
	}
	if utils.StringInSlice(producer_topic_name, core.Vars.ProducerTopics) == false {
		log.Panic("Blocks Worker: invalid BLOCKS_WORKER_PRODUCER_TOPIC value. MUST be a topic in PRODUCER_TOPICS")
	}

	consumer_topic_chan := make(chan *sarama.ConsumerMessage)
	producer_topic_chan := kafka.KafkaTopicProducers[producer_topic_name].TopicChan
	// create a channel for postgres

	// Register consumer channel
	broadcaster_output_chan_id := kafka.Broadcasters[consumer_topic_name].AddBroadcastChannel(consumer_topic_chan)
	defer func() {
		kafka.Broadcasters[consumer_topic_name].RemoveBroadcastChannel(broadcaster_output_chan_id)
	}()

	log.Debug("Blocks Worker: started working")
	for {
		consumer_topic_msg := <-consumer_topic_chan

		producer_topic_msg := &sarama.ProducerMessage{
			Topic: producer_topic_name,
			Key:   sarama.ByteEncoder(consumer_topic_msg.Key),
			Value: sarama.ByteEncoder(consumer_topic_msg.Value),
		}

		producer_topic_chan <- producer_topic_msg

		log.Debug("Blocks worker: last seen block #", string(consumer_topic_msg.Key))
	}
}
