package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port         int
	LogLevel     string
	CORSOrigins  string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func Load() Config {
	return Config{
		Port:         getEnvInt("PORT", 8080),
		LogLevel:     getEnvStr("LOG_LEVEL", "info"),
		CORSOrigins:  getEnvStr("CORS_ORIGINS", "*"),
		ReadTimeout:  getEnvDuration("READ_TIMEOUT", 5*time.Second),
		WriteTimeout: getEnvDuration("WRITE_TIMEOUT", 10*time.Second),
	}
}

func getEnvStr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return fallback
}
