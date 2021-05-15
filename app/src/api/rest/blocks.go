package rest

import (
	"github.com/gofiber/fiber/v2"

	"github.com/geometry-labs/app/config"
)

func BlocksAddHandlers(app *fiber.App) {

	prefix := config.Vars.RestPrefix + "/blocks"

	app.Get(prefix+"/", handlerGetBlock)
}

func handlerGetBlock(c *fiber.Ctx) error {
	return c.SendString(`{"block": "rests"}`)
}
