package server

import (
	"encoding/json"

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
		log.Info(c.Method(), " ", c.Path())

		// Go to next middleware:
		return c.Next()
	})

	// Add version handlers
	app.Get("/version", handlerVersion)
	app.Get("/metadata", handlerMetadata)

	// Add handlers
	rest.BlocksAddHandlers(app)
	ws.BlocksAddHandlers(app)

	app.Listen(":" + config.Vars.Port)
}

func handlerVersion(c *fiber.Ctx) error {
	message := map[string]string{
		"version": config.Vars.Version,
	}

	json_message, _ := json.Marshal(message)

	return c.SendString(string(json_message))
}

func handlerMetadata(c *fiber.Ctx) error {
	message := map[string]string{
		"version":     config.Vars.Version,
		"name":        config.Vars.Name,
		"description": "a go api template",
	}

	json_message, _ := json.Marshal(message)

	return c.SendString(string(json_message))
}
