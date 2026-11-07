// backend/internal/domain/operation_test.go
package domain_test

import (
	"testing"

	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

type mockBinaryOp struct{}
func (m *mockBinaryOp) Execute(a, b float64) (float64, error) { return a + b, nil }
func (m *mockBinaryOp) Name() string        { return "mock_binary" }
func (m *mockBinaryOp) Description() string  { return "A mock binary operation" }
func (m *mockBinaryOp) Category() string     { return "test" }

type mockUnaryOp struct{}
func (m *mockUnaryOp) Execute(a float64) (float64, error) { return a * 2, nil }
func (m *mockUnaryOp) Name() string        { return "mock_unary" }
func (m *mockUnaryOp) Description() string  { return "A mock unary operation" }
func (m *mockUnaryOp) Category() string     { return "test" }

func TestBinaryOperation_Interface(t *testing.T) {
	var op domain.BinaryOperation = &mockBinaryOp{}
	result, err := op.Execute(5, 3)
	assert.NoError(t, err)
	assert.Equal(t, 8.0, result)
	assert.Equal(t, "mock_binary", op.Name())
	assert.Equal(t, "A mock binary operation", op.Description())
	assert.Equal(t, "test", op.Category())
}

func TestUnaryOperation_Interface(t *testing.T) {
	var op domain.UnaryOperation = &mockUnaryOp{}
	result, err := op.Execute(5)
	assert.NoError(t, err)
	assert.Equal(t, 10.0, result)
	assert.Equal(t, "mock_unary", op.Name())
	assert.Equal(t, "A mock unary operation", op.Description())
	assert.Equal(t, "test", op.Category())
}

func TestOperationType_Constants(t *testing.T) {
	assert.Equal(t, domain.OperationType("binary"), domain.Binary)
	assert.Equal(t, domain.OperationType("unary"), domain.Unary)
}
