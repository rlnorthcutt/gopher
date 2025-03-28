package core

import (
	"github.com/labstack/echo/v4"
	"github.com/pocketbase/pocketbase"
)

// Bootstrap initializes the core application and returns Echo + CoreServices.
func Bootstrap() (*echo.Echo, *CoreServices, error) {
	// Load config
	config := LoadConfig()

	// Set up logger and tracer
	logger := SetupLogger(config.AppName)
	tracer, err := SetupTracer(config.AppName)
	if err != nil {
		logger.Error().Err(err).Msg("failed to initialize tracer")
		return nil, nil, err
	}

	// Initialize Echo and PocketBase
	e := echo.New()
	pb := pocketbase.New()

	// Init other services
	cache := NewCache()
	jobs := NewJobScheduler(nil)

	// Build CoreServices
	services := New(e, pb, logger, tracer, config, cache, jobs)
	jobs.services = services // Inject services into JobScheduler

	// Register middleware
	e.Use(RequestIDMiddleware())
	e.Use(LoggingMiddleware(logger))

	// Init route system and metrics
	InitRouter(e)
	AttachMetricsRoute(e)

	return e, services, nil
}
