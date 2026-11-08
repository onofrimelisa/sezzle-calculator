# Calculator UX Improvements Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add result carry-forward (result becomes next operand_a) and Clear/Reset buttons to the calculator.

**Architecture:** Add `clear()` and `reset()` methods to `useCalculator` hook. Add carry-forward logic and C/AC button rendering in `Calculator.tsx`. Use `useRef` for auto-focus on operand_b after carry-forward.

**Tech Stack:** React 18, TypeScript, React Hook Form, Vitest, React Testing Library, msw

**Spec:** `docs/specs/2026-11-07-calculator-ux-improvements-design.md`

---

### File Structure

| File | Action | Responsibility |
|---|---|---|
| `src/hooks/useCalculator.ts` | Modify | Add `clear()` and `reset()` methods |
| `src/hooks/useCalculator.test.ts` | Modify | Tests for `clear()` and `reset()` |
| `src/components/Calculator.tsx` | Modify | Carry-forward in onSubmit, C/AC buttons, useRef for autofocus |
| `src/components/Calculator.test.tsx` | Modify | Tests for carry-forward, clear, reset |

---

### Task 1: Add `clear()` and `reset()` to useCalculator hook

**Files:**
- Modify: `frontend/src/hooks/useCalculator.ts`
- Modify: `frontend/src/hooks/useCalculator.test.ts`

- [ ] **Step 1: Write failing tests for `clear()` and `reset()`**

Add these tests to `frontend/src/hooks/useCalculator.test.ts`:

```typescript
it('clear() resets result and error but preserves operation', async () => {
  mockCalculate.mockResolvedValueOnce({
    operation: 'add', operand_a: 5, operand_b: 3, result: 8, timestamp: '2026-11-07T15:04:05Z',
  })
  const { result } = renderHook(() => useCalculator())
  await act(async () => { await result.current.performCalculation(5, 3) })
  expect(result.current.result).not.toBeNull()

  act(() => result.current.clear())
  expect(result.current.result).toBeNull()
  expect(result.current.error).toBeNull()
  expect(result.current.selectedOperation).toBe('add')
})

it('reset() resets everything including operation to add', async () => {
  mockCalculate.mockResolvedValueOnce({
    operation: 'multiply', operand_a: 2, operand_b: 3, result: 6, timestamp: '2026-11-07T15:04:05Z',
  })
  const { result } = renderHook(() => useCalculator())
  act(() => result.current.selectOperation('multiply'))
  await act(async () => { await result.current.performCalculation(2, 3) })

  act(() => result.current.reset())
  expect(result.current.result).toBeNull()
  expect(result.current.error).toBeNull()
  expect(result.current.selectedOperation).toBe('add')
})
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd frontend && npx vitest run src/hooks/useCalculator.test.ts`
Expected: FAIL — `clear` and `reset` are not functions on the hook return value.

- [ ] **Step 3: Implement `clear()` and `reset()` in useCalculator**

In `frontend/src/hooks/useCalculator.ts`, add two new callbacks and include them in the return value:

```typescript
const clear = useCallback(() => {
  setResult(null)
  setError(null)
}, [])

const reset = useCallback(() => {
  setResult(null)
  setError(null)
  setSelectedOperation('add')
}, [])

return { selectedOperation, selectOperation, result, error, isLoading, performCalculation, clear, reset }
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `cd frontend && npx vitest run src/hooks/useCalculator.test.ts`
Expected: All tests PASS.

- [ ] **Step 5: Commit**

```bash
git add frontend/src/hooks/useCalculator.ts frontend/src/hooks/useCalculator.test.ts
git commit -m "feat: add clear() and reset() methods to useCalculator hook"
```

---

### Task 2: Add C/AC buttons to Calculator component

**Files:**
- Modify: `frontend/src/components/Calculator.tsx`
- Modify: `frontend/src/components/Calculator.test.tsx`

- [ ] **Step 1: Write failing tests for C and AC buttons**

Add these tests to `frontend/src/components/Calculator.test.tsx`:

```typescript
it('C button clears inputs and result but preserves operation', async () => {
  render(<Calculator />)
  const user = userEvent.setup()
  await user.type(screen.getByLabelText('Operand A'), '5')
  await user.type(screen.getByLabelText('Operand B'), '3')
  await user.click(screen.getByText('Calculate'))
  await waitFor(() => { expect(screen.getByText('8')).toBeInTheDocument() })

  await user.click(screen.getByText('C'))

  expect(screen.getByLabelText('Operand A')).toHaveValue('')
  expect(screen.getByLabelText('Operand B')).toHaveValue('')
  expect(screen.getByText('0')).toBeInTheDocument()
})

it('AC button clears everything and resets operation to add', async () => {
  render(<Calculator />)
  const user = userEvent.setup()
  await user.click(screen.getByLabelText('Divide'))
  await user.type(screen.getByLabelText('Operand A'), '10')
  await user.type(screen.getByLabelText('Operand B'), '2')
  await user.click(screen.getByText('Calculate'))
  await waitFor(() => { expect(screen.queryByText('Calculating...')).not.toBeInTheDocument() })

  await user.click(screen.getByText('AC'))

  expect(screen.getByLabelText('Operand A')).toHaveValue('')
  expect(screen.getByLabelText('Operand B')).toHaveValue('')
  expect(screen.getByText('0')).toBeInTheDocument()
  // Operation should reset to add — Divide should not be selected
  expect(screen.getByLabelText('Add')).toBeChecked()
})
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd frontend && npx vitest run src/components/Calculator.test.tsx`
Expected: FAIL — no elements with text "C" or "AC".

- [ ] **Step 3: Add C/AC buttons and handlers to Calculator.tsx**

In `frontend/src/components/Calculator.tsx`:

1. Destructure `clear` and `reset` from `useCalculator()`:
```typescript
const { selectedOperation, selectOperation, result, error, isLoading, performCalculation, clear, reset } = useCalculator()
```

2. Destructure `reset: resetForm` from `useForm` and add `clearHistory` usage:
```typescript
const { register, handleSubmit, setValue, reset: resetForm, formState: { errors } } = useForm<CalculatorFormData>({
  defaultValues: { operandA: '', operandB: '' },
})
```

3. Add handler functions:
```typescript
const handleClear = () => {
  clear()
  resetForm({ operandA: '', operandB: '' })
}

const handleAllClear = () => {
  reset()
  resetForm({ operandA: '', operandB: '' })
  clearHistory()
}
```

4. Add the button row between `<Display>` and `<form>`, inside the root `<div>` (before the `<form>` tag):
```tsx
<div className="grid grid-cols-2 gap-2">
  <button
    type="button"
    onClick={handleClear}
    className="rounded-lg bg-gray-200 p-2 text-sm font-semibold text-gray-700 transition-colors hover:bg-gray-300"
  >
    C
  </button>
  <button
    type="button"
    onClick={handleAllClear}
    className="rounded-lg bg-red-100 p-2 text-sm font-semibold text-red-700 transition-colors hover:bg-red-200"
  >
    AC
  </button>
</div>
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `cd frontend && npx vitest run src/components/Calculator.test.tsx`
Expected: All tests PASS.

- [ ] **Step 5: Commit**

```bash
git add frontend/src/components/Calculator.tsx frontend/src/components/Calculator.test.tsx
git commit -m "feat: add C (clear) and AC (all clear) buttons to calculator"
```

---

### Task 3: Add result carry-forward with auto-focus

**Files:**
- Modify: `frontend/src/components/Calculator.tsx`
- Modify: `frontend/src/components/Calculator.test.tsx`

- [ ] **Step 1: Write failing tests for carry-forward**

Add these tests to `frontend/src/components/Calculator.test.tsx`. Update the msw handler to also handle `multiply`:

At the top of the file, update the server handler to support more operations:

```typescript
const server = setupServer(
  http.post('/api/calculate', async ({ request }) => {
    const body = (await request.json()) as Record<string, unknown>
    const a = body.operand_a as number
    const b = body.operand_b as number | undefined

    if (body.operation === 'add') {
      return HttpResponse.json({
        operation: 'add', operand_a: a, operand_b: b,
        result: a + (b ?? 0),
        timestamp: '2026-11-07T15:04:05Z',
      })
    }
    if (body.operation === 'multiply') {
      return HttpResponse.json({
        operation: 'multiply', operand_a: a, operand_b: b,
        result: a * (b ?? 1),
        timestamp: '2026-11-07T15:04:05Z',
      })
    }
    if (body.operation === 'sqrt') {
      if (a < 0) {
        return HttpResponse.json(
          { error: { code: 'NEGATIVE_SQRT', message: 'Cannot take square root of negative number' } },
          { status: 422 }
        )
      }
      return HttpResponse.json({
        operation: 'sqrt', operand_a: a,
        result: Math.sqrt(a),
        timestamp: '2026-11-07T15:04:05Z',
      })
    }
    if (body.operation === 'divide' && b === 0) {
      return HttpResponse.json(
        { error: { code: 'DIVISION_BY_ZERO', message: 'Cannot divide by zero' } },
        { status: 422 }
      )
    }
    return HttpResponse.json({ operation: body.operation, operand_a: a, result: 0, timestamp: '' })
  })
)
```

Then add the carry-forward test:

```typescript
it('carries result forward as operand A after successful calculation', async () => {
  render(<Calculator />)
  const user = userEvent.setup()
  await user.type(screen.getByLabelText('Operand A'), '5')
  await user.type(screen.getByLabelText('Operand B'), '3')
  await user.click(screen.getByText('Calculate'))
  await waitFor(() => { expect(screen.getByText('8')).toBeInTheDocument() })

  // After calculation, operand A should have the result, operand B should be empty
  expect(screen.getByLabelText('Operand A')).toHaveValue('8')
  expect(screen.getByLabelText('Operand B')).toHaveValue('')
})

it('does not carry forward on error', async () => {
  render(<Calculator />)
  const user = userEvent.setup()
  await user.click(screen.getByLabelText('Divide'))
  await user.type(screen.getByLabelText('Operand A'), '10')
  await user.type(screen.getByLabelText('Operand B'), '0')
  await user.click(screen.getByText('Calculate'))
  await waitFor(() => { expect(screen.getByText('Cannot divide by zero')).toBeInTheDocument() })

  // Inputs should stay as they were
  expect(screen.getByLabelText('Operand A')).toHaveValue('10')
  expect(screen.getByLabelText('Operand B')).toHaveValue('0')
})
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd frontend && npx vitest run src/components/Calculator.test.tsx`
Expected: FAIL — operand A still has '5' after calculation, not '8'.

- [ ] **Step 3: Implement carry-forward logic in onSubmit**

In `frontend/src/components/Calculator.tsx`:

1. Add `useRef` import:
```typescript
import { useEffect, useRef } from 'react'
```

2. Create a ref for operand B:
```typescript
const operandBRef = useRef<HTMLInputElement | null>(null)
```

3. Update `onSubmit` to carry forward on success:
```typescript
const onSubmit = async (data: CalculatorFormData) => {
  const parsedA = operandSchema.safeParse(data.operandA)
  if (!parsedA.success) return

  let operandB: number | undefined
  if (!unary && data.operandB) {
    const parsedB = operandSchema.safeParse(data.operandB)
    if (!parsedB.success) return
    operandB = parsedB.data
  }

  const response = await performCalculation(parsedA.data, operandB)
  if (response) {
    addEntry(response)
    // Carry-forward: load result into operand A, clear operand B
    setValue('operandA', String(response.result))
    setValue('operandB', '')
    // Auto-focus operand B for binary operations
    if (!unary) {
      operandBRef.current?.focus()
    }
  }
}
```

4. Pass the ref to `OperandInputs`. First update the component's props interface in `OperandInputs.tsx` to accept a ref for operand B. **However**, since `register` already handles the ref, we need a different approach. Use `register` with a ref callback.

Update `Calculator.tsx` — replace the `register` usage for operandB with a merged ref approach:

```typescript
const { ref: operandBFormRef, ...operandBRest } = register('operandB')

// Merge refs
const setOperandBRef = (el: HTMLInputElement | null) => {
  operandBFormRef(el)
  operandBRef.current = el
}
```

Then update `OperandInputs` to accept an optional `operandBRef` prop. In `frontend/src/components/OperandInputs.tsx`:

Add `operandBRef` to the props:
```typescript
interface OperandInputsProps {
  register: UseFormRegister<CalculatorFormData>
  errors: FieldErrors<CalculatorFormData>
  isUnary: boolean
  operandBRef?: (el: HTMLInputElement | null) => void
}
```

Update the operand B input to use the ref:
```tsx
<input
  id="operandB"
  type="text"
  inputMode="decimal"
  {...register('operandB')}
  ref={operandBRef}  // Override the register ref
  onKeyDown={handleNumericKeyDown}
  disabled={isUnary}
  className={`w-full rounded-lg border p-3 text-lg ${isUnary ? 'cursor-not-allowed bg-gray-100 text-gray-400' : ''} ${errors.operandB ? 'border-red-500' : 'border-gray-300'}`}
  placeholder={isUnary ? 'N/A' : 'Enter number'}
/>
```

Wait — spreading `register('operandB')` and then overriding `ref` will break React Hook Form. Better approach: destructure register in Calculator and pass the merged ref.

**Simpler approach** — handle ref merging in `Calculator.tsx` only:

In `Calculator.tsx`, instead of passing `register` directly to `OperandInputs`, pass individual registration objects. Actually, this would be a large refactor. Let's keep it simpler:

Just use `setTimeout` to focus after state updates:
```typescript
if (response) {
  addEntry(response)
  setValue('operandA', String(response.result))
  setValue('operandB', '')
  if (!unary) {
    setTimeout(() => {
      const operandBInput = document.getElementById('operandB') as HTMLInputElement | null
      operandBInput?.focus()
    }, 0)
  }
}
```

This avoids any ref plumbing changes to `OperandInputs`.

- [ ] **Step 4: Run tests to verify they pass**

Run: `cd frontend && npx vitest run src/components/Calculator.test.tsx`
Expected: All tests PASS.

- [ ] **Step 5: Run all frontend tests**

Run: `cd frontend && npx vitest run`
Expected: All tests PASS (including Display, OperandInputs, History, useCalculator, useHistory, Calculator).

- [ ] **Step 6: Commit**

```bash
git add frontend/src/components/Calculator.tsx frontend/src/components/Calculator.test.tsx
git commit -m "feat: carry forward result as operand A after successful calculation"
```

---

### Task 4: Final verification

- [ ] **Step 1: Run all frontend tests**

Run: `cd frontend && npx vitest run`
Expected: All tests PASS.

- [ ] **Step 2: Manual smoke test**

Rebuild Docker and verify in browser:
```bash
docker build -t sezzle-calculator .
docker run -p 3000:80 sezzle-calculator
```

Verify:
1. Calculate 5 + 3 = 8 → operand A becomes "8", operand B is cleared
2. Press C → inputs clear, result shows 0, operation stays the same
3. Press AC → inputs clear, result shows 0, history clears, operation resets to Add
4. Division by zero → inputs stay as they were (no carry-forward)
5. Browser tab says "Calculator"

- [ ] **Step 3: Commit any final adjustments**

```bash
git add -A
git commit -m "feat: calculator UX improvements - carry-forward, C/AC buttons"
```
