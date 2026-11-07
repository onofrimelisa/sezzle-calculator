package middleware_test

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sezzle/calculator-backend/internal/middleware"
	"github.com/stretchr/testify/assert"
)

func TestLogging_LogsRequest(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))
	handler := middleware.Logging(logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	req := httptest.NewRequest(http.MethodPost, "/api/calculate", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	output := buf.String()
	assert.Contains(t, output, "/api/calculate")
	assert.Contains(t, output, "POST")
}
