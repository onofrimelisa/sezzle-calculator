# Sezzle Calculator

A full-stack calculator application with a Go backend API and React frontend, built using Spec-Driven Development (SDD).

## Quick Start with Docker

Build and run as a single container:

```bash
docker build -t sezzle-calculator .
docker run -p 80:80 sezzle-calculator
```

Open [http://localhost](http://localhost) in your browser.

## Quick Start without Docker

### Backend

```bash
cd backend
go run ./cmd/server
```

The API runs on [http://localhost:8080](http://localhost:8080).

### Frontend

```bash
cd frontend
npm install
npm run dev
```

The UI runs on [http://localhost:5173](http://localhost:5173).

## Development with Docker Compose

Run backend and frontend with hot reload:

```bash
docker compose up
```

- Frontend: [http://localhost:5173](http://localhost:5173)
- Backend: [http://localhost:8080](http://localhost:8080)

## API Examples

**Addition:**

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "add", "operand_a": 10, "operand_b": 5}'
# {"result": 15}
```

**Square root:**

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "sqrt", "operand_a": 144}'
# {"result": 12}
```

**Division by zero (error):**

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "divide", "operand_a": 10, "operand_b": 0}'
# {"error": {"code": "DIVISION_BY_ZERO", "message": "cannot divide by zero"}}
```

## Available Operations

| Operation  | Type   | Example                          |
|------------|--------|----------------------------------|
| `add`      | Binary | `10 + 5 = 15`                    |
| `subtract` | Binary | `10 - 5 = 5`                     |
| `multiply` | Binary | `10 * 5 = 50`                    |
| `divide`   | Binary | `10 / 5 = 2`                     |
| `negate`   | Unary  | `negate(10) = -10`               |
| `abs`      | Unary  | `abs(-10) = 10`                  |
| `sqrt`     | Unary  | `sqrt(144) = 12`                 |

## Testing

```bash
# Backend
cd backend && go test ./... -v

# Frontend
cd frontend && npm test
```

## Swagger UI

When the backend is running: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Design Decisions

1. **Single endpoint with Strategy pattern** -- All operations route through `POST /api/calculate` with a strategy dispatcher. Adding a new operation means implementing an interface and registering it -- no new routes, handlers, or middleware needed.

2. **Layered architecture over hexagonal** -- With no database or external service calls, there are no secondary/output ports. Hexagonal architecture would add adapter ceremony without benefit. A clean handler-service-operation layering keeps the code simple and testable.

3. **slog (stdlib) over third-party loggers** -- Go 1.21+ structured logging covers all needs. No dependency to manage, no API to learn beyond the standard library.

4. **Vite + Tailwind CSS** -- Fast build times, hot module replacement, and utility-first CSS for rapid, mobile-first UI development.

5. **Custom hooks for logic separation** -- `useCalculator` and `useHistory` encapsulate business logic, keeping components focused on rendering. This pattern is testable and composable without a state management library.

6. **Spec-Driven Development (SDD)** -- Each operation starts as a spec document defining inputs, outputs, edge cases, and test cases. Implementation follows the spec, ensuring consistency and traceability.

## Documentation

- [Backend Design](backend/docs/DESIGN.md) -- Architecture, patterns, and backend decisions
- [Frontend Design](frontend/docs/DESIGN.md) -- Component structure and frontend patterns
- [API Reference](backend/docs/API.md) -- Endpoint details and request/response schemas
- [SDD Guide](backend/docs/SDD.md) -- Spec-Driven Development workflow
- [Assignment](docs/ASSIGNMENT.md) -- Original project requirements
- [AI Prompts](docs/PROMPTS.md) -- Key prompts used during development
