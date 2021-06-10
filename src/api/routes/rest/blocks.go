package rest

import (
	"encoding/json"
	"github.com/geometry-labs/go-service-template/api/service"
	"github.com/geometry-labs/go-service-template/config"
	"github.com/geometry-labs/go-service-template/global"
	"github.com/geometry-labs/go-service-template/models"
	fiber "github.com/gofiber/fiber/v2"
)

func BlocksAddHandlers(app *fiber.App) {

	prefix := config.Config.RestPrefix + "/blocks"

	app.Get(prefix+"/", handlerGetAllBlock)
	app.Get(prefix+"/index", handlerIndex)
	app.Get(prefix+"/hash/:hash", handlerGetBlockWhereHash)
	app.Get(prefix+"/height/:height", handlerGetBlockWhereHeight)
	app.Get(prefix+"/start/:height", handlerGetBlockFromHeight)
	app.Get(prefix+"/created_by/:address", handlerGetBlockWhereCreatedBy)
}

func handlerIndex(c *fiber.Ctx) error {
	err := global.GetGlobal().Blocks.CreateIndex()
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString("migrate done")

}

// Blocks
// @Summary Get all blocks.
// @Description Get all blocks in the system.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} *[]models.BlockRaw
// @Router /blocks [get]
func handlerGetAllBlock(c *fiber.Ctx) error {
	blocks := service.AllBlocks()
	jsonBytes, err := json.Marshal(blocks)
	if err != nil {
		jsonBytes = []byte("{}")
	}
	return c.SendString(string(jsonBytes))
}

// Blocks/hash
// @Summary Get all blocks for a Hash.
// @Description Get all blocks in the system for a given hash.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} *[]models.BlockRaw
// @Router /blocks/hash/:hash [get]
func handlerGetBlockWhereHash(c *fiber.Ctx) error {
	blocks := service.FindWhereHash(c.Params("hash"))
	jsonBytes, err := json.Marshal(blocks)
	if err != nil {
		jsonBytes = []byte("{}")
	}
	return c.SendString(string(jsonBytes))
}

// Blocks/height
// @Summary Get all blocks for a Height.
// @Description Get all blocks in the system for a given height.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} *[]models.BlockRaw
// @Router /blocks/height/:height [get]
func handlerGetBlockWhereHeight(c *fiber.Ctx) error {
	heightString := c.Params("height")
	_ = handlerHelperInvalidHeight(c, heightString)
	blocks := service.FindWhereHeight(heightString)
	jsonBytes, err := json.Marshal(blocks)
	if err != nil {
		jsonBytes = []byte("[]")
	}
	return c.SendString(string(jsonBytes))
}

// Blocks/start
// @Summary Get all blocks from a Height.
// @Description Get all blocks in the system from a given height.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} *[]models.BlockRaw
// @Router /blocks/start/:height [get]
func handlerGetBlockFromHeight(c *fiber.Ctx) error {
	heightString := c.Params("height")
	_ = handlerHelperInvalidHeight(c, heightString)
	blocks := service.FindFromHeight(heightString)
	jsonBytes, err := json.Marshal(blocks)
	if err != nil {
		jsonBytes = []byte("[]")
	}
	return c.SendString(string(jsonBytes))
}

// Blocks/created_by
// @Summary Get all blocks by a created_by address.
// @Description Get all blocks in the system by a created_by address.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} *[]models.BlockRaw
// @Router /blocks/created_by/:peerId [get]
func handlerGetBlockWhereCreatedBy(c *fiber.Ctx) error {
	address := c.Params("address")
	blocks := service.FindWhereCreatedBy(address)
	jsonBytes, err := json.Marshal(blocks)
	if err != nil {
		jsonBytes = []byte("[]")
	}
	return c.SendString(string(jsonBytes))
}

func handlerHelperInvalidHeight(c *fiber.Ctx, heightString string) error {
	validParam := models.ValidateHeight(heightString)
	if validParam == false {
		return c.SendString("wrong param")
	}
	return nil

}
