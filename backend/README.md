# Calculator Backend

Go REST API microservice for arithmetic operations with Strategy pattern, structured logging, and Prometheus metrics.

## Quick Start

### Prerequisites
- Go 1.22+
- swag CLI: `go install github.com/swaggo/swag/cmd/swag@latest`

### Run
```bash
make run
```
Server starts at `http://localhost:8080`.

### Test
```bash
make test
make test-coverage
```

### API Examples

**Addition:**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "add", "operand_a": 5, "operand_b": 3}'
```

**Square Root (unary):**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "sqrt", "operand_a": 9}'
```

**Health Check:**
```bash
curl http://localhost:8080/healthz
```

### Swagger UI
Open `http://localhost:8080/swagger/index.html` after starting the server.

### Environment Variables

| Variable | Default | Description |
|---|---|---|
| `PORT` | `8080` | Server port |
| `LOG_LEVEL` | `info` | Log level (debug, info, warn, error) |
| `CORS_ORIGINS` | `*` | Allowed CORS origins |
| `READ_TIMEOUT` | `5s` | HTTP read timeout |
| `WRITE_TIMEOUT` | `10s` | HTTP write timeout |

## Documentation
- [Design & Architecture](docs/DESIGN.md)
- [API Reference](docs/API.md)
- [Spec-Driven Development](docs/SDD.md)
