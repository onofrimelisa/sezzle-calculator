package operations

// Percentage implements domain.BinaryOperation for percentage calculation.
type Percentage struct{}

func NewPercentage() *Percentage { return &Percentage{} }

func (p *Percentage) Execute(operandA, operandB float64) (float64, error) {
	return operandA * (operandB / 100), nil
}

func (p *Percentage) Name() string        { return "percentage" }
func (p *Percentage) Description() string { return "Percentage: operand_b% of operand_a" }
func (p *Percentage) Category() string    { return "Advanced Operations" }
