package kafka

import (
	"github.com/Shopify/sarama"
)

// TODO use uuid for larger ID range
type BroadcasterID int

var LAST_BROADCASTER_ID BroadcasterID = 0

type TopicBroadcaster struct {

	// Input
	ConsumerChan chan *sarama.ConsumerMessage

	// Output
	WorkerChans map[BroadcasterID]chan *sarama.ConsumerMessage
}

var Broadcasters = map[string]*TopicBroadcaster{}

func newBroadcaster(topic_name string) {
	Broadcasters[topic_name] = &TopicBroadcaster{
		make(chan *sarama.ConsumerMessage),
		make(map[BroadcasterID]chan *sarama.ConsumerMessage),
	}

	go Broadcasters[topic_name].Start()
}

func (tb *TopicBroadcaster) AddWorkerChannel(topic_chan chan *sarama.ConsumerMessage) BroadcasterID {
	id := LAST_BROADCASTER_ID
	LAST_BROADCASTER_ID++

	tb.WorkerChans[id] = topic_chan

	return id
}

func (tb *TopicBroadcaster) RemoveWorkerChannel(id BroadcasterID) {
	_, ok := tb.WorkerChans[id]
	if ok {
		delete(tb.WorkerChans, id)
	}
}

func (tb *TopicBroadcaster) Start() {
	for {
		msg := <-tb.ConsumerChan

		for _, channel := range tb.WorkerChans {
			channel <- msg
		}
	}
}
