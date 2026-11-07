# Sezzle Calculator — Project Instructions

## Overview

Full-stack calculator application for Sezzle Senior Software Engineer technical assessment.
React (TypeScript) frontend + Go backend REST API.

## Reference Documents

- **Assignment:** `docs/ASSIGNMENT.md`
- **Design Spec:** `docs/specs/2026-11-07-fullstack-calculator-design.md`
- **Backend Plan:** `docs/plans/2026-11-07-backend.md`
- **Frontend Plan:** `docs/plans/2026-11-07-frontend.md`
- **Deployment Plan:** `docs/plans/2026-11-07-deployment.md`

## Architecture

```
├── backend/          # Go REST API (Chi, slog, Prometheus)
├── frontend/         # React + TypeScript (Vite, Tailwind, Zod)
├── nginx/            # Reverse proxy config
├── docs/             # Specs, plans, prompts
├── Dockerfile        # Multi-stage production build
└── docker-compose.yml # Dev environment
```

### Backend Layers
- `domain/` — Interfaces (Strategy pattern), Registry, domain errors
- `operations/` — Strategy implementations (add, subtract, multiply, divide, power, sqrt, percentage)
- `service/` — CalculatorService orchestrates domain + observability
- `handler/` — HTTP handlers with Swagger annotations
- `dto/` — Request/Response structs with validation
- `middleware/` — RequestID, logging, CORS, recovery, metrics
- `platform/` — Config, logger, metrics definitions

### Frontend Patterns
- Container/Presentational components
- Custom Hooks (useCalculator, useHistory)
- Adapter Pattern for API layer
- Error Boundary

## Key Commands

### Backend
```bash
cd backend
make run              # Start server on :8080
make test             # Run all tests
make test-coverage    # Generate coverage report
make swagger          # Regenerate Swagger docs
```

### Backend SDD
```bash
cd backend
make propose op=<name>   # Generate spec template
make analyze op=<name>   # Validate spec completeness
make apply op=<name>     # Scaffold from spec
make verify op=<name>    # Test + lint + coverage
make sync                # Sync swagger + validate alignment
```

### Frontend
```bash
cd frontend
npm run dev           # Dev server on :5173
npx vitest run        # Run all tests
npm run build         # Production build
```

### Docker
```bash
docker build -t sezzle-calculator .    # Production image
docker run -p 3000:80 sezzle-calculator
docker-compose up --build              # Dev with hot-reload
```

## Conventions

- All documentation, error messages, and code comments in **English**
- API contract: single endpoint `POST /api/calculate` with `operation`, `operand_a`, `operand_b` fields
- Backend uses **table-driven tests** (Go idiomatic)
- Frontend uses **React Testing Library** (no enzyme, no shallow render)
- Every operation follows TDD: test first, then implement
- Domain errors use specific codes: `DIVISION_BY_ZERO`, `NEGATIVE_SQRT`, `INVALID_OPERATION`, `INVALID_INPUT`, `MISSING_OPERAND`, `UNEXPECTED_OPERAND`

## Implementation Order

1. **Backend** — domain → operations → service → handler → middleware → server
2. **Frontend** — api layer → hooks → components → app
3. **Deployment** — nginx → Dockerfile → docker-compose → root docs

## Testing

- Backend coverage target: 90%+
- Frontend coverage target: 80%+
- Backend: `testify` + `httptest` for integration
- Frontend: `vitest` + `React Testing Library` + `msw` for API mocking
