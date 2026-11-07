package operations_test

import (
	"testing"

	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/sezzle/calculator-backend/internal/operations"
	"github.com/stretchr/testify/assert"
)

func TestSqrt_Execute(t *testing.T) {
	op := operations.NewSqrt()

	tests := []struct {
		name     string
		a        float64
		expected float64
	}{
		{"perfect square", 9, 3},
		{"irrational result", 2, 1.4142135623730951},
		{"zero", 0, 0},
		{"one", 1, 1},
		{"large number", 10000, 100},
		{"decimal", 0.25, 0.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := op.Execute(tt.a)
			assert.NoError(t, err)
			assert.InDelta(t, tt.expected, result, 1e-9)
		})
	}
}

func TestSqrt_Execute_NegativeInput(t *testing.T) {
	op := operations.NewSqrt()

	result, err := op.Execute(-4)
	assert.ErrorIs(t, err, domain.ErrNegativeSqrt)
	assert.Equal(t, float64(0), result)
}

func TestSqrt_Metadata(t *testing.T) {
	var op domain.UnaryOperation = operations.NewSqrt()
	assert.Equal(t, "sqrt", op.Name())
	assert.Equal(t, "Advanced Operations", op.Category())
	assert.NotEmpty(t, op.Description())
}
