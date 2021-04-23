package ws

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	"github.com/geometry-labs/api/config"
	"github.com/geometry-labs/api/kafka"
	"github.com/geometry-labs/api/metrics"
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

	app.Get(prefix+"/", websocket.New(handlerGetBlocks))
}

func handlerGetBlocks(c *websocket.Conn) {
	metrics.Metrics["websockets_connected"].Inc()

	// Add broadcaster
	topic_chan := make(chan *kafka.Message)
	id := kafka.Broadcasters["blocks"].AddOutputChannel(topic_chan)
	defer func() {
		// Remove broadcaster
		kafka.Broadcasters["blocks"].RemoveOutputChannel(id)
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
		err = c.WriteMessage(websocket.TextMessage, msg.Value)
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
