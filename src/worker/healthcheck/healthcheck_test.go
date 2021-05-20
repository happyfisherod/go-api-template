package healthcheck

import (
	"net/http"
	"testing"

	"github.com/geometry-labs/app/api"
	"github.com/geometry-labs/app/config"

	"github.com/stretchr/testify/assert"
)

func init() {
	config.GetEnvironment()
}

func TestHealthCheck(t *testing.T) {
	assert := assert.New(t)

	// Start api
	api.Start()

	// Start healthcheck
	Start()

	resp, err := http.Get("http://localhost:" + config.Vars.HealthPort + config.Vars.HealthPrefix)
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)
}
