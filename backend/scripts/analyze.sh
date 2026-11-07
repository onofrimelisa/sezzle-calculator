#!/usr/bin/env bash
set -euo pipefail
OP_NAME="${1:?Usage: analyze.sh <operation_name>}"
SPEC_FILE="spec/operations/${OP_NAME}.md"
if [ ! -f "$SPEC_FILE" ]; then echo "Error: Spec not found at ${SPEC_FILE}"; exit 1; fi

ERRORS=0
check_section() {
    if ! grep -q "## ${1}" "$SPEC_FILE"; then echo "MISSING: ## ${1} section"; ERRORS=$((ERRORS + 1)); fi
}

echo "Analyzing spec: ${SPEC_FILE}"
echo "---"
for section in Description Category Type Inputs Output "Edge Cases" Examples "Test Cases" Swagger; do
    check_section "$section"
done

if grep -q "Brief description" "$SPEC_FILE"; then echo "INCOMPLETE: Description still contains placeholder"; ERRORS=$((ERRORS + 1)); fi
if grep -q "one-line summary" "$SPEC_FILE"; then echo "INCOMPLETE: Swagger Summary placeholder"; ERRORS=$((ERRORS + 1)); fi

echo "---"
if [ "$ERRORS" -gt 0 ]; then echo "FAILED: ${ERRORS} issue(s) found"; exit 1; else echo "PASSED: Spec is complete. Run: make apply op=${OP_NAME}"; fi
