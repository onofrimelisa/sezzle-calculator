// backend/internal/domain/registry_test.go
package domain_test

import (
	"testing"

	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistry_RegisterBinary(t *testing.T) {
	registry := domain.NewRegistry()
	op := &mockBinaryOp{}
	registry.RegisterBinary(op)
	entry, err := registry.Get("mock_binary")
	require.NoError(t, err)
	assert.Equal(t, domain.Binary, entry.Type)
	assert.NotNil(t, entry.Binary)
	assert.Nil(t, entry.Unary)
}

func TestRegistry_RegisterUnary(t *testing.T) {
	registry := domain.NewRegistry()
	op := &mockUnaryOp{}
	registry.RegisterUnary(op)
	entry, err := registry.Get("mock_unary")
	require.NoError(t, err)
	assert.Equal(t, domain.Unary, entry.Type)
	assert.NotNil(t, entry.Unary)
	assert.Nil(t, entry.Binary)
}

func TestRegistry_Get_NotFound(t *testing.T) {
	registry := domain.NewRegistry()
	_, err := registry.Get("nonexistent")
	assert.ErrorIs(t, err, domain.ErrInvalidOperation)
}

func TestRegistry_List(t *testing.T) {
	registry := domain.NewRegistry()
	registry.RegisterBinary(&mockBinaryOp{})
	registry.RegisterUnary(&mockUnaryOp{})
	entries := registry.List()
	assert.Len(t, entries, 2)
}
