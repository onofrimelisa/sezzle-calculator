package logger_test

import (
	"bytes"
	"testing"

	"github.com/sezzle/calculator-backend/internal/platform/logger"
	"github.com/stretchr/testify/assert"
)

func TestNew_DefaultLevel(t *testing.T) {
	var buf bytes.Buffer
	l := logger.New("info", &buf)
	l.Info("test message", "key", "value")
	assert.Contains(t, buf.String(), "test message")
	assert.Contains(t, buf.String(), "key")
}

func TestNew_DebugLevel(t *testing.T) {
	var buf bytes.Buffer
	l := logger.New("debug", &buf)
	l.Debug("debug message")
	assert.Contains(t, buf.String(), "debug message")
}

func TestNew_DebugSuppressedAtInfoLevel(t *testing.T) {
	var buf bytes.Buffer
	l := logger.New("info", &buf)
	l.Debug("should not appear")
	assert.Empty(t, buf.String())
}
