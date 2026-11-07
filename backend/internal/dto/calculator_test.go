package dto_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sezzle/calculator-backend/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestCalculateRequest_Valid_Binary(t *testing.T) {
	v := validator.New()
	b := 3.0
	req := dto.CalculateRequest{Operation: "add", OperandA: 5.0, OperandB: &b}
	err := v.Struct(req)
	assert.NoError(t, err)
}

func TestCalculateRequest_Valid_Unary(t *testing.T) {
	v := validator.New()
	req := dto.CalculateRequest{Operation: "sqrt", OperandA: 9.0}
	err := v.Struct(req)
	assert.NoError(t, err)
}

func TestCalculateRequest_Invalid_MissingOperation(t *testing.T) {
	v := validator.New()
	req := dto.CalculateRequest{OperandA: 5.0}
	err := v.Struct(req)
	assert.Error(t, err)
}

func TestErrorResponse_Structure(t *testing.T) {
	resp := dto.NewErrorResponse("DIVISION_BY_ZERO", "Cannot divide by zero", map[string]interface{}{
		"operation": "divide", "operand_a": 10.0, "operand_b": 0.0,
	})
	assert.Equal(t, "DIVISION_BY_ZERO", resp.Error.Code)
	assert.Equal(t, "Cannot divide by zero", resp.Error.Message)
	assert.Equal(t, "divide", resp.Error.Details["operation"])
}
