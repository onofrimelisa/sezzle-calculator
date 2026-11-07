// backend/internal/domain/errors.go
package domain

import "fmt"

type DomainError struct {
	Code    string
	Message string
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewDomainError(code, message string) *DomainError {
	return &DomainError{Code: code, Message: message}
}

var (
	ErrDivisionByZero   = NewDomainError("DIVISION_BY_ZERO", "Cannot divide by zero")
	ErrNegativeSqrt     = NewDomainError("NEGATIVE_SQRT", "Cannot calculate square root of a negative number")
	ErrInvalidOperation = NewDomainError("INVALID_OPERATION", "The requested operation does not exist")
	ErrInvalidInput     = NewDomainError("INVALID_INPUT", "The provided input is invalid")
	ErrMissingOperand   = NewDomainError("MISSING_OPERAND", "Binary operations require both operand_a and operand_b")
	ErrUnexpectedOperand = NewDomainError("UNEXPECTED_OPERAND", "Unary operations only accept operand_a")
)
