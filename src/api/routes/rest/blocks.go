package rest

import (
	"github.com/geometry-labs/go-service-template/config"
	fiber "github.com/gofiber/fiber/v2"
)

func BlocksAddHandlers(app *fiber.App) {

	prefix := config.Config.RestPrefix + "/blocks"

	app.Get(prefix+"/", handlerGetBlock)
}

func handlerGetBlock(c *fiber.Ctx) error {
	return c.SendString(`{"block": "rests"}`)
}
