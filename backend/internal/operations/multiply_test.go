package operations_test

import (
	"testing"

	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/sezzle/calculator-backend/internal/operations"
	"github.com/stretchr/testify/assert"
)

func TestMultiply_Execute(t *testing.T) {
	op := operations.NewMultiply()

	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"positive numbers", 5, 3, 15},
		{"negative times negative", -5, -3, 15},
		{"negative times positive", -5, 3, -15},
		{"multiply by zero", 5, 0, 0},
		{"multiply by one", 5, 1, 5},
		{"decimals", 2.5, 4, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := op.Execute(tt.a, tt.b)
			assert.NoError(t, err)
			assert.InDelta(t, tt.expected, result, 1e-9)
		})
	}
}

func TestMultiply_Metadata(t *testing.T) {
	var op domain.BinaryOperation = operations.NewMultiply()
	assert.Equal(t, "multiply", op.Name())
	assert.Equal(t, "Basic Operations", op.Category())
	assert.NotEmpty(t, op.Description())
}
