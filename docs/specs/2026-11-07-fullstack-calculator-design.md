# Full-Stack Calculator Application — Design Spec

## Overview

Full-stack calculator application for Sezzle Senior Software Engineer technical assessment. React (TypeScript) frontend consuming a Go backend REST API for arithmetic operations. Focus on clean architecture, testability, observability, and spec-driven development.

## Architecture

```
┌─────────────────────────────────────────────────┐
│                   Docker                         │
│  ┌──────────┐    nginx    ┌──────────────────┐  │
│  │ Frontend │◄──────────►│     Backend       │  │
│  │  React   │  /api/*    │       Go          │  │
│  │  (static)│  proxy     │                   │  │
│  └──────────┘            └──────────────────┘  │
└─────────────────────────────────────────────────┘
```

- **Frontend:** React + TypeScript served as static assets via nginx
- **Backend:** Go REST API behind nginx reverse proxy at `/api/*`
- **Deployment:** Multi-stage Dockerfile + docker-compose for development

---

## Backend

### Tech Stack

| Tool | Purpose |
|---|---|
| Go 1.22+ | Language |
| Chi | HTTP router (idiomatic, net/http compatible) |
| swaggo/swag | Swagger/OpenAPI generation from annotations |
| validator/v10 | Struct validation with tags |
| testify | Test assertions and mocks |
| slog (stdlib) | Structured logging |
| prometheus/client_golang | Metrics |

### Project Structure

```
backend/
├── cmd/server/              # Entrypoint, wiring, graceful shutdown
├── internal/
│   ├── domain/              # Operation interfaces (Strategy), Registry, domain errors
│   │   ├── operation.go     # BinaryOperation + UnaryOperation interfaces
│   │   ├── registry.go      # Registry pattern: map[string]Operation with metadata
│   │   └── errors.go        # Domain errors: DivisionByZero, InvalidOperation, etc.
│   ├── operations/          # Strategy implementations
│   │   ├── add.go
│   │   ├── subtract.go
│   │   ├── multiply.go
│   │   ├── divide.go
│   │   ├── power.go
│   │   ├── sqrt.go
│   │   └── percentage.go
│   ├── service/             # CalculatorService: orchestrates domain, emits metrics/logs
│   │   └── calculator.go
│   ├── handler/             # HTTP handlers + Swagger annotations
│   │   └── calculator.go
│   ├── dto/                 # Request/Response structs with validation
│   │   └── calculator.go
│   ├── middleware/          # RequestID, logging, CORS, recovery, metrics
│   │   ├── requestid.go
│   │   ├── logging.go
│   │   ├── cors.go
│   │   ├── recovery.go
│   │   └── metrics.go
│   └── platform/
│       ├── logger/          # slog configuration
│       ├── metrics/         # Prometheus metrics definitions
│       └── config/          # Env-based configuration
├── docs/                    # Swagger generated files
└── spec/                    # SDD operation specs
    └── operations/
```

### Design Patterns

- **Strategy Pattern:** `BinaryOperation` and `UnaryOperation` interfaces. Each operation is an independent implementation.
- **Registry Pattern:** Map of `string -> Operation` with metadata (category, type, description). Dynamic dispatch based on `operation` field in request.
- **DTO Pattern:** Separate request/response structs. Domain models are not exposed to HTTP layer.
- **Middleware Chain:** Cross-cutting concerns as composable decorators.

### API

#### `POST /api/calculate`

**Request (binary operation):**
```json
{
  "operation": "add",
  "operand_a": 5,
  "operand_b": 3
}
```

**Request (unary operation):**
```json
{
  "operation": "sqrt",
  "operand_a": 9
}
```

**Success response (200):**
```json
{
  "operation": "add",
  "operand_a": 5,
  "operand_b": 3,
  "result": 8,
  "timestamp": "2026-11-07T15:04:05Z"
}
```

**Error response (400/422):**
```json
{
  "error": {
    "code": "DIVISION_BY_ZERO",
    "message": "Cannot divide by zero",
    "details": {
      "operation": "divide",
      "operand_a": 10,
      "operand_b": 0
    }
  }
}
```

#### Error Codes

| Code | HTTP Status | When |
|---|---|---|
| `INVALID_OPERATION` | 400 | Operation does not exist in registry |
| `INVALID_INPUT` | 422 | Missing required field, invalid type |
| `DIVISION_BY_ZERO` | 422 | Division by zero |
| `NEGATIVE_SQRT` | 422 | Square root of negative number |
| `MISSING_OPERAND` | 422 | Binary operation without `operand_b` |
| `UNEXPECTED_OPERAND` | 422 | Unary operation with `operand_b` provided |

#### Operations

| Operation | Type | Category | Description |
|---|---|---|---|
| `add` | Binary | Basic Operations | Addition: operand_a + operand_b |
| `subtract` | Binary | Basic Operations | Subtraction: operand_a - operand_b |
| `multiply` | Binary | Basic Operations | Multiplication: operand_a * operand_b |
| `divide` | Binary | Basic Operations | Division: operand_a / operand_b |
| `power` | Binary | Advanced Operations | Exponentiation: operand_a ^ operand_b |
| `sqrt` | Unary | Advanced Operations | Square root of operand_a |
| `percentage` | Binary | Advanced Operations | Percentage: operand_b% of operand_a |

#### Additional Endpoints

- `GET /healthz` — Liveness/readiness check
- `GET /metrics` — Prometheus metrics
- `GET /swagger/*` — Swagger UI

### Observability

**Structured Logging (slog):**
- Every request logged with: `request_id`, `operation`, `duration_ms`, `status_code`
- Domain errors logged as `WARN`, unexpected errors as `ERROR`
- Log level configurable via `LOG_LEVEL` env var

**Prometheus Metrics:**

| Metric | Type | Labels |
|---|---|---|
| `calculator_requests_total` | Counter | `operation`, `status` |
| `calculator_request_duration_seconds` | Histogram | `operation` |
| `calculator_errors_total` | Counter | `operation`, `error_code` |

**Middleware Chain (execution order):**
1. Recovery — panic → 500 with log
2. Request ID — generates UUID, injects in context and `X-Request-ID` response header
3. Logging — logs every request with duration
4. Metrics — increments counters, records latency
5. CORS — configured via env var

**Configuration (env vars):**

| Variable | Default | Description |
|---|---|---|
| `PORT` | `8080` | Server port |
| `LOG_LEVEL` | `info` | Log verbosity (debug, info, warn, error) |
| `CORS_ORIGINS` | `*` | Allowed origins |
| `READ_TIMEOUT` | `5s` | HTTP read timeout |
| `WRITE_TIMEOUT` | `10s` | HTTP write timeout |

**Graceful Shutdown:**
- Listens for `SIGINT`/`SIGTERM`
- Drains active connections with configurable timeout
- Clean shutdown log

### SDD (Spec-Driven Development)

Commands (Makefile):

| Command | Description |
|---|---|
| `make propose op=<name>` | Generates spec template at `spec/operations/<name>.md` |
| `make analyze op=<name>` | Validates spec completeness: inputs, outputs, edge cases, tests |
| `make apply op=<name>` | Scaffolds implementation from spec: operation, tests, swagger |
| `make sync` | Regenerates swagger docs, validates spec-code alignment |
| `make verify op=<name>` | Runs tests + lint + coverage for that operation |

**Spec template:**
```markdown
# Operation: <name>

## Description
Brief description of what this operation does.

## Category
Basic | Advanced

## Type
Binary | Unary

## Inputs
- operand_a: float64 (required) — description
- operand_b: float64 (required|optional) — description

## Output
- result: float64 — description

## Edge Cases
- [ ] Case 1 → expected error code
- [ ] Case 2 → expected error code

## Examples
| operand_a | operand_b | result |
|-----------|-----------|--------|
| 5         | 3         | ?      |

## Test Cases
- [ ] Happy path with positive numbers
- [ ] Happy path with negative numbers
- [ ] Edge case 1
- [ ] Edge case 2

## Swagger
- Tags: Basic Operations | Advanced Operations
- Summary: ...
- Error responses: ...
```

**Workflow:**
```
make propose op=sqrt → edit spec → make analyze op=sqrt → make apply op=sqrt → make verify op=sqrt → make sync
```

---

## Frontend

### Tech Stack

| Tool | Purpose |
|---|---|
| React 18+ | UI framework |
| TypeScript | Type safety |
| Vite | Build tool |
| Tailwind CSS | Utility-first CSS, responsive design |
| React Hook Form | Form state and validation |
| Zod | Schema validation |
| Axios | HTTP client with interceptors |
| Vitest | Test runner |
| React Testing Library | Component testing |
| msw (Mock Service Worker) | API mocking for integration tests |

### Project Structure

```
frontend/
├── src/
│   ├── components/
│   │   ├── App.tsx                # Error Boundary, layout
│   │   ├── Calculator.tsx         # Container: orchestrates state via hooks
│   │   ├── Display.tsx            # Presentational: shows result or error
│   │   ├── OperandInputs.tsx      # Presentational: two inputs, B disabled for unary
│   │   ├── OperationButtons.tsx   # Presentational: buttons grouped by category
│   │   └── History.tsx            # Presentational: list of past calculations
│   ├── hooks/
│   │   ├── useCalculator.ts       # Calculation state, API calls, loading/error/result
│   │   └── useHistory.ts          # In-memory history array
│   ├── api/
│   │   ├── client.ts              # Axios instance, interceptors, base URL
│   │   ├── calculator.ts          # calculate(operation, operandA, operandB?) → Result
│   │   └── types.ts               # CalculateRequest, CalculateResponse, ApiError
│   ├── schemas/
│   │   └── calculator.ts          # Zod schemas for input validation
│   └── utils/
│       └── operations.ts          # Operation metadata: labels, categories, unary/binary
├── docs/
│   └── DESIGN.md                  # Architecture, patterns, decisions
└── README.md                      # Setup, run, test
```

### Design Patterns

- **Container/Presentational:** `Calculator` manages logic, child components only render.
- **Custom Hooks:** `useCalculator` and `useHistory` encapsulate all business logic away from components.
- **Adapter Pattern:** `api/` layer abstracts HTTP details. Components never touch Axios directly.
- **Error Boundary:** Catches unexpected React errors with a fallback UI.

### UI Behavior

- **Layout:** Display at top, two operand inputs, operation buttons grouped (Basic / Advanced), history below.
- **Unary operations:** When a unary operation (sqrt) is selected, Operand B input is disabled (grayed out) and its value is cleared. Re-enabled when switching back to a binary operation.
- **Validation:** Inline error messages under each input. Operands must be valid numbers.
- **Responsive (mobile-first with Tailwind):**
  - Mobile: inputs stack vertically, buttons full width, history below
  - Desktop: compact layout as described in the wireframe

---

## Testing Strategy

### Backend

| Level | What | Tools |
|---|---|---|
| Unit | Each operation in isolation (Strategy) | testify, table-driven tests |
| Unit | CalculatorService with mocked operations | testify/mock |
| Unit | DTO validation | validator/v10 |
| Integration | Handlers via httptest.Server | httptest, testify |
| Integration | Full pipeline: request → middleware → handler → service → operation → response | httptest |

### Frontend

| Level | What | Tools |
|---|---|---|
| Unit | `useCalculator` hook with mocked API | Vitest, renderHook |
| Unit | `useHistory` hook | Vitest, renderHook |
| Unit | Presentational components (Display, OperationButtons) | React Testing Library |
| Integration | Full Calculator: input → click operation → display result | React Testing Library |
| Integration | Error flow: API error → display error message | msw |

### Coverage Targets

- Backend: 90%+
- Frontend: 80%+

### Commands

- `make test` — runs all tests
- `make test-coverage` — generates coverage report

---

## Deployment

### Multi-stage Dockerfile

- **Stage 1:** Build Go binary (static, from scratch-friendly)
- **Stage 2:** Build React assets (`npm run build`)
- **Stage 3:** Runtime with nginx serving frontend static files and proxying `/api/*` to Go backend

### docker-compose.yml (development)

- Frontend container with Vite hot-reload
- Backend container with Go hot-reload (air)
- Shared network

---

## Documentation

### Backend
- `backend/README.md` — Setup, run, test, API examples with curl
- `backend/docs/DESIGN.md` — Architecture, patterns, trade-offs, diagram
- `backend/docs/API.md` — Endpoint reference (complements Swagger)
- `backend/docs/SDD.md` — Spec-driven development flow

### Frontend
- `frontend/README.md` — Setup, run, test, component structure
- `frontend/docs/DESIGN.md` — Architecture, patterns, decisions

### Root
- `README.md` — Project overview, quick start with Docker, links to each service's docs
- `docs/ASSIGNMENT.md` — Original assignment
- `docs/PROMPTS.md` — Relevant AI prompts used (focus on SDD prompts)
- `Dockerfile` — Multi-stage build
- `docker-compose.yml` — Dev environment
