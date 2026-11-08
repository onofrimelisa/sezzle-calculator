# Calculator UX Improvements — Design Spec

## Overview

Improve calculator UX with two features: result carry-forward (result becomes next operand_a) and Clear/Reset buttons.

## Carry-Forward Behavior

- After a successful calculation, the result is loaded into operand_a
- operand_b is cleared
- Focus moves to operand_b (for binary operations) or stays on operand_a (for unary)
- If the calculation results in an error, no carry-forward happens — inputs stay as they are

## Clear and Reset Buttons

### C (Clear)
- Clears operand_a, operand_b, result, and error
- Preserves history and selected operation

### AC (All Clear)
- Clears operand_a, operand_b, result, error, and history
- Resets selected operation to "add"

### Placement
- Row of two buttons between the display/error area and the operand inputs

## Files Changed

| File | Change |
|---|---|
| `src/hooks/useCalculator.ts` | Add `clear()` and `reset()` methods |
| `src/components/Calculator.tsx` | Carry-forward logic in onSubmit, C/AC buttons, useRef for autofocus on operand_b |
| `src/components/Calculator.test.tsx` | Tests for carry-forward, clear, reset |
| `src/hooks/useCalculator.test.ts` | Tests for clear and reset methods |

## No Changes

Backend, API layer, Display, OperandInputs, OperationButtons, History — all unchanged.
