# Spec-Driven Development (SDD)

## What is SDD?

Spec-Driven Development is a workflow where every new operation begins as a **specification document** before any code is written. The spec defines the operation's behavior, edge cases, test cases, and API metadata in a structured Markdown template. Only after the spec is reviewed and analyzed does implementation begin.

This approach ensures:
- **Deliberate design** -- edge cases and error handling are considered before writing code
- **Living documentation** -- specs stay in sync with implementations via automated checks
- **Consistent structure** -- every operation follows the same template, making the codebase predictable
- **Review-friendly** -- specs can be reviewed independently from implementation PRs

## SDD Commands

All commands are run from the `backend/` directory.

### `make propose op=<name>`

Creates a new spec file from a template at `spec/operations/<name>.md`. The template includes all required sections with placeholder values.

```bash
make propose op=modulo
# Creates spec/operations/modulo.md with template
```

### `make analyze op=<name>`

Validates that the spec file is complete. Checks for:
- All required sections are present (Description, Category, Type, Inputs, Output, Edge Cases, Examples, Test Cases, Swagger)
- No placeholder text remains (e.g., "Brief description", "one-line summary")

```bash
make analyze op=modulo
# PASSED: Spec is complete. Run: make apply op=modulo
# -- or --
# FAILED: 2 issue(s) found
```

### `make apply op=<name>`

Scaffolds the Go implementation file from the spec. Reads the `## Type` section to determine whether to generate a Binary or Unary operation struct. Creates `internal/operations/<name>.go` with a skeleton implementation (returns zero, needs `Execute()` filled in).

```bash
make apply op=modulo
# Scaffolded: internal/operations/modulo.go
```

If the implementation file already exists, the command exits without overwriting.

### `make sync`

Checks alignment between specs and implementations. For every `spec/operations/*.md` file, verifies that a corresponding `internal/operations/*.go` file exists. Also regenerates Swagger documentation.

```bash
make sync
# SYNC OK: All specs have implementations
# -- or --
# SYNC FAILED: 1 spec(s) without implementation
```

### `make verify op=<name>`

Runs lint (`go vet`) and tests for a specific operation.

```bash
make verify op=modulo
# === Verifying operation: modulo ===
# --- Lint ---
# --- Tests ---
# === Verification complete for: modulo ===
```

## Walkthrough: Adding a `modulo` Operation

This example walks through the full SDD workflow for adding a hypothetical modulo (remainder) operation.

### Step 1: Propose the Spec

```bash
make propose op=modulo
```

This creates `spec/operations/modulo.md` with a template. Edit it to fill in the details:

```markdown
# Operation: modulo

## Description
Computes the remainder of dividing operand_a by operand_b using Go's math.Mod function.

## Category
Advanced Operations

## Type
Binary

## Inputs
- operand_a: float64 (required) -- dividend
- operand_b: float64 (required) -- divisor

## Output
- result: float64 -- remainder of operand_a / operand_b

## Edge Cases
- operand_b == 0 --> DIVISION_BY_ZERO (reuse existing domain error)

## Examples
| operand_a | operand_b | result | notes |
|-----------|-----------|--------|-------|
| 10        | 3         | 1      | basic modulo |
| 10        | 5         | 0      | evenly divisible |
| -10       | 3         | -1     | negative dividend |
| 10.5      | 3         | 1.5    | decimal operands |

## Test Cases
- [x] Happy path with positive integers
- [x] Evenly divisible (result is 0)
- [x] Negative dividend
- [x] Decimal operands
- [x] Division by zero error

## Swagger
- Tags: Advanced Operations
- Summary: Modulo (remainder) of two operands
- Error responses: INVALID_INPUT, MISSING_OPERAND, DIVISION_BY_ZERO
```

### Step 2: Analyze the Spec

```bash
make analyze op=modulo
# PASSED: Spec is complete. Run: make apply op=modulo
```

If any section is missing or still has placeholder text, the command fails and tells you what to fix.

### Step 3: Scaffold the Implementation

```bash
make apply op=modulo
# Scaffolded: internal/operations/modulo.go
```

This generates a skeleton file. Open it and implement `Execute()`:

```go
func (m *Modulo) Execute(operandA, operandB float64) (float64, error) {
    if operandB == 0 {
        return 0, domain.ErrDivisionByZero
    }
    return math.Mod(operandA, operandB), nil
}
```

Update `Name()`, `Description()`, and `Category()` to match the spec. Write tests in `internal/operations/modulo_test.go` covering all test cases listed in the spec.

### Step 4: Register the Operation

Add to `buildRegistry()` in `cmd/server/main.go`:

```go
registry.RegisterBinary(operations.NewModulo())
```

### Step 5: Verify

```bash
make verify op=modulo
make sync
```

Both should pass. The operation is now live.

## Spec Template

The `make propose` command generates specs with this structure:

| Section | Purpose |
|---|---|
| `# Operation: <name>` | Operation identifier, matches filename |
| `## Description` | What the operation does, one paragraph |
| `## Category` | `Basic Operations` or `Advanced Operations` |
| `## Type` | `Binary` or `Unary` |
| `## Inputs` | List of inputs with types and descriptions |
| `## Output` | Return value with type and description |
| `## Edge Cases` | Domain-specific error conditions and their error codes |
| `## Examples` | Table of input/output examples with notes |
| `## Test Cases` | Checklist of scenarios that must be tested |
| `## Swagger` | Tags, summary, and applicable error codes for API docs |

## Integration with TDD

SDD and TDD complement each other in a four-phase cycle:

```
Spec (SDD)  -->  Test (TDD)  -->  Implement  -->  Verify
   |                                                  |
   +--------------------------------------------------+
                   (make sync ensures alignment)
```

1. **Spec**: Write the spec with `make propose` and fill in all sections. This is the design phase.
2. **Test**: Write tests from the spec's Examples and Test Cases sections. Tests should fail initially (red).
3. **Implement**: Write the operation's `Execute()` method to make tests pass (green).
4. **Verify**: Run `make verify` to confirm tests pass, then `make sync` to confirm spec-code alignment.

The spec is the source of truth for *what* the operation should do. Tests verify *that* it does it. Code is the *how*.
