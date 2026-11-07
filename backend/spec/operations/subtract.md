# Operation: subtract

## Description
Performs subtraction of two operands (operand_a minus operand_b).

## Category
Basic Operations

## Type
Binary

## Inputs
- operand_a: float64 (required) -- minuend
- operand_b: float64 (required) -- subtrahend

## Output
- result: float64 -- difference of operand_a and operand_b (operand_a - operand_b)

## Edge Cases
No domain-specific edge cases. All float64 inputs are valid.

## Examples
| operand_a | operand_b | result | notes |
|-----------|-----------|--------|-------|
| 10        | 4         | 6      |       |
| -5        | -3        | -2     | negative numbers |
| -5        | 3         | -8     | mixed signs |
| 0         | 0         | 0      | zeros |
| 5.5       | 2.3       | 3.2    | decimals |
| 3         | 10        | -7     | result is negative |

## Test Cases
- [x] Happy path with positive numbers
- [x] Happy path with negative numbers
- [x] Mixed positive and negative
- [x] Both zeros
- [x] Decimal numbers
- [x] Result is negative
- [x] Large numbers

## Swagger
- Tags: Basic Operations
- Summary: Subtraction of two operands
- Error responses: INVALID_INPUT, MISSING_OPERAND
