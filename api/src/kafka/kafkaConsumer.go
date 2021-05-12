package kafka

import (
	"fmt"
	"github.com/geometry-labs/api/models"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KafkaConsumer struct {
	consumer  *kafka.Consumer
	isReading bool
	isClose   bool
}

func NewKafkaConsumer(broker string, group string, offset string) (*KafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          group,
		"auto.offset.reset": offset,
	})
	if err != nil {
		err := consumer.Close()
		return nil, err
	}
	return &KafkaConsumer{
		consumer:  consumer,
		isReading: false,
		isClose:   false,
	}, err
}

func (kafkaConsumer *KafkaConsumer) Subscribe(topics []string, rebalanceCb kafka.RebalanceCb) error {
	err := kafkaConsumer.consumer.SubscribeTopics(topics, rebalanceCb)
	return err
}

func (kafkaConsumer *KafkaConsumer) StartRead(returnChan chan string) {
	kafkaConsumer.isReading = true
	kafkaConsumer.read(returnChan)
	defer kafkaConsumer.StopRead()
}

func (kafkaConsumer *KafkaConsumer) read(returnChan chan string) {
	for kafkaConsumer.isReading {
		ev, err := kafkaConsumer.consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Block message read error: %v\n", err)
		} else {
			fmt.Printf("Consuming Block message\n")
			kafkaConsumer.handleEventRead(ev, returnChan)
		}
	}
}

func (kafkaConsumer *KafkaConsumer) handleEventRead(ev *kafka.Message, returnChan chan string) {
	block := models.BlockRaw{}
	err := protojson.Unmarshal(ev.Value, &block)
	if err != nil {
		fmt.Printf("Error in Unmarshall of Block comming from IconEtl: %v", err)
	}
	fmt.Printf("%s %v \n", "Block consumed from iconetl:", block.String())
	go sendToReturnChan(returnChan, block.String())
}

func sendToReturnChan(returnChan chan string, block string) {
	fmt.Println("sendToReturnChan()")
	returnChan <- block
}

func (kafkaConsumer *KafkaConsumer) StopRead() {
	fmt.Println("StopRead")
	kafkaConsumer.isReading = false
}

func (kafkaConsumer *KafkaConsumer) CloseConsumer() error {
	fmt.Println("CloseConsumer")
	kafkaConsumer.isReading = false
	err := kafkaConsumer.consumer.Close()
	kafkaConsumer.isClose = true
	return err
}
