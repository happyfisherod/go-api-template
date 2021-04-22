package endpoints

import (
	"github.com/gofiber/fiber/v2"

	"github.com/geometry-labs/api/config"
)

func BlocksAddHandlers(app *fiber.App) {

	prefix := config.Vars.Prefix + "/blocks"

	app.Get(prefix+"/", handlerGetBlock)
}

func handlerGetBlock(c *fiber.Ctx) error {
	return c.SendString(`{"block": "mock"}`)
}
