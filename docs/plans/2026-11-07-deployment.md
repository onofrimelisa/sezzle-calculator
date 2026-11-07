# Deployment & Root Documentation Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Create unified Docker deployment (multi-stage Dockerfile + nginx), docker-compose for development, and root-level project documentation.

**Architecture:** Multi-stage Docker build: Go binary + React static assets served by nginx with reverse proxy. docker-compose for local dev with hot-reload.

**Tech Stack:** Docker, nginx, docker-compose

**Spec:** `docs/specs/2026-11-07-fullstack-calculator-design.md`

---

### Task 1: Nginx Configuration

**Files:**
- Create: `nginx/nginx.conf`

- [ ] **Step 1: Create nginx config**

```nginx
# nginx/nginx.conf
server {
    listen 80;
    server_name localhost;

    # Frontend static files
    root /usr/share/nginx/html;
    index index.html;

    # SPA fallback
    location / {
        try_files $uri $uri/ /index.html;
    }

    # Proxy API requests to Go backend
    location /api/ {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Proxy health check
    location /healthz {
        proxy_pass http://localhost:8080;
    }

    # Proxy metrics
    location /metrics {
        proxy_pass http://localhost:8080;
    }

    # Proxy swagger
    location /swagger/ {
        proxy_pass http://localhost:8080;
    }
}
```

- [ ] **Step 2: Commit**

```bash
git add .
git commit -m "chore: add nginx configuration for reverse proxy"
```

---

### Task 2: Multi-stage Dockerfile

**Files:**
- Create: `Dockerfile`

- [ ] **Step 1: Create Dockerfile**

```dockerfile
# Dockerfile
# Multi-stage build for full-stack calculator application

# Stage 1: Build Go backend
FROM golang:1.22-alpine AS backend-builder

WORKDIR /app/backend

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ .

# Generate swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/server/main.go -o docs

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server

# Stage 2: Build React frontend
FROM node:20-alpine AS frontend-builder

WORKDIR /app/frontend

COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci

COPY frontend/ .
RUN npm run build

# Stage 3: Runtime
FROM nginx:alpine

# Install supervisor to run both nginx and Go binary
RUN apk add --no-cache supervisor

# Copy Go binary
COPY --from=backend-builder /server /usr/local/bin/server

# Copy frontend build
COPY --from=frontend-builder /app/frontend/dist /usr/share/nginx/html

# Copy nginx config
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf

# Copy supervisor config
COPY supervisord.conf /etc/supervisord.conf

EXPOSE 80

CMD ["supervisord", "-c", "/etc/supervisord.conf"]
```

- [ ] **Step 2: Create supervisord config**

```ini
; supervisord.conf
[supervisord]
nodaemon=true
user=root
logfile=/var/log/supervisord.log

[program:backend]
command=/usr/local/bin/server
autostart=true
autorestart=true
stdout_logfile=/dev/stdout
stdout_logfile_maxbytes=0
stderr_logfile=/dev/stderr
stderr_logfile_maxbytes=0
environment=PORT="8080",LOG_LEVEL="info",CORS_ORIGINS="*"

[program:nginx]
command=nginx -g "daemon off;"
autostart=true
autorestart=true
stdout_logfile=/dev/stdout
stdout_logfile_maxbytes=0
stderr_logfile=/dev/stderr
stderr_logfile_maxbytes=0
```

- [ ] **Step 3: Create .dockerignore**

```
# .dockerignore
**/node_modules
**/dist
**/bin
**/.git
**/.env
**/coverage*
**/*.out
```

- [ ] **Step 4: Build and test Docker image**

```bash
docker build -t sezzle-calculator .
docker run -p 3000:80 sezzle-calculator
```

Open `http://localhost:3000` — should see calculator UI.
Open `http://localhost:3000/healthz` — should return `{"status":"healthy"}`.
Open `http://localhost:3000/swagger/index.html` — should see Swagger UI.

- [ ] **Step 5: Commit**

```bash
git add .
git commit -m "feat: add multi-stage Dockerfile with nginx and supervisor"
```

---

### Task 3: Docker Compose for Development

**Files:**
- Create: `docker-compose.yml`

- [ ] **Step 1: Create docker-compose.yml**

```yaml
# docker-compose.yml
# Development environment with hot-reload

services:
  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile.dev
    ports:
      - "8080:8080"
    volumes:
      - ./backend:/app
    environment:
      - PORT=8080
      - LOG_LEVEL=debug
      - CORS_ORIGINS=http://localhost:5173

  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile.dev
    ports:
      - "5173:5173"
    volumes:
      - ./frontend:/app
      - /app/node_modules
    environment:
      - VITE_API_URL=http://localhost:8080
    depends_on:
      - backend
```

- [ ] **Step 2: Create backend dev Dockerfile**

```dockerfile
# backend/Dockerfile.dev
FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]
```

- [ ] **Step 3: Create air config for Go hot-reload**

```toml
# backend/.air.toml
root = "."
tmp_dir = "tmp"

[build]
  cmd = "go build -o ./tmp/server ./cmd/server"
  bin = "./tmp/server"
  include_ext = ["go"]
  exclude_dir = ["tmp", "docs", "spec", "scripts"]

[log]
  time = true
```

- [ ] **Step 4: Create frontend dev Dockerfile**

```dockerfile
# frontend/Dockerfile.dev
FROM node:20-alpine

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]
```

- [ ] **Step 5: Test docker-compose**

```bash
docker-compose up --build
```

Expected:
- Frontend at `http://localhost:5173`
- Backend at `http://localhost:8080`
- Both with hot-reload

- [ ] **Step 6: Commit**

```bash
git add .
git commit -m "feat: add docker-compose for development with hot-reload"
```

---

### Task 4: Root README

**Files:**
- Create: `README.md`

- [ ] **Step 1: Write root README**

```markdown
# Sezzle Calculator

Full-stack calculator application — React (TypeScript) frontend + Go backend REST API.

## Quick Start

### With Docker (recommended)

```bash
docker build -t sezzle-calculator .
docker run -p 3000:80 sezzle-calculator
```

Open `http://localhost:3000`.

### Without Docker

**Backend:**
```bash
cd backend
go run ./cmd/server
```
Server starts at `http://localhost:8080`.

**Frontend:**
```bash
cd frontend
npm install
npm run dev
```
App starts at `http://localhost:5173` (proxies API to `:8080`).

### Development (docker-compose)

```bash
docker-compose up --build
```
Frontend at `:5173`, Backend at `:8080`, both with hot-reload.

## API Examples

**Addition:**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "add", "operand_a": 5, "operand_b": 3}'
```

Response:
```json
{
  "operation": "add",
  "operand_a": 5,
  "operand_b": 3,
  "result": 8,
  "timestamp": "2026-11-07T15:04:05Z"
}
```

**Square Root (unary):**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "sqrt", "operand_a": 9}'
```

**Division by Zero (error):**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "divide", "operand_a": 10, "operand_b": 0}'
```

Response:
```json
{
  "error": {
    "code": "DIVISION_BY_ZERO",
    "message": "Cannot divide by zero",
    "details": {"operation": "divide", "operand_a": 10, "operand_b": 0}
  }
}
```

### Available Operations

| Operation | Type | Example |
|---|---|---|
| `add` | Binary | 5 + 3 = 8 |
| `subtract` | Binary | 10 - 4 = 6 |
| `multiply` | Binary | 6 × 7 = 42 |
| `divide` | Binary | 15 ÷ 3 = 5 |
| `power` | Binary | 2 ^ 8 = 256 |
| `sqrt` | Unary | √9 = 3 |
| `percentage` | Binary | 15% of 200 = 30 |

## Testing

```bash
# Backend
cd backend && make test && make test-coverage

# Frontend
cd frontend && npx vitest run --coverage
```

## Swagger UI

Start the backend and open `http://localhost:8080/swagger/index.html`.

## Design Decisions

See detailed documentation:
- [Backend Architecture](backend/docs/DESIGN.md)
- [Frontend Architecture](frontend/docs/DESIGN.md)
- [API Reference](backend/docs/API.md)
- [Spec-Driven Development](backend/docs/SDD.md)
- [Original Assignment](docs/ASSIGNMENT.md)
- [AI Prompts Used](docs/PROMPTS.md)

### Key Decisions

1. **Single endpoint (`POST /api/calculate`)** over multiple endpoints — cleaner API, easier to extend via Strategy pattern.
2. **Strategy + Registry patterns** — each operation is an independent implementation, registered at startup. Adding a new operation requires zero changes to existing code.
3. **Layered architecture** (domain → service → handler) — separation of concerns without over-engineering (no hexagonal since there are no secondary ports).
4. **Spec-Driven Development** — each operation has a spec that defines behavior before implementation. `make propose` → `make apply` workflow.
5. **slog (stdlib)** over third-party loggers — demonstrates knowledge of modern Go stdlib.
6. **Vite + Tailwind** over CRA + CSS modules — modern, fast, mobile-first responsive out of the box.
7. **Custom hooks** (useCalculator, useHistory) — all business logic separated from UI components.
```

- [ ] **Step 2: Commit**

```bash
git add .
git commit -m "docs: add root README with setup, examples, and design decisions"
```

---

### Task 5: PROMPTS.md

**Files:**
- Create: `docs/PROMPTS.md`

- [ ] **Step 1: Create PROMPTS.md structure**

```markdown
# AI Prompts Used

This document captures the key AI prompts used during development, focusing on the Spec-Driven Development (SDD) workflow and architectural decisions.

## SDD Workflow Prompts

### Proposing a New Operation

When adding a new operation, the SDD flow starts with `make propose op=<name>`, which generates a spec template. The spec is then refined through AI-assisted analysis:

```
> Analyze the spec for the divide operation. Check for completeness:
> inputs, outputs, edge cases, error codes, test cases.
> Flag anything missing or ambiguous.
```

### Implementing from Spec

After the spec passes analysis, implementation is generated from the spec:

```
> Implement the divide operation from spec/operations/divide.md.
> Follow the Strategy pattern established by add.go.
> Include: operation struct, Execute method, metadata methods,
> table-driven tests covering all spec test cases.
```

### Verifying Implementation

```
> Run make verify op=divide. Check that all test cases from the spec
> are covered. Verify edge cases: division by zero returns
> DIVISION_BY_ZERO error code.
```

## Architecture Decision Prompts

*(To be populated during implementation with actual prompts used)*

## Key Observations

- SDD prompts reference the spec document explicitly, ensuring AI output aligns with documented behavior
- Implementation prompts reference existing code patterns for consistency
- Verification prompts cross-reference spec test cases against actual test coverage
```

- [ ] **Step 2: Commit**

```bash
git add .
git commit -m "docs: add PROMPTS.md for AI prompt documentation"
```

---

### Task 6: Git Repository Setup

**Files:**
- Create: `.gitignore`

- [ ] **Step 1: Create root .gitignore**

```
# .gitignore

# Go
backend/bin/
backend/tmp/
backend/coverage.*
*.out

# Node
frontend/node_modules/
frontend/dist/

# IDE
.idea/
.vscode/
*.swp

# OS
.DS_Store

# Docker
docker-compose.override.yml

# Environment
.env
.env.local
```

- [ ] **Step 2: Initialize git and make initial commit**

```bash
git init
git add .
git commit -m "initial commit: project structure and documentation"
```

---

### Final Verification

- [ ] **Docker build succeeds**

```bash
docker build -t sezzle-calculator .
```

- [ ] **Docker container runs and serves app**

```bash
docker run -p 3000:80 sezzle-calculator
curl http://localhost:3000/healthz
curl -X POST http://localhost:3000/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "add", "operand_a": 2, "operand_b": 3}'
```

Expected: Health check returns healthy, calculate returns result 5.
