package ws

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	confluent "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	"github.com/geometry-labs/app/config"
	"github.com/geometry-labs/app/kafka"
	"github.com/geometry-labs/app/metrics"
)

func BlocksAddHandlers(app *fiber.App) {

	prefix := config.Vars.WebsocketPrefix + "/blocks"

	app.Use(prefix, func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get(prefix+"/", websocket.New(handlerGetBlocks(kafka.Broadcasters["blocks"])))
}

func handlerGetBlocks(broadcaster *kafka.TopicBroadcaster) func(c *websocket.Conn) {

	return func(c *websocket.Conn) {
		metrics.Metrics["websockets_connected"].Inc()

		// Add broadcaster
		topic_chan := make(chan *confluent.Message)
		id := broadcaster.AddOutputChannel(topic_chan)
		defer func() {
			// Remove broadcaster
			broadcaster.RemoveOutputChannel(id)
		}()

		// Read for close
		client_close_sig := make(chan bool)
		go func() {
			for {
				_, _, err := c.ReadMessage()
				if err != nil {
					client_close_sig <- true
					break
				}
			}
		}()

		for {
			// Read
			msg := <-topic_chan

			// Broadcast
			err := c.WriteMessage(websocket.TextMessage, msg.Value)
			metrics.Metrics["websockets_bytes_written"].Add(float64(len(msg.Value)))
			if err != nil {
				break
			}

			// check for client close
			select {
			case _ = <-client_close_sig:
				break
			default:
				continue
			}
		}
	}
}
