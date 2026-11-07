package operations

// Multiply implements domain.BinaryOperation for multiplication.
type Multiply struct{}

func NewMultiply() *Multiply { return &Multiply{} }

func (m *Multiply) Execute(operandA, operandB float64) (float64, error) {
	return operandA * operandB, nil
}

func (m *Multiply) Name() string        { return "multiply" }
func (m *Multiply) Description() string { return "Multiplication: operand_a * operand_b" }
func (m *Multiply) Category() string    { return "Basic Operations" }
