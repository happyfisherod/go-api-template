package rest

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/geometry-labs/api/config"
	"github.com/geometry-labs/api/rest/endpoints"
)

type Router map[string]func(c *fiber.Ctx)

func StartHttpServer() {

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		// logging
		log.WithFields(log.Fields{
			"hostname": c.Hostname(),
			"IP":       c.IP(),
			"Method":   c.Method(),
			"Path":     c.Path(),
		}).Info("")

		// Go to next middleware:
		return c.Next()
	})

	// Add handlers
	endpoints.BlocksAddHandlers(app)

	app.Listen(":" + config.Vars.Port)
}
