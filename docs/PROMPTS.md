# AI Prompts Used

This document captures the key AI prompts used during development, focusing on the Spec-Driven Development (SDD) workflow and architectural decisions.

## Development Approach

This project was built using AI-assisted development with a structured workflow:

1. **Design Phase** -- Collaborative brainstorming to define architecture, patterns, and API contract
2. **Planning Phase** -- Detailed implementation plans with TDD steps for each task
3. **Implementation Phase** -- Spec-driven development with atomic operations

## SDD Workflow Prompts

### Proposing a New Operation

```
make propose op=sqrt
```

This generates a spec template. The spec is then refined:

```
Analyze the spec for the sqrt operation. Validate completeness:
- Inputs: operand_a (float64, required)
- Output: result (float64, square root of operand_a)
- Edge cases: negative input returns NEGATIVE_SQRT error
- Test cases: perfect squares, non-perfect, zero, decimals, negative
```

### Implementing from Spec

```
Implement the sqrt operation from spec/operations/sqrt.md.
Follow the UnaryOperation interface from domain/operation.go.
Include:
- Sqrt struct implementing UnaryOperation
- Execute method with negative number check
- Table-driven tests covering all spec test cases
- Metadata methods (Name, Description, Category)
```

### Verifying Implementation

```
make verify op=sqrt
```

Runs lint, tests, and coverage for the specific operation.

## Architecture Decision Prompts

### API Design
```
For the calculator API, evaluate three approaches:
A) Single endpoint POST /api/calculate with operation field -- Strategy pattern dispatch
B) Multiple endpoints POST /api/add, /api/subtract -- one handler per operation
C) Expression-based POST /api/evaluate with expression parser

Recommend based on: extensibility, testability, API cleanliness.
```
Decision: Option A -- single endpoint with Strategy pattern for clean extensibility.

### Backend Architecture
```
Evaluate hexagonal architecture vs layered architecture for a calculator
microservice with no database, no external service calls, only HTTP input.
Consider: does hexagonal add value when there are no secondary ports?
```
Decision: Layered architecture -- hexagonal without output ports adds ceremony without benefit.

### Frontend Patterns
```
For a React calculator, recommend patterns that we should follow to avoid reinventing the wheel:
- State management approach
- Form handling and validation
- API integration pattern
- Component organization
```
Decision: Custom hooks + React Hook Form + Zod + Adapter pattern for API layer.

## Key Observations

- SDD prompts reference spec documents explicitly, ensuring AI output aligns with documented behavior
- Architecture prompts present trade-offs and ask for reasoned evaluation, not just implementation
- Implementation prompts reference existing patterns for consistency across operations
- Each operation follows the same TDD cycle: spec -> test -> implement -> verify
