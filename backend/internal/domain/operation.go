// backend/internal/domain/operation.go
package domain

type OperationType string

const (
	Binary OperationType = "binary"
	Unary  OperationType = "unary"
)

type OperationMeta interface {
	Name() string
	Description() string
	Category() string
}

type BinaryOperation interface {
	OperationMeta
	Execute(a, b float64) (float64, error)
}

type UnaryOperation interface {
	OperationMeta
	Execute(a float64) (float64, error)
}
