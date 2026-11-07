# Operation: divide

## Description
Performs division of two operands (operand_a divided by operand_b). Returns a DIVISION_BY_ZERO error when operand_b is zero.

## Category
Basic Operations

## Type
Binary

## Inputs
- operand_a: float64 (required) -- dividend
- operand_b: float64 (required) -- divisor

## Output
- result: float64 -- quotient of operand_a and operand_b (operand_a / operand_b)

## Edge Cases
- operand_b == 0 --> DIVISION_BY_ZERO: Cannot divide by zero

## Examples
| operand_a | operand_b | result | notes |
|-----------|-----------|--------|-------|
| 20        | 4         | 5      |       |
| -12       | -3        | 4      | negative divided by negative |
| -12       | 3         | -4     | mixed signs |
| 0         | 5         | 0      | zero dividend |
| 7         | 2         | 3.5    | non-integer result |
| 10        | 0         | error  | DIVISION_BY_ZERO |
| 1         | 3         | 0.3333 | repeating decimal |

## Test Cases
- [x] Happy path with positive numbers
- [x] Happy path with negative numbers
- [x] Mixed positive and negative
- [x] Zero dividend (0 / n = 0)
- [x] Non-integer result
- [x] Division by zero returns DIVISION_BY_ZERO error
- [x] Decimal numbers
- [x] Large numbers

## Swagger
- Tags: Basic Operations
- Summary: Division of two operands
- Error responses: INVALID_INPUT, MISSING_OPERAND, DIVISION_BY_ZERO
