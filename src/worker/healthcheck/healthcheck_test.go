package healthcheck

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/geometry-labs/go-service-template/api/routes"
	"github.com/geometry-labs/go-service-template/core"
)

func init() {
	//core.GetEnvironment()
	core.Vars.ConfigFile = "config.api.test"
	core.Vars.ConfigType = "yaml"
	core.Vars.ConfigPath = "../../../envfiles"
	core.ConfigInit()
}

func TestHealthCheck(t *testing.T) {
	assert := assert.New(t)

	// Start api
	routes.Start()

	// Start healthcheck
	Start()

	resp, err := http.Get("http://localhost:" + core.Config.HealthPort + core.Config.HealthPrefix)
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)
}
