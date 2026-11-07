# Design & Architecture

## Architecture Overview

The calculator backend follows a **layered architecture** with strict dependency direction: outer layers depend on inner layers, never the reverse. Cross-cutting concerns (logging, metrics, request tracing) live in a platform layer accessible to all other layers.

```
                    +--------------------------+
                    |       cmd/server         |  Composition root
                    +--------------------------+
                              |
                    +--------------------------+
                    |   middleware (chain)      |  Recovery, RequestID, Logging, Metrics, CORS
                    +--------------------------+
                              |
                    +--------------------------+
                    |       handler            |  HTTP layer: decode, validate, encode
                    +--------------------------+
                              |
                    +--------------------------+
                    |       service            |  Orchestration: registry lookup, dispatch
                    +--------------------------+
                              |
              +---------------+---------------+
              |                               |
    +-----------------+             +-----------------+
    |    operations    |             |     domain      |
    |  (strategies)   |             |  (interfaces,   |
    |  add, subtract, |             |   registry,     |
    |  multiply, ...  |             |   errors)       |
    +-----------------+             +-----------------+
              |                               |
              +---------------+---------------+
                              |
                    +--------------------------+
                    |   platform (cross-cut)   |
                    |   config, logger, metrics|
                    +--------------------------+
```

### Layer Responsibilities

| Layer | Package | Responsibility |
|---|---|---|
| **Domain** | `internal/domain` | Interfaces (`BinaryOperation`, `UnaryOperation`), `Registry`, domain errors |
| **Operations** | `internal/operations` | Concrete strategy implementations (add, subtract, multiply, divide, power, sqrt, percentage) |
| **Service** | `internal/service` | `CalculatorService` orchestrates registry lookup, operand validation, and dispatch |
| **DTO** | `internal/dto` | Request/response structs for HTTP serialization, decoupled from domain types |
| **Handler** | `internal/handler` | HTTP handlers: JSON decode, validation, error mapping to HTTP status codes |
| **Middleware** | `internal/middleware` | Composable HTTP middleware: recovery, request ID, logging, metrics |
| **Platform** | `internal/platform/*` | Cross-cutting: config loading, slog logger setup, Prometheus metrics registration |
| **Composition Root** | `cmd/server` | Wires everything together, starts HTTP server with graceful shutdown |

## Design Patterns

### Strategy Pattern

The core extensibility mechanism. Each arithmetic operation is an independent struct implementing one of two interfaces:

```go
type BinaryOperation interface {
    OperationMeta
    Execute(a, b float64) (float64, error)
}

type UnaryOperation interface {
    OperationMeta
    Execute(a float64) (float64, error)
}
```

Both embed `OperationMeta` which provides `Name()`, `Description()`, and `Category()`. This makes each operation self-describing and independently testable.

**Why two interfaces instead of one:** A single `Execute(a float64, b *float64)` interface would push nil-checking into every operation. Splitting into Binary and Unary makes the contract explicit at the type level -- a binary operation always receives two operands, a unary always receives one. The service layer handles the dispatch.

### Registry Pattern

The `Registry` maps operation names (strings) to `RegistryEntry` structs, enabling dynamic dispatch without a switch statement:

```go
type RegistryEntry struct {
    Type        OperationType   // "binary" or "unary"
    Name        string
    Description string
    Category    string
    Binary      BinaryOperation // populated if Type == Binary
    Unary       UnaryOperation  // populated if Type == Unary
}
```

Adding a new operation requires:
1. Implement the interface
2. Call `registry.RegisterBinary(op)` or `registry.RegisterUnary(op)` in `main.go`

No other code changes needed. The registry is populated in `buildRegistry()` in the composition root.

### DTO Pattern

The HTTP layer never exposes domain types. `dto.CalculateRequest` and `dto.CalculateResponse` are dedicated serialization structs:

- `OperandB` is `*float64` (pointer) so unary operations can omit it in JSON
- `Timestamp` is added by the response constructor, not by the domain
- `ErrorResponse` wraps error codes in a structured JSON envelope

This separation means the domain can evolve without breaking the API contract, and vice versa.

### Middleware Chain

Five middleware functions compose via Chi's `r.Use()`:

1. **Recovery** -- catches panics, logs stack trace, returns 500
2. **RequestID** -- generates UUID, stores in context, sets `X-Request-ID` header
3. **Logging** -- logs method, path, status, duration, request ID after each request
4. **Metrics** -- records `calculator_requests_total` and `calculator_request_duration_seconds`
5. **CORS** -- configurable origins via `CORS_ORIGINS` env var

Order matters: Recovery wraps everything (outermost), RequestID generates the ID before Logging needs it.

## Observability

### Structured Logging

Uses Go's `log/slog` (stdlib, introduced in Go 1.21) with JSON output. Every log entry includes:
- `method`, `path`, `status`, `duration_ms` (from Logging middleware)
- `request_id` (from RequestID middleware)
- `operation`, `operand_a`, `result` (from service layer, at debug level)

Log level is configurable via `LOG_LEVEL` environment variable.

### Prometheus Metrics

Three metrics registered via `promauto`:

| Metric | Type | Labels | Purpose |
|---|---|---|---|
| `calculator_requests_total` | Counter | `operation`, `status` | Total requests by operation and HTTP status |
| `calculator_request_duration_seconds` | Histogram | `operation` | Latency distribution per operation |
| `calculator_errors_total` | Counter | `operation`, `error_code` | Error count by operation and domain error code |

Exposed at `GET /metrics` in Prometheus text format.

### Request ID Tracing

Every request gets a UUID stored in context and returned via `X-Request-ID` response header. The ID appears in all log entries for that request, enabling correlation across log lines.

## Technology Choices

### Chi over Gin

Chi is a lightweight, idiomatic Go router built on `net/http`. It uses standard `http.Handler` and `http.HandlerFunc` signatures, making middleware composable without framework lock-in. Gin uses its own `gin.Context` which couples handler code to the framework.

### slog over zerolog/zap

`log/slog` is part of the Go standard library since 1.21. It provides structured logging with zero external dependencies. While zerolog and zap offer marginally better performance, slog is sufficient for this workload and eliminates a dependency.

### Layered over Hexagonal

A full hexagonal (ports and adapters) architecture would introduce interfaces for secondary ports (database, external APIs). This service has no secondary ports -- it is a pure computation engine. The layered approach is simpler and sufficient. If persistence were added later, introducing a repository interface at the domain layer would be straightforward.

### `go-playground/validator` for Input Validation

Struct tag validation (`validate:"required"`) keeps validation rules co-located with the DTO definition. The handler only needs to call `validator.Struct(req)` rather than writing manual checks.

## Trade-offs

| Decision | Benefit | Cost |
|---|---|---|
| Two operation interfaces (Binary/Unary) | Type-safe dispatch, no nil checks in operations | Service layer must switch on type |
| Registry with string keys | Easy to add operations, human-readable API | No compile-time guarantee that an operation name is valid |
| Pointer for `OperandB` | Clean JSON omission for unary ops | Nil checks required in service layer |
| No database | Zero infrastructure, stateless, horizontally scalable | No history, no persistence |
| Single `/api/calculate` endpoint | Uniform API, easy to extend | Less RESTful than `/api/add`, `/api/divide` etc. |
| Composition root wiring | Explicit dependencies, no DI framework | Verbose startup code |
