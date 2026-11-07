#!/usr/bin/env bash
set -euo pipefail
echo "Regenerating Swagger docs..."
swag init -g cmd/server/main.go -o docs 2>/dev/null && echo "Swagger docs regenerated." || echo "Warning: swag not installed, skipping."

echo ""
echo "Checking spec-code alignment..."
SPECS_DIR="spec/operations"
OPS_DIR="internal/operations"
ERRORS=0

for spec in "${SPECS_DIR}"/*.md; do
    [ -f "$spec" ] || continue
    op_name=$(basename "$spec" .md)
    if [ ! -f "${OPS_DIR}/${op_name}.go" ]; then
        echo "MISSING IMPLEMENTATION: ${op_name}"
        ERRORS=$((ERRORS + 1))
    fi
done

if [ "$ERRORS" -gt 0 ]; then echo "SYNC FAILED: ${ERRORS} spec(s) without implementation"; exit 1
else echo "SYNC OK: All specs have implementations"; fi
