package healthcheck

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/geometry-labs/go-service-template/api/routes"
	"github.com/geometry-labs/go-service-template/core"
)

func init() {
	core.GetEnvironment()
}

func TestHealthCheck(t *testing.T) {
	assert := assert.New(t)

	// Start api
	routes.Start()

	// Start healthcheck
	Start()

	resp, err := http.Get("http://localhost:" + core.Vars.HealthPort + core.Vars.HealthPrefix)
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)
}
