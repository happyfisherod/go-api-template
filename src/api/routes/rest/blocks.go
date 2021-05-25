package rest

import (
	fiber "github.com/gofiber/fiber/v2"

	"github.com/geometry-labs/go-service-template/core"
)

func BlocksAddHandlers(app *fiber.App) {

	prefix := core.Vars.RestPrefix + "/blocks"

	app.Get(prefix+"/", handlerGetBlock)
}

func handlerGetBlock(c *fiber.Ctx) error {
	return c.SendString(`{"block": "rests"}`)
}
