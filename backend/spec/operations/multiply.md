# Operation: multiply

## Description
Performs multiplication of two operands.

## Category
Basic Operations

## Type
Binary

## Inputs
- operand_a: float64 (required) -- first factor
- operand_b: float64 (required) -- second factor

## Output
- result: float64 -- product of operand_a and operand_b (operand_a * operand_b)

## Edge Cases
No domain-specific edge cases. All float64 inputs are valid (multiplying by zero returns zero per IEEE 754).

## Examples
| operand_a | operand_b | result | notes |
|-----------|-----------|--------|-------|
| 6         | 7         | 42     |       |
| -3        | -4        | 12     | negative times negative |
| -3        | 4         | -12    | mixed signs |
| 0         | 100       | 0      | multiply by zero |
| 1.5       | 2.0       | 3.0    | decimals |
| 1         | 999       | 999    | identity |

## Test Cases
- [x] Happy path with positive numbers
- [x] Happy path with negative numbers
- [x] Mixed positive and negative
- [x] Multiply by zero
- [x] Decimal numbers
- [x] Identity (multiply by 1)
- [x] Large numbers

## Swagger
- Tags: Basic Operations
- Summary: Multiplication of two operands
- Error responses: INVALID_INPUT, MISSING_OPERAND
