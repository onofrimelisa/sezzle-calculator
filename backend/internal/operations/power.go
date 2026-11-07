package operations

import "math"

// Power implements domain.BinaryOperation for exponentiation.
type Power struct{}

func NewPower() *Power { return &Power{} }

func (p *Power) Execute(operandA, operandB float64) (float64, error) {
	return math.Pow(operandA, operandB), nil
}

func (p *Power) Name() string        { return "power" }
func (p *Power) Description() string { return "Exponentiation: operand_a ^ operand_b" }
func (p *Power) Category() string    { return "Advanced Operations" }
