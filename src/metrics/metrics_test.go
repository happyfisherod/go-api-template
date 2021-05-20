package metrics

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/geometry-labs/app/config"
)

func init() {
	config.GetEnvironment()
}

func TestStart(t *testing.T) {
	assert := assert.New(t)

	// Start metrics server
	Start()

	Metrics["requests_amount"].Inc()
	Metrics["kafka_messages_consumed"].Inc()
	Metrics["websockets_connected"].Inc()
	Metrics["websockets_bytes_written"].Inc()

	resp, err := http.Get(fmt.Sprintf("http://localhost:%s%s", config.Vars.MetricsPort, config.Vars.MetricsPrefix))
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)
}
