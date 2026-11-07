#!/usr/bin/env bash
set -euo pipefail
OP_NAME="${1:?Usage: verify.sh <operation_name>}"
OP_FILE="internal/operations/${OP_NAME}.go"
if [ ! -f "$OP_FILE" ]; then echo "Error: Implementation not found at ${OP_FILE}"; exit 1; fi

echo "=== Verifying operation: ${OP_NAME} ==="
echo "--- Lint ---"
go vet ./internal/operations/
echo "--- Tests ---"
STRUCT_NAME="$(echo "${OP_NAME}" | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) tolower(substr($i,2))}1' | tr -d ' ')"
go test ./internal/operations/ -v -run "Test${STRUCT_NAME}"
echo "=== Verification complete for: ${OP_NAME} ==="
