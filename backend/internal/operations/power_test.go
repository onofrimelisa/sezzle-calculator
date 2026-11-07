package operations_test

import (
	"testing"

	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/sezzle/calculator-backend/internal/operations"
	"github.com/stretchr/testify/assert"
)

func TestPower_Execute(t *testing.T) {
	op := operations.NewPower()

	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"square", 5, 2, 25},
		{"cube", 2, 3, 8},
		{"power of zero", 5, 0, 1},
		{"zero base", 0, 5, 0},
		{"negative exponent", 2, -1, 0.5},
		{"fractional exponent", 4, 0.5, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := op.Execute(tt.a, tt.b)
			assert.NoError(t, err)
			assert.InDelta(t, tt.expected, result, 1e-9)
		})
	}
}

func TestPower_Metadata(t *testing.T) {
	var op domain.BinaryOperation = operations.NewPower()
	assert.Equal(t, "power", op.Name())
	assert.Equal(t, "Advanced Operations", op.Category())
	assert.NotEmpty(t, op.Description())
}
