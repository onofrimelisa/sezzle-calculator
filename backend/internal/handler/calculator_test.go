package handler_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/sezzle/calculator-backend/internal/dto"
	"github.com/sezzle/calculator-backend/internal/handler"
	"github.com/sezzle/calculator-backend/internal/operations"
	"github.com/sezzle/calculator-backend/internal/platform/metrics"
	"github.com/sezzle/calculator-backend/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testMetrics = metrics.New()

func newTestHandler() *handler.CalculatorHandler {
	registry := domain.NewRegistry()
	registry.RegisterBinary(operations.NewAdd())
	registry.RegisterBinary(operations.NewSubtract())
	registry.RegisterBinary(operations.NewDivide())
	registry.RegisterUnary(operations.NewSqrt())
	logger := slog.Default()
	svc := service.NewCalculatorService(registry, logger, testMetrics)
	v := validator.New()
	return handler.NewCalculatorHandler(svc, v, logger)
}

func ptrF(f float64) *float64 { return &f }

func TestCalculateHandler_Success(t *testing.T) {
	h := newTestHandler()
	body := dto.CalculateRequest{Operation: "add", OperandA: 5, OperandB: ptrF(3)}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.Calculate(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	var resp dto.CalculateResponse
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, 8.0, resp.Result)
}

func TestCalculateHandler_InvalidJSON(t *testing.T) {
	h := newTestHandler()
	req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.Calculate(rec, req)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	var resp dto.ErrorResponse
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.Equal(t, "INVALID_INPUT", resp.Error.Code)
}

func TestCalculateHandler_DivisionByZero(t *testing.T) {
	h := newTestHandler()
	body := dto.CalculateRequest{Operation: "divide", OperandA: 10, OperandB: ptrF(0)}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.Calculate(rec, req)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	var resp dto.ErrorResponse
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.Equal(t, "DIVISION_BY_ZERO", resp.Error.Code)
}

func TestCalculateHandler_InvalidOperation(t *testing.T) {
	h := newTestHandler()
	body := dto.CalculateRequest{Operation: "nonexistent", OperandA: 5, OperandB: ptrF(3)}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.Calculate(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var resp dto.ErrorResponse
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.Equal(t, "INVALID_OPERATION", resp.Error.Code)
}

func TestCalculateHandler_MissingOperation(t *testing.T) {
	h := newTestHandler()
	body := map[string]interface{}{"operand_a": 5}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.Calculate(rec, req)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}
