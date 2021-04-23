package healthcheck

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/InVisionApp/go-health/v2"
	"github.com/InVisionApp/go-health/v2/checkers"
	"github.com/InVisionApp/go-health/v2/handlers"

	"github.com/geometry-labs/api/config"
)

func Start() {
	// Create a new health instance
	h := health.New()
	goodTestURL, _ := url.Parse("https://google.com")

	// Create a couple of checks
	goodHTTPCheck, _ := checkers.NewHTTP(&checkers.HTTPConfig{
		URL: goodTestURL,
	})

	// Add the checks to the health instance
	h.AddChecks([]*health.Config{
		{
			Name:     "good-check",
			Checker:  goodHTTPCheck,
			Interval: time.Duration(2) * time.Second,
			Fatal:    true,
		},
	})

	//  Start the healthcheck process
	if err := h.Start(); err != nil {
		log.Fatalf("Unable to start healthcheck: %v", err)
	}

	// Define a healthcheck endpoint and use the built-in JSON handler
	http.HandleFunc("/healthcheck", handlers.NewJSONHandlerFunc(h, nil))
	http.ListenAndServe(":"+config.Vars.HealthPort, nil)
}
