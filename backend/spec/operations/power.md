# Operation: power

## Description
Performs exponentiation: raises operand_a to the power of operand_b using Go's math.Pow function.

## Category
Advanced Operations

## Type
Binary

## Inputs
- operand_a: float64 (required) -- base
- operand_b: float64 (required) -- exponent

## Output
- result: float64 -- operand_a raised to the power of operand_b (operand_a ^ operand_b)

## Edge Cases
No domain-specific edge cases. Go's math.Pow handles special values per IEEE 754 (e.g., 0^0 = 1, x^0 = 1, 0^y = 0 for y > 0).

## Examples
| operand_a | operand_b | result | notes |
|-----------|-----------|--------|-------|
| 2         | 10        | 1024   |       |
| 5         | 0         | 1      | any number to the power of 0 |
| 0         | 5         | 0      | 0 raised to positive power |
| 2         | -1        | 0.5    | negative exponent |
| 9         | 0.5       | 3      | fractional exponent (square root) |
| -2        | 3         | -8     | negative base, odd exponent |

## Test Cases
- [x] Happy path with positive base and exponent
- [x] Exponent is zero (result is 1)
- [x] Base is zero
- [x] Negative exponent
- [x] Fractional exponent
- [x] Negative base with integer exponent
- [x] Large exponent

## Swagger
- Tags: Advanced Operations
- Summary: Exponentiation of two operands
- Error responses: INVALID_INPUT, MISSING_OPERAND
