package operations

// Add implements domain.BinaryOperation for addition.
type Add struct{}

func NewAdd() *Add { return &Add{} }

func (a *Add) Execute(operandA, operandB float64) (float64, error) {
	return operandA + operandB, nil
}

func (a *Add) Name() string        { return "add" }
func (a *Add) Description() string { return "Addition: operand_a + operand_b" }
func (a *Add) Category() string    { return "Basic Operations" }
