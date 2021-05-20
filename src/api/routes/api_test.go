package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/geometry-labs/app/config"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.GetEnvironment()
}

func TestHandlerVersion(t *testing.T) {
	assert := assert.New(t)

	app := fiber.New()

	app.Get("/", handlerVersion)

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	// Read body
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	body_map := make(map[string]interface{})
	err = json.Unmarshal(bytes, &body_map)
	assert.Equal(nil, err)

	// Verify body
	assert.NotEqual(0, len(body_map["version"].(string)))
}

func TestHandlerMetadata(t *testing.T) {
	assert := assert.New(t)

	app := fiber.New()

	app.Get("/", handlerMetadata)

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)

	// Read body
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	assert.Equal(nil, err)

	body_map := make(map[string]interface{})
	err = json.Unmarshal(bytes, &body_map)
	assert.Equal(nil, err)

	// Verify body
	assert.NotEqual(0, len(body_map["version"].(string)))
	assert.NotEqual(0, len(body_map["name"].(string)))
	assert.NotEqual(0, len(body_map["description"].(string)))
}
