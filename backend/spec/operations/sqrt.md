# Operation: sqrt

## Description
Computes the square root of operand_a using Go's math.Sqrt function. Returns a NEGATIVE_SQRT error when operand_a is negative.

## Category
Advanced Operations

## Type
Unary

## Inputs
- operand_a: float64 (required) -- the value to compute the square root of

## Output
- result: float64 -- square root of operand_a

## Edge Cases
- operand_a < 0 --> NEGATIVE_SQRT: Cannot calculate square root of a negative number

## Examples
| operand_a | result | notes |
|-----------|--------|-------|
| 9         | 3      |       |
| 144       | 12     |       |
| 0         | 0      | square root of zero |
| 2         | 1.4142 | irrational result |
| 0.25      | 0.5    | fractional input |
| -4        | error  | NEGATIVE_SQRT |

## Test Cases
- [x] Happy path with perfect square
- [x] Happy path with large perfect square
- [x] Square root of zero
- [x] Irrational result (non-perfect square)
- [x] Fractional input
- [x] Negative input returns NEGATIVE_SQRT error

## Swagger
- Tags: Advanced Operations
- Summary: Square root of operand_a
- Error responses: INVALID_INPUT, UNEXPECTED_OPERAND, NEGATIVE_SQRT
