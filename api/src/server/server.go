package server

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/geometry-labs/api/config"
	"github.com/geometry-labs/api/server/rest"
	"github.com/geometry-labs/api/server/ws"
)

func Start() {

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		// logging
		log.WithFields(log.Fields{
			"hostname": c.Hostname(),
			"IP":       c.IP(),
			"Method":   c.Method(),
		}).Info(c.Path())

		// Go to next middleware:
		return c.Next()
	})

	// Add handlers
	rest.BlocksAddHandlers(app)
	ws.BlocksAddHandlers(app)

	app.Listen(":" + config.Vars.Port)
}
