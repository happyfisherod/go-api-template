package ws

import (
	"fmt"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	gorilla "github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	confluent "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	"github.com/geometry-labs/api/config"
	"github.com/geometry-labs/api/kafka"
	"github.com/geometry-labs/api/metrics"
)

func init() {
	config.GetEnvironment()
	metrics.Start()
}

func TestHandlerGetBlocks(t *testing.T) {
	_ = assert.New(t)

	// Create topic broadcaster
	input_chan := make(chan *confluent.Message)
	broadcaster := &kafka.TopicBroadcaster{
		input_chan,
		make(map[kafka.BroadcasterID]chan *confluent.Message),
	}

	app := fiber.New()

	app.Use("/", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/", websocket.New(handlerGetBlocks(broadcaster)))
	go app.Listen(":9999")

	test_data := "Test Data"
	go func() {
		for {
			msg := &(confluent.Message{})
			msg.Value = []byte(test_data)

			input_chan <- msg

			time.Sleep(1 * time.Second)
		}
	}()

	fmt.Printf("here")
	// Validate message
	websocket_client, _, err := gorilla.DefaultDialer.Dial("ws://localhost:9999/", nil)
	if err != nil {
		t.Logf("Failed to connect to KafkaWebsocketServer")
		t.Fail()
	}
	defer websocket_client.Close()

	fmt.Printf("here")
	_, message, err := websocket_client.ReadMessage()
	if err != nil {
		t.Logf("Failed to read websocket")
		t.Fail()
	}

	fmt.Printf("here")
	if string(message) != test_data {
		t.Logf("Failed to validate data")
		t.Fail()
	}

}
