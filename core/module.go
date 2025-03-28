package core

import (
	"errors"

	"github.com/labstack/echo/v4"
)

// Module defines the interface that all GoPHER modules must implement.
type Module interface {
	Key() string                         // Unique machine-readable key (e.g. "auth", "blog")
	Name() string                        // Human-readable display name
	Description() string                 // Short summary for UI or logging

	Init(services *CoreServices) error   // Called during app boot
	RegisterRoutes(e *echo.Echo)         // Hook for HTTP route setup
	Migrate() error                      // Create DB schema or seed data
	Update() error                       // Update schema or config (e.g. on deploy)
	RegisterJobs(scheduler *JobScheduler) // Optional background task registration
}

// moduleRegistry holds and manages enabled modules.
type moduleRegistry struct {
	modules map[string]Module
	order   []string // preserve order of registration
}

var registry = &moduleRegistry{
	modules: make(map[string]Module),
	order:   []string{},
}

// Enable registers a module with the system.
func Enable(m Module) {
	key := m.Key()
	if _, exists := registry.modules[key]; !exists {
		registry.modules[key] = m
		registry.order = append(registry.order, key)
	}
}

// Disable removes a module from the registry (non-persistent for now).
func Disable(key string) {
	delete(registry.modules, key)
	// Optionally remove from order (not essential now)
}

// InitAll initializes, registers routes, and hooks jobs for all enabled modules.
func InitAll(services *CoreServices) error {
	for _, key := range registry.order {
		m := registry.modules[key]
		services.Logger.Info().
			Str("module", m.Key()).
			Msg("Initializing module")

		if err := m.Init(services); err != nil {
			return err
		}

		m.RegisterRoutes(services.Echo)
		m.RegisterJobs(services.Jobs)
	}
	return nil
}

// MigrateAll runs Migrate on all enabled modules.
func MigrateAll() error {
	for _, key := range registry.order {
		if err := registry.modules[key].Migrate(); err != nil {
			return err
		}
	}
	return nil
}

// UpdateAll runs Update on all enabled modules.
func UpdateAll() error {
	for _, key := range registry.order {
		if err := registry.modules[key].Update(); err != nil {
			return err
		}
	}
	return nil
}

// Get returns a module by key if registered.
func Get(key string) (Module, error) {
	if m, ok := registry.modules[key]; ok {
		return m, nil
	}
	return nil, errors.New("module not found")
}

// List returns all registered modules.
func List() []Module {
	var all []Module
	for _, key := range registry.order {
		all = append(all, registry.modules[key])
	}
	return all
}
