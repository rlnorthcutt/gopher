package core

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
)

// NewPB initializes and returns a PocketBase instance using the configured directory.
func NewPB(config *Config) *pocketbase.PocketBase {
	if err := os.MkdirAll(config.PocketBaseDir, 0755); err != nil {
		log.Fatalf("Failed to create PB data directory: %v", err)
	}

	app := pocketbase.New()
	app.DataDir = config.PocketBaseDir

	// @TODO: Attach automatic migration tool
	app.OnBeforeServe().Add(pbAutoMigrate)

	return app
}

// Placeholder - to be built out later
func pbAutoMigrate(e *core.ServeEvent) error {
	migrate.MustRegister(app, migrate.FS(
		os.DirFS("pb_migrations"),
		".",
	))
	return nil
}
