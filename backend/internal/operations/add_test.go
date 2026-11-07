package operations_test

import (
	"testing"

	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/sezzle/calculator-backend/internal/operations"
	"github.com/stretchr/testify/assert"
)

func TestAdd_Execute(t *testing.T) {
	op := operations.NewAdd()

	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -5, -3, -8},
		{"mixed sign", -5, 3, -2},
		{"both zeros", 0, 0, 0},
		{"one zero", 5, 0, 5},
		{"decimals", 1.5, 2.3, 3.8},
		{"large numbers", 1e15, 2e15, 3e15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := op.Execute(tt.a, tt.b)
			assert.NoError(t, err)
			assert.InDelta(t, tt.expected, result, 1e-9)
		})
	}
}

func TestAdd_Metadata(t *testing.T) {
	var op domain.BinaryOperation = operations.NewAdd()
	assert.Equal(t, "add", op.Name())
	assert.Equal(t, "Basic Operations", op.Category())
	assert.NotEmpty(t, op.Description())
}
