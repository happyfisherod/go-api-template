package metrics

import (
	"fmt"
	"github.com/geometry-labs/go-service-template/config"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	//core.GetEnvironment()
	config.Vars.ConfigFile = "config.api.test"
	config.Vars.ConfigType = "yaml"
	config.Vars.ConfigPath = "../../../envfiles"
	config.ConfigInit()
}

func TestStart(t *testing.T) {
	assert := assert.New(t)

	// Start metrics server
	Start()

	Metrics["requests_amount"].Inc()
	Metrics["kafka_messages_consumed"].Inc()
	Metrics["websockets_connected"].Inc()
	Metrics["websockets_bytes_written"].Inc()

	resp, err := http.Get(fmt.Sprintf("http://localhost:%s%s", config.Config.MetricsPort, config.Config.MetricsPrefix))
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)
}
