package workers

import (
	"github.com/geometry-labs/worker/kafka"

	log "github.com/sirupsen/logrus"
	confluent "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func StartBlocksWorker() {
	go blocksWorker()
}

func blocksWorker() {
	consumer_topic_name := "raw-blocks"
	producer_topic_name := "blocks"

	consumer_topic_chan := make(chan *confluent.Message)
	producer_topic_chan := kafka.KafkaTopicProducers[producer_topic_name].TopicChan

	// Register consumer channel
	broadcaster_output_chan_id := kafka.Broadcasters[consumer_topic_name].AddOutputChannel(consumer_topic_chan)
	defer func() {
		kafka.Broadcasters[consumer_topic_name].RemoveOutputChannel(broadcaster_output_chan_id)
	}()

	for {
		topic_msg := <-consumer_topic_chan

		producer_topic_chan <- topic_msg
	}
}
