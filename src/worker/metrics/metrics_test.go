package metrics

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/geometry-labs/go-service-template/core"
)

func init() {
	//core.GetEnvironment()
	core.Vars.ConfigFile = "config.worker.test"
	core.Vars.ConfigType = "yaml"
	core.Vars.ConfigPath = "../../../envfiles"
	core.ConfigInit()
}

func TestStart(t *testing.T) {
	assert := assert.New(t)

	// Start metrics server
	Start()

	Metrics["requests_amount"].Inc()
	Metrics["kafka_messages_consumed"].Inc()
	Metrics["websockets_connected"].Inc()
	Metrics["websockets_bytes_written"].Inc()

	resp, err := http.Get(fmt.Sprintf("http://localhost:%s%s", core.Config.MetricsPort, core.Config.MetricsPrefix))
	assert.Equal(nil, err)
	assert.Equal(200, resp.StatusCode)
}
