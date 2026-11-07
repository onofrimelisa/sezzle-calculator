package operations

// Subtract implements domain.BinaryOperation for subtraction.
type Subtract struct{}

func NewSubtract() *Subtract { return &Subtract{} }

func (s *Subtract) Execute(operandA, operandB float64) (float64, error) {
	return operandA - operandB, nil
}

func (s *Subtract) Name() string        { return "subtract" }
func (s *Subtract) Description() string { return "Subtraction: operand_a - operand_b" }
func (s *Subtract) Category() string    { return "Basic Operations" }
