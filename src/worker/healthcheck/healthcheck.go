package healthcheck

import (
	"net/http"
	"net/url"
	"time"

	"github.com/InVisionApp/go-health/v2"
	"github.com/InVisionApp/go-health/v2/checkers"
	"github.com/InVisionApp/go-health/v2/handlers"
	log "github.com/sirupsen/logrus"

	"github.com/geometry-labs/app/config"
)

// TODO split API and WORKER
func Start() {
	// Create a new health instance
	h := health.New()

	// Create a couple of checks
	blocksCheckerURL, _ := url.Parse("http://localhost:" + config.Vars.Port + config.Vars.RestPrefix + "/blocks")
	blocksChecker, _ := checkers.NewHTTP(&checkers.HTTPConfig{
		URL: blocksCheckerURL,
	})

	// Add the checks to the health instance
	h.AddChecks([]*health.Config{
		{
			Name:     "blocks-rest-check",
			Checker:  blocksChecker,
			Interval: time.Duration(config.Vars.HealthPollingInterval) * time.Second,
			Fatal:    true,
		},
	})

	//  Start the healthcheck process
	if err := h.Start(); err != nil {
		log.Fatalf("Unable to start healthcheck: %v", err)
	}

	// Define a healthcheck endpoint and use the built-in JSON handler
	http.HandleFunc(config.Vars.HealthPrefix, handlers.NewJSONHandlerFunc(h, nil))
	go http.ListenAndServe(":"+config.Vars.HealthPort, nil)
	log.Println("Started Healthcheck:", config.Vars.HealthPort)
}
