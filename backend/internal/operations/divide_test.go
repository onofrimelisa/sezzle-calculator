package operations_test

import (
	"testing"

	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/sezzle/calculator-backend/internal/operations"
	"github.com/stretchr/testify/assert"
)

func TestDivide_Execute(t *testing.T) {
	op := operations.NewDivide()

	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"positive numbers", 10, 2, 5},
		{"negative divided by negative", -10, -2, 5},
		{"negative divided by positive", -10, 2, -5},
		{"non-integer result", 10, 3, 3.333333333333333},
		{"divide by one", 7, 1, 7},
		{"decimals", 0.1, 0.2, 0.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := op.Execute(tt.a, tt.b)
			assert.NoError(t, err)
			assert.InDelta(t, tt.expected, result, 1e-9)
		})
	}
}

func TestDivide_Execute_DivisionByZero(t *testing.T) {
	op := operations.NewDivide()

	tests := []struct {
		name string
		a    float64
		b    float64
	}{
		{"nonzero divided by zero", 10, 0},
		{"zero divided by zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := op.Execute(tt.a, tt.b)
			assert.ErrorIs(t, err, domain.ErrDivisionByZero)
			assert.Equal(t, float64(0), result)
		})
	}
}

func TestDivide_Metadata(t *testing.T) {
	var op domain.BinaryOperation = operations.NewDivide()
	assert.Equal(t, "divide", op.Name())
	assert.Equal(t, "Basic Operations", op.Category())
	assert.NotEmpty(t, op.Description())
}
