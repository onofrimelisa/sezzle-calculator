# Sezzle Calculator

Full-stack calculator application — React (TypeScript) frontend + Go backend REST API.
Built with clean architecture, Spec-Driven Development (SDD), and comprehensive testing.

## Quick Start

### Option 1: Docker (recommended)

```bash
docker build -t sezzle-calculator .
docker run -p 3000:80 sezzle-calculator
```

Open [http://localhost:3000](http://localhost:3000) in your browser.

The container runs both the backend API and frontend behind nginx on **port 3000**:
- Calculator UI: [http://localhost:3000](http://localhost:3000)
- Health check: [http://localhost:3000/healthz](http://localhost:3000/healthz)
- Swagger UI: [http://localhost:3000/swagger/index.html](http://localhost:3000/swagger/index.html)
- Prometheus metrics: [http://localhost:3000/metrics](http://localhost:3000/metrics)

### Option 2: Run locally (without Docker)

**Prerequisites:** Go 1.22+, Node.js 18+

**Backend** (terminal 1):
```bash
cd backend
make run
```
API starts at [http://localhost:8080](http://localhost:8080).

**Frontend** (terminal 2):
```bash
cd frontend
npm install
npm run dev
```
UI starts at [http://localhost:5173](http://localhost:5173) — automatically proxies `/api/*` to the backend.

### Option 3: Docker Compose (development with hot-reload)

```bash
docker compose up --build
```

- Frontend with hot-reload: [http://localhost:5173](http://localhost:5173)
- Backend with hot-reload: [http://localhost:8080](http://localhost:8080)

## API Examples

All operations go through a single endpoint: `POST /api/calculate`

**Addition (binary):**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "add", "operand_a": 5, "operand_b": 3}'
```
```json
{
  "operation": "add",
  "operand_a": 5,
  "operand_b": 3,
  "result": 8,
  "timestamp": "2026-11-07T15:04:05Z"
}
```

**Square root (unary — no operand_b):**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "sqrt", "operand_a": 9}'
```
```json
{
  "operation": "sqrt",
  "operand_a": 9,
  "result": 3,
  "timestamp": "2026-11-07T15:04:05Z"
}
```

**Division by zero (error response):**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "divide", "operand_a": 10, "operand_b": 0}'
```
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

> When running via Docker, replace `localhost:8080` with `localhost:3000` in the curl examples above.

## Available Operations

| Operation    | Type   | Category  | Example              |
|-------------|--------|-----------|----------------------|
| `add`       | Binary | Basic     | 5 + 3 = 8            |
| `subtract`  | Binary | Basic     | 10 - 4 = 6           |
| `multiply`  | Binary | Basic     | 6 x 7 = 42           |
| `divide`    | Binary | Basic     | 15 / 3 = 5           |
| `power`     | Binary | Advanced  | 2 ^ 8 = 256          |
| `sqrt`      | Unary  | Advanced  | sqrt(9) = 3          |
| `percentage`| Binary | Advanced  | 15% of 200 = 30      |

## Testing

```bash
# Backend (Go) — 59 tests
cd backend && make test

# Backend with coverage report
cd backend && make test-coverage

# Frontend (React) — 43 tests
cd frontend && npx vitest run

# Frontend with coverage
cd frontend && npx vitest run --coverage
```

## Swagger UI

Interactive API documentation is available when the backend is running:
- Local: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- Docker: [http://localhost:3000/swagger/index.html](http://localhost:3000/swagger/index.html)

## Design Decisions

1. **Single endpoint with Strategy pattern** — All operations route through `POST /api/calculate` with a Strategy + Registry dispatcher. Adding a new operation means implementing an interface and registering it — zero changes to existing code.

2. **Layered architecture over hexagonal** — With no database or external service calls, there are no secondary/output ports. Hexagonal would add ceremony without benefit. Clean handler → service → operation layering keeps the code simple and testable.

3. **slog (stdlib) over third-party loggers** — Go 1.21+ structured JSON logging. No dependency to manage, demonstrates knowledge of modern Go stdlib.

4. **Vite + Tailwind CSS** — Modern build tooling (CRA is deprecated). Tailwind provides mobile-first responsive design with utility classes — no custom CSS needed.

5. **Custom hooks for logic separation** — `useCalculator` and `useHistory` encapsulate all business logic. Components only render. Testable and composable without a state management library.

6. **Spec-Driven Development (SDD)** — Each backend operation starts as a spec document defining inputs, outputs, edge cases, and test cases. See the [SDD Guide](backend/docs/SDD.md) for the full workflow.

## Project Structure

```
.
├── backend/                  # Go REST API
│   ├── cmd/server/           # Entrypoint
│   ├── internal/
│   │   ├── domain/           # Strategy interfaces, Registry, errors
│   │   ├── operations/       # 7 operation implementations
│   │   ├── service/          # CalculatorService
│   │   ├── handler/          # HTTP handlers + Swagger
│   │   ├── dto/              # Request/Response structs
│   │   ├── middleware/       # RequestID, logging, metrics, recovery
│   │   └── platform/        # Config, logger, Prometheus metrics
│   ├── spec/operations/      # SDD specs for each operation
│   └── docs/                 # DESIGN.md, API.md, SDD.md
├── frontend/                 # React + TypeScript
│   ├── src/
│   │   ├── api/              # Axios client + adapter
│   │   ├── hooks/            # useCalculator, useHistory
│   │   ├── components/       # Calculator, Display, OperandInputs, etc.
│   │   ├── schemas/          # Zod validation
│   │   └── utils/            # Operation metadata
│   └── docs/                 # DESIGN.md
├── nginx/                    # Reverse proxy config
├── Dockerfile                # Multi-stage production build
├── docker-compose.yml        # Development with hot-reload
└── docs/                     # Assignment, prompts, specs
```

## Documentation

- [Backend Design & Architecture](backend/docs/DESIGN.md)
- [Frontend Design & Patterns](frontend/docs/DESIGN.md)
- [API Reference](backend/docs/API.md)
- [SDD Workflow Guide](backend/docs/SDD.md)
- [Original Assignment](docs/ASSIGNMENT.md)
- [AI Prompts Used](docs/PROMPTS.md)
