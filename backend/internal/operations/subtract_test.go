package operations_test

import (
	"testing"

	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/sezzle/calculator-backend/internal/operations"
	"github.com/stretchr/testify/assert"
)

func TestSubtract_Execute(t *testing.T) {
	op := operations.NewSubtract()

	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"positive numbers", 5, 3, 2},
		{"negative minus negative", -5, -3, -2},
		{"negative minus positive", -5, 3, -8},
		{"both zeros", 0, 0, 0},
		{"result is negative", 3, 5, -2},
		{"decimals", 5.5, 2.3, 3.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := op.Execute(tt.a, tt.b)
			assert.NoError(t, err)
			assert.InDelta(t, tt.expected, result, 1e-9)
		})
	}
}

func TestSubtract_Metadata(t *testing.T) {
	var op domain.BinaryOperation = operations.NewSubtract()
	assert.Equal(t, "subtract", op.Name())
	assert.Equal(t, "Basic Operations", op.Category())
	assert.NotEmpty(t, op.Description())
}
