# Operation: percentage

## Description
Computes operand_b percent of operand_a. The formula is `operand_a * (operand_b / 100)`. For example, 15% of 200 is calculated as `200 * (15 / 100) = 30`.

## Category
Advanced Operations

## Type
Binary

## Inputs
- operand_a: float64 (required) -- the base value to calculate the percentage of
- operand_b: float64 (required) -- the percentage rate

## Output
- result: float64 -- operand_b percent of operand_a, computed as `operand_a * (operand_b / 100)`

## Edge Cases
No domain-specific edge cases. All float64 inputs are valid (0% returns 0, 100% returns operand_a, values above 100% are valid).

## Examples
| operand_a | operand_b | result | notes |
|-----------|-----------|--------|-------|
| 200       | 15        | 30     | 15% of 200 |
| 100       | 50        | 50     | 50% of 100 |
| 100       | 100       | 100    | 100% returns the original value |
| 100       | 0         | 0      | 0% is always 0 |
| 50        | 200       | 100    | percentage above 100% |
| -100      | 10        | -10    | negative base |
| 200       | -15       | -30    | negative percentage |

## Test Cases
- [x] Happy path with standard percentage
- [x] 100% returns original value
- [x] 0% returns zero
- [x] Percentage above 100%
- [x] Negative base value
- [x] Negative percentage rate
- [x] Decimal percentage
- [x] Both zeros

## Swagger
- Tags: Advanced Operations
- Summary: Percentage calculation (operand_b% of operand_a)
- Error responses: INVALID_INPUT, MISSING_OPERAND
