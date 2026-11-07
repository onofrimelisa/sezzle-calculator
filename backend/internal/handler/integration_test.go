package handler_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/sezzle/calculator-backend/internal/dto"
	"github.com/sezzle/calculator-backend/internal/handler"
	"github.com/sezzle/calculator-backend/internal/middleware"
	"github.com/sezzle/calculator-backend/internal/operations"
	"github.com/sezzle/calculator-backend/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupRouter() *chi.Mux {
	registry := domain.NewRegistry()
	registry.RegisterBinary(operations.NewAdd())
	registry.RegisterBinary(operations.NewSubtract())
	registry.RegisterBinary(operations.NewMultiply())
	registry.RegisterBinary(operations.NewDivide())
	registry.RegisterBinary(operations.NewPower())
	registry.RegisterUnary(operations.NewSqrt())
	registry.RegisterBinary(operations.NewPercentage())

	logger := slog.Default()
	svc := service.NewCalculatorService(registry, logger, testMetrics)
	v := validator.New()
	h := handler.NewCalculatorHandler(svc, v, logger)

	r := chi.NewRouter()
	r.Use(middleware.Recovery(logger))
	r.Use(middleware.RequestID)
	r.Post("/api/calculate", h.Calculate)
	r.Get("/healthz", handler.HealthCheck)
	return r
}

func TestIntegration_AllOperations(t *testing.T) {
	router := setupRouter()
	tests := []struct {
		name      string
		operation string
		a         float64
		b         *float64
		expected  float64
	}{
		{"add", "add", 10, ptrF(5), 15},
		{"subtract", "subtract", 10, ptrF(5), 5},
		{"multiply", "multiply", 4, ptrF(3), 12},
		{"divide", "divide", 10, ptrF(4), 2.5},
		{"power", "power", 2, ptrF(8), 256},
		{"sqrt", "sqrt", 16, nil, 4},
		{"percentage", "percentage", 200, ptrF(15), 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := dto.CalculateRequest{Operation: tt.operation, OperandA: tt.a, OperandB: tt.b}
			b, _ := json.Marshal(body)
			req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader(b))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
			var resp dto.CalculateResponse
			err := json.Unmarshal(rec.Body.Bytes(), &resp)
			require.NoError(t, err)
			assert.InDelta(t, tt.expected, resp.Result, 1e-9)
		})
	}
}

func TestIntegration_ErrorCases(t *testing.T) {
	router := setupRouter()
	tests := []struct {
		name           string
		body           interface{}
		expectedStatus int
		expectedCode   string
	}{
		{"division by zero", dto.CalculateRequest{Operation: "divide", OperandA: 10, OperandB: ptrF(0)}, http.StatusUnprocessableEntity, "DIVISION_BY_ZERO"},
		{"negative sqrt", dto.CalculateRequest{Operation: "sqrt", OperandA: -4}, http.StatusUnprocessableEntity, "NEGATIVE_SQRT"},
		{"invalid operation", dto.CalculateRequest{Operation: "modulo", OperandA: 5, OperandB: ptrF(3)}, http.StatusBadRequest, "INVALID_OPERATION"},
		{"missing operand for binary", dto.CalculateRequest{Operation: "add", OperandA: 5}, http.StatusUnprocessableEntity, "MISSING_OPERAND"},
		{"unexpected operand for unary", dto.CalculateRequest{Operation: "sqrt", OperandA: 9, OperandB: ptrF(3)}, http.StatusUnprocessableEntity, "UNEXPECTED_OPERAND"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.body)
			req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader(b))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedStatus, rec.Code)
			var resp dto.ErrorResponse
			err := json.Unmarshal(rec.Body.Bytes(), &resp)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedCode, resp.Error.Code)
		})
	}
}

func TestIntegration_RequestID_Propagated(t *testing.T) {
	router := setupRouter()
	body := dto.CalculateRequest{Operation: "add", OperandA: 1, OperandB: ptrF(1)}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.NotEmpty(t, rec.Header().Get("X-Request-ID"))
}

func TestIntegration_HealthCheck(t *testing.T) {
	router := setupRouter()
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}
