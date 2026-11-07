package service_test

import (
	"log/slog"
	"testing"

	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/sezzle/calculator-backend/internal/dto"
	"github.com/sezzle/calculator-backend/internal/operations"
	"github.com/sezzle/calculator-backend/internal/platform/metrics"
	"github.com/sezzle/calculator-backend/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testMetrics = metrics.New()

func newTestService() *service.CalculatorService {
	registry := domain.NewRegistry()
	registry.RegisterBinary(operations.NewAdd())
	registry.RegisterBinary(operations.NewDivide())
	registry.RegisterUnary(operations.NewSqrt())
	logger := slog.Default()
	return service.NewCalculatorService(registry, logger, testMetrics)
}

func ptrFloat(f float64) *float64 { return &f }

func TestCalculatorService_BinaryOperation(t *testing.T) {
	svc := newTestService()
	req := dto.CalculateRequest{Operation: "add", OperandA: 5, OperandB: ptrFloat(3)}
	resp, err := svc.Calculate(req)
	require.NoError(t, err)
	assert.Equal(t, 8.0, resp.Result)
	assert.Equal(t, "add", resp.Operation)
}

func TestCalculatorService_UnaryOperation(t *testing.T) {
	svc := newTestService()
	req := dto.CalculateRequest{Operation: "sqrt", OperandA: 9}
	resp, err := svc.Calculate(req)
	require.NoError(t, err)
	assert.InDelta(t, 3.0, resp.Result, 1e-9)
}

func TestCalculatorService_InvalidOperation(t *testing.T) {
	svc := newTestService()
	req := dto.CalculateRequest{Operation: "nonexistent", OperandA: 5, OperandB: ptrFloat(3)}
	_, err := svc.Calculate(req)
	assert.ErrorIs(t, err, domain.ErrInvalidOperation)
}

func TestCalculatorService_MissingOperandForBinary(t *testing.T) {
	svc := newTestService()
	req := dto.CalculateRequest{Operation: "add", OperandA: 5}
	_, err := svc.Calculate(req)
	assert.ErrorIs(t, err, domain.ErrMissingOperand)
}

func TestCalculatorService_UnexpectedOperandForUnary(t *testing.T) {
	svc := newTestService()
	req := dto.CalculateRequest{Operation: "sqrt", OperandA: 9, OperandB: ptrFloat(3)}
	_, err := svc.Calculate(req)
	assert.ErrorIs(t, err, domain.ErrUnexpectedOperand)
}

func TestCalculatorService_DomainError(t *testing.T) {
	svc := newTestService()
	req := dto.CalculateRequest{Operation: "divide", OperandA: 10, OperandB: ptrFloat(0)}
	_, err := svc.Calculate(req)
	assert.ErrorIs(t, err, domain.ErrDivisionByZero)
}
