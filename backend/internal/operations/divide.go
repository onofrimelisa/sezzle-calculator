package operations

import "github.com/sezzle/calculator-backend/internal/domain"

// Divide implements domain.BinaryOperation for division.
type Divide struct{}

func NewDivide() *Divide { return &Divide{} }

func (d *Divide) Execute(operandA, operandB float64) (float64, error) {
	if operandB == 0 {
		return 0, domain.ErrDivisionByZero
	}
	return operandA / operandB, nil
}

func (d *Divide) Name() string        { return "divide" }
func (d *Divide) Description() string { return "Division: operand_a / operand_b" }
func (d *Divide) Category() string    { return "Basic Operations" }
