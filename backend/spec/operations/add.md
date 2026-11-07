# Operation: add

## Description
Performs addition of two operands.

## Category
Basic Operations

## Type
Binary

## Inputs
- operand_a: float64 (required) -- first operand
- operand_b: float64 (required) -- second operand

## Output
- result: float64 -- sum of operand_a and operand_b

## Edge Cases
No domain-specific edge cases. All float64 inputs are valid.

## Examples
| operand_a | operand_b | result | notes |
|-----------|-----------|--------|-------|
| 5         | 3         | 8      |       |
| -5        | -3        | -8     | negative numbers |
| -5        | 3         | -2     | mixed signs |
| 0         | 0         | 0      | zeros |
| 1.5       | 2.3       | 3.8    | decimals |

## Test Cases
- [x] Happy path with positive numbers
- [x] Happy path with negative numbers
- [x] Mixed positive and negative
- [x] Both zeros
- [x] Decimal numbers
- [x] Large numbers

## Swagger
- Tags: Basic Operations
- Summary: Addition of two operands
- Error responses: INVALID_INPUT, MISSING_OPERAND
