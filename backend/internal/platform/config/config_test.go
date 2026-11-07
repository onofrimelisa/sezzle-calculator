package config_test

import (
	"testing"
	"time"

	"github.com/sezzle/calculator-backend/internal/platform/config"
	"github.com/stretchr/testify/assert"
)

func TestLoad_Defaults(t *testing.T) {
	cfg := config.Load()
	assert.Equal(t, 8080, cfg.Port)
	assert.Equal(t, "info", cfg.LogLevel)
	assert.Equal(t, "*", cfg.CORSOrigins)
	assert.Equal(t, 5*time.Second, cfg.ReadTimeout)
	assert.Equal(t, 10*time.Second, cfg.WriteTimeout)
}

func TestLoad_FromEnv(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("CORS_ORIGINS", "http://localhost:3000")
	t.Setenv("READ_TIMEOUT", "10s")
	t.Setenv("WRITE_TIMEOUT", "30s")

	cfg := config.Load()
	assert.Equal(t, 9090, cfg.Port)
	assert.Equal(t, "debug", cfg.LogLevel)
	assert.Equal(t, "http://localhost:3000", cfg.CORSOrigins)
	assert.Equal(t, 10*time.Second, cfg.ReadTimeout)
	assert.Equal(t, 30*time.Second, cfg.WriteTimeout)
}
