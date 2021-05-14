package service

import (
	"github.com/geometry-labs/api/config"
	"github.com/geometry-labs/api/kafka"
)

func StartConsumeBlocks(returnChan chan string) {
	consumer, err := kafka.NewKafkaConsumer(config.Vars.KafkaBrokerURL, config.Vars.Name+"-group", "latest")
	if err != nil {
		panic(err)
	}
	defer consumer.CloseConsumer()

	err = consumer.Subscribe(config.Vars.TopicNames, nil)
	if err != nil {
		panic(err)
	}

	consumer.StartRead(returnChan)
}
