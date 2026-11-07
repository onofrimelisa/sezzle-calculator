package operations

import (
	"math"

	"github.com/sezzle/calculator-backend/internal/domain"
)

// Sqrt implements domain.UnaryOperation for square root.
type Sqrt struct{}

func NewSqrt() *Sqrt { return &Sqrt{} }

func (s *Sqrt) Execute(operandA float64) (float64, error) {
	if operandA < 0 {
		return 0, domain.ErrNegativeSqrt
	}
	return math.Sqrt(operandA), nil
}

func (s *Sqrt) Name() string        { return "sqrt" }
func (s *Sqrt) Description() string { return "Square root of operand_a" }
func (s *Sqrt) Category() string    { return "Advanced Operations" }
