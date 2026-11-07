#!/usr/bin/env bash
set -euo pipefail
OP_NAME="${1:?Usage: apply.sh <operation_name>}"
SPEC_FILE="spec/operations/${OP_NAME}.md"
if [ ! -f "$SPEC_FILE" ]; then echo "Error: Spec not found. Run: make propose op=${OP_NAME}"; exit 1; fi

OP_FILE="internal/operations/${OP_NAME}.go"
TEST_FILE="internal/operations/${OP_NAME}_test.go"
if [ -f "$OP_FILE" ]; then echo "Warning: ${OP_FILE} already exists. Skipping scaffold."; exit 0; fi

OP_TYPE=$(grep -A1 "## Type" "$SPEC_FILE" | tail -1 | tr -d '[:space:]')
STRUCT_NAME="$(echo "${OP_NAME}" | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) tolower(substr($i,2))}1' | tr -d ' ')"

echo "Scaffolding ${OP_NAME} (type: ${OP_TYPE})"

if [ "${OP_TYPE}" = "Binary" ]; then
cat > "$OP_FILE" << EOF
package operations

type ${STRUCT_NAME} struct{}
func New${STRUCT_NAME}() *${STRUCT_NAME} { return &${STRUCT_NAME}{} }
func (o *${STRUCT_NAME}) Execute(operandA, operandB float64) (float64, error) { return 0, nil }
func (o *${STRUCT_NAME}) Name() string        { return "${OP_NAME}" }
func (o *${STRUCT_NAME}) Description() string  { return "TODO: from spec" }
func (o *${STRUCT_NAME}) Category() string     { return "TODO" }
EOF
else
cat > "$OP_FILE" << EOF
package operations

type ${STRUCT_NAME} struct{}
func New${STRUCT_NAME}() *${STRUCT_NAME} { return &${STRUCT_NAME}{} }
func (o *${STRUCT_NAME}) Execute(operandA float64) (float64, error) { return 0, nil }
func (o *${STRUCT_NAME}) Name() string        { return "${OP_NAME}" }
func (o *${STRUCT_NAME}) Description() string  { return "TODO: from spec" }
func (o *${STRUCT_NAME}) Category() string     { return "TODO" }
EOF
fi

echo "Scaffolded: ${OP_FILE}. Implement Execute() and fill tests."
