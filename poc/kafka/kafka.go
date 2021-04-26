package kafka

import (
	"encoding/json"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Ok struct {
	Value int
}

func ProducerPoc() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "myTopic"
	//for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
	//	p.Produce(&kafka.Message{
	//		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	//		Value:          []byte(word),
	//	}, nil)
	//}

	for _, word := range []Ok{Ok{20215}, Ok{20216}, Ok{20217}, Ok{20218}, Ok{20219}} {
		val, _ := json.Marshal(word)
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          val,
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

func ConsumerPoc() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}

/* Noting api for schema registry
1. register json schema for myTopic-Key
curl -X POST -H "Content-Type: application/vnd.schemaregistry.v1+json" \
  --data '{"schema": "{\"type\": \"string\"}"}' \
  http://localhost:8081/subjects/myTopic-key/versions
2. curl -X GET http://localhost:8081/subjects
3. curl -X GET http://localhost:8081/subjects/myTopic-key/versions
4. Get schema for  subject myTopic-key version 1
curl -X GET http://localhost:8081/subjects/myTopic-key/versions/1
5. Register a protobuf schema
curl -X POST -H "Content-Type: application/vnd.schemaregistry.v1+json" \
  --data '{"schemaType": "PROTOBUF", "schema": "syntax = \"proto3\";\npackage com.acme;\n\nmessage value_myTopic {\n  string myField1 = 1;\n}\n"}' \
  http://localhost:8081/subjects/myTopic-value/versions


*/
