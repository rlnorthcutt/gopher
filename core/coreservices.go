package core

import (
	"github.com/labstack/echo/v4"
	"github.com/pocketbase/pocketbase"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

// CoreServices contains all shared systems used across the application.
type CoreServices struct {
	Echo        *echo.Echo
	PB          *pocketbase.PocketBase
	Logger      zerolog.Logger
	Tracer      trace.Tracer
	Cache       *Cache
	Config      *Config
	Jobs        *JobScheduler
	Data        *DataService
	Permissions *PermissionsManager
}

func New(
	echo *echo.Echo,
	pb *pocketbase.PocketBase,
	logger zerolog.Logger,
	tracer trace.Tracer,
	config *Config,
	cache *Cache,
	jobs *JobScheduler,
) *CoreServices {
	data := NewDataService(pb)
	permissions := NewPermissionsManager(pb)

	// Register entity runtime for global Save/Delete access
	RegisterEntityRuntime(data)

	return &CoreServices{
		Echo:        echo,
		PB:          pb,
		Logger:      logger,
		Tracer:      tracer,
		Cache:       cache,
		Config:      config,
		Jobs:        jobs,
		Data:        data,
		Permissions: permissions,
	}
}

