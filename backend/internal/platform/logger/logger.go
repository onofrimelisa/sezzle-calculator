package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

func New(level string, w ...io.Writer) *slog.Logger {
	var writer io.Writer = os.Stdout
	if len(w) > 0 {
		writer = w[0]
	}
	lvl := parseLevel(level)
	opts := &slog.HandlerOptions{Level: lvl}
	handler := slog.NewJSONHandler(writer, opts)
	return slog.New(handler)
}

func parseLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
