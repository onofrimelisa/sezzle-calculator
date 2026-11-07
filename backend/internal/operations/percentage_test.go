package operations_test

import (
	"testing"

	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/sezzle/calculator-backend/internal/operations"
	"github.com/stretchr/testify/assert"
)

func TestPercentage_Execute(t *testing.T) {
	op := operations.NewPercentage()

	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"15 percent of 200", 200, 15, 30},
		{"50 percent of 100", 100, 50, 50},
		{"100 percent of 50", 50, 100, 50},
		{"0 percent of 100", 100, 0, 0},
		{"200 percent of 50", 50, 200, 100},
		{"decimal percentage", 100, 33.33, 33.33},
		{"negative base", -200, 15, -30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := op.Execute(tt.a, tt.b)
			assert.NoError(t, err)
			assert.InDelta(t, tt.expected, result, 1e-9)
		})
	}
}

func TestPercentage_Metadata(t *testing.T) {
	var op domain.BinaryOperation = operations.NewPercentage()
	assert.Equal(t, "percentage", op.Name())
	assert.Equal(t, "Advanced Operations", op.Category())
	assert.NotEmpty(t, op.Description())
}
