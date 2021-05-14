package kafka

import (
	"strings"

	"github.com/geometry-labs/api/config"
	"github.com/geometry-labs/api/metrics"

	log "github.com/sirupsen/logrus"
	confluent "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func Start() {
	kafka_broker := config.Vars.KafkaBrokerURL
	topics := config.Vars.TopicNames
	schemas := config.Vars.SchemaNames

	if kafka_broker == "" {
		log.Panic("No kafka broker url provided")
	}

	//time.Sleep(time.Minute)
	for _, schemaNameAndFilePairs := range schemas {
		schemaNameAndFile := strings.Split(schemaNameAndFilePairs, ":")
		//_, _ = RegisterSchema(schemaNameAndFile[0], false, schemaNameAndFile[1], true)
		id, _ := RetriableRegisterSchema(RegisterSchema, schemaNameAndFile[0], false, schemaNameAndFile[1], true)
		log.Info("Schema id for ", schemaNameAndFile[0], " is ", id)
	}

	for _, t := range topics {
		// Broadcaster indexed in Broadcasters map
		newBroadcaster(t, make(chan *confluent.Message))

		topic_consumer := &KafkaTopicConsumer{
			kafka_broker,
			t,
			Broadcasters[t],
		}

		go topic_consumer.consumeAndBroadcastTopics()
	}
}

type KafkaTopicConsumer struct {
	BrokerURL   string
	TopicName   string
	Broadcaster *TopicBroadcaster
}

func (k *KafkaTopicConsumer) consumeAndBroadcastTopics() {

	consumer, err := confluent.NewConsumer(&confluent.ConfigMap{
		"bootstrap.servers": k.BrokerURL,
		"group.id":          config.Vars.KafkaGroupID,
		"auto.offset.reset": "latest",
	})

	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	consumer.SubscribeTopics([]string{k.TopicName}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		metrics.Metrics["kafka_messages_consumed"].Inc()

		if err == nil {

			// NOTE: use select statement for non-blocking channels
			select {
			case k.Broadcaster.InputChan <- msg:
			default:
			}
		}
	}
}
