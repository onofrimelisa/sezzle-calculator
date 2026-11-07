#!/usr/bin/env bash
set -euo pipefail
OP_NAME="${1:?Usage: propose.sh <operation_name>}"
SPEC_DIR="spec/operations"
SPEC_FILE="${SPEC_DIR}/${OP_NAME}.md"
if [ -f "$SPEC_FILE" ]; then echo "Error: Spec already exists at ${SPEC_FILE}"; exit 1; fi
mkdir -p "$SPEC_DIR"

cat > "$SPEC_FILE" << TEMPLATE
# Operation: ${OP_NAME}

## Description
Brief description of what this operation does.

## Category
Basic | Advanced

## Type
Binary | Unary

## Inputs
- operand_a: float64 (required) — first operand
- operand_b: float64 (required|optional) — second operand (remove if unary)

## Output
- result: float64 — description of the result

## Edge Cases
- [ ] Describe edge case → EXPECTED_ERROR_CODE

## Examples
| operand_a | operand_b | result | notes |
|-----------|-----------|--------|-------|
| 5         | 3         | ?      |       |

## Test Cases
- [ ] Happy path with positive numbers
- [ ] Happy path with negative numbers
- [ ] Happy path with zero
- [ ] Happy path with decimals

## Swagger
- Tags: Basic Operations | Advanced Operations
- Summary: one-line summary
- Error responses: list applicable error codes
TEMPLATE

echo "Spec created at ${SPEC_FILE}"
echo "Edit the spec, then run: make analyze op=${OP_NAME}"
