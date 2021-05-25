package routes

import (
	"encoding/json"

	swagger "github.com/arsmn/fiber-swagger/v2"
	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/geometry-labs/go-service-template/core"

	_ "github.com/geometry-labs/go-service-template/api/docs"
	"github.com/geometry-labs/go-service-template/api/routes/rest"
	"github.com/geometry-labs/go-service-template/api/routes/ws"
)

// @title Go api template docs
// @version 2.0
// @description This is a sample server server.
func Start() {

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		// logging
		log.Info(c.Method(), " ", c.Path())

		// Go to next middleware:
		return c.Next()
	})

	// Swagger docs
	app.Get("/docs/*", swagger.Handler)

	// Add version handlers
	app.Get("/version", handlerVersion)
	app.Get("/metadata", handlerMetadata)

	// Add handlers
	rest.BlocksAddHandlers(app)
	ws.BlocksAddHandlers(app)

	go app.Listen(":" + core.Vars.Port)
}

// Version
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /version [get]
func handlerVersion(c *fiber.Ctx) error {
	message := map[string]string{
		"version": core.Vars.Version,
	}

	json_message, _ := json.Marshal(message)

	return c.SendString(string(json_message))
}

// Metadata
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /metadata [get]
func handlerMetadata(c *fiber.Ctx) error {
	message := map[string]string{
		"version":     core.Vars.Version,
		"name":        core.Vars.Name,
		"description": "a go api template",
	}

	json_message, _ := json.Marshal(message)

	return c.SendString(string(json_message))
}
