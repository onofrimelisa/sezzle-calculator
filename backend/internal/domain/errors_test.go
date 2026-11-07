// backend/internal/domain/errors_test.go
package domain_test

import (
	"errors"
	"testing"

	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestDomainError_Error(t *testing.T) {
	err := domain.NewDomainError("DIVISION_BY_ZERO", "Cannot divide by zero")
	assert.Equal(t, "DIVISION_BY_ZERO: Cannot divide by zero", err.Error())
}

func TestDomainError_Is(t *testing.T) {
	err := domain.NewDomainError("DIVISION_BY_ZERO", "Cannot divide by zero")
	var domErr *domain.DomainError
	assert.True(t, errors.As(err, &domErr))
	assert.Equal(t, "DIVISION_BY_ZERO", domErr.Code)
}

func TestPredefinedErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		code string
	}{
		{"DivisionByZero", domain.ErrDivisionByZero, "DIVISION_BY_ZERO"},
		{"NegativeSqrt", domain.ErrNegativeSqrt, "NEGATIVE_SQRT"},
		{"InvalidOperation", domain.ErrInvalidOperation, "INVALID_OPERATION"},
		{"InvalidInput", domain.ErrInvalidInput, "INVALID_INPUT"},
		{"MissingOperand", domain.ErrMissingOperand, "MISSING_OPERAND"},
		{"UnexpectedOperand", domain.ErrUnexpectedOperand, "UNEXPECTED_OPERAND"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var domErr *domain.DomainError
			assert.True(t, errors.As(tt.err, &domErr))
			assert.Equal(t, tt.code, domErr.Code)
		})
	}
}
