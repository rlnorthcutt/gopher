package core

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all environment-level application settings.
type Config struct {
	Port           string
	Environment    string
	Debug          bool
	BaseURL        string
	PocketBaseDir  string
	CacheEnabled   bool
	EnableTelemetry bool
}

// LoadConfig reads from .env and environment variables into Config.
func LoadConfig() *Config {
	_ = godotenv.Load(".env") // silently ignore if missing

	return &Config{
		Port:            getEnv("PORT", "8080"),
		Environment:     getEnv("ENV", "development"),
		Debug:           getEnvBool("DEBUG", true),
		BaseURL:         getEnv("BASE_URL", "http://localhost:8080"),
		PocketBaseDir:   getEnv("PB_DIR", "./pb_data"),
		CacheEnabled:    getEnvBool("CACHE_ENABLED", true),
		EnableTelemetry: getEnvBool("TELEMETRY", false),
	}
}

// Helpers
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		log.Printf("invalid bool value for %s: %v", key, err)
		return defaultVal
	}
	return b
}