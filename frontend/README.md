# Calculator Frontend

React + TypeScript calculator UI built with Vite.

## Quick Start

### Prerequisites
- Node.js 18+

### Install
```bash
npm install
```

### Run (development)
```bash
npm run dev
```
App starts at `http://localhost:5173`. Proxies `/api/*` to `http://localhost:8080`.

### Test
```bash
npx vitest run
npx vitest run --coverage
```

### Build
```bash
npm run build
```
Static files output to `dist/`.

## Architecture

See [docs/DESIGN.md](docs/DESIGN.md) for detailed architecture documentation.

### Component Structure
- **App** — Error Boundary + layout
- **Calculator** (Container) — Orchestrates hooks and child components
- **Display** (Presentational) — Shows result or error
- **OperandInputs** (Presentational) — Input fields with validation
- **OperationButtons** (Presentational) — Operation selection buttons
- **History** (Presentational) — Past calculation list

### Key Patterns
- Container/Presentational components
- Custom Hooks (useCalculator, useHistory)
- Adapter Pattern (API layer)
- Error Boundary
