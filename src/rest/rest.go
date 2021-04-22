package rest

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func StartHttpServer(port string, prefix string) {

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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":" + port)
}
