package metrics

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/geometry-labs/worker/config"
)

func init() {
	config.GetEnvironment()
}

func TestStart(t *testing.T) {
	assert := assert.New(t)

	// Start metrics server
	Start()

	Metrics["kafka_messages_produced"].Inc()
	Metrics["kafka_messages_consumed"].Inc()

	resp, err := http.Get(fmt.Sprintf("http://localhost:%s%s", config.Vars.MetricsPort, config.Vars.MetricsPrefix))
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)
}
