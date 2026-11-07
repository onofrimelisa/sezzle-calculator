# API Reference

Base URL: `http://localhost:8080`

## Endpoints

### POST /api/calculate

Perform an arithmetic calculation.

**Request**

```json
{
  "operation": "string (required)",
  "operand_a": 0.0,
  "operand_b": 0.0
}
```

| Field | Type | Required | Description |
|---|---|---|---|
| `operation` | string | Yes | Operation name (see operations table below) |
| `operand_a` | float64 | Yes | First operand |
| `operand_b` | float64 | Binary only | Second operand. Required for binary operations, must be omitted for unary operations. |

**Success Response (200 OK)**

```json
{
  "operation": "add",
  "operand_a": 5,
  "operand_b": 3,
  "result": 8,
  "timestamp": "2026-11-07T15:04:05Z"
}
```

**Error Response (400 or 422)**

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

#### Examples

**Addition (binary)**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "add", "operand_a": 5, "operand_b": 3}'
# Response: {"operation":"add","operand_a":5,"operand_b":3,"result":8,"timestamp":"..."}
```

**Subtraction (binary)**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "subtract", "operand_a": 10, "operand_b": 4}'
# Response: {"operation":"subtract","operand_a":10,"operand_b":4,"result":6,"timestamp":"..."}
```

**Multiplication (binary)**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "multiply", "operand_a": 6, "operand_b": 7}'
# Response: {"operation":"multiply","operand_a":6,"operand_b":7,"result":42,"timestamp":"..."}
```

**Division (binary)**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "divide", "operand_a": 20, "operand_b": 4}'
# Response: {"operation":"divide","operand_a":20,"operand_b":4,"result":5,"timestamp":"..."}
```

**Power (binary)**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "power", "operand_a": 2, "operand_b": 10}'
# Response: {"operation":"power","operand_a":2,"operand_b":10,"result":1024,"timestamp":"..."}
```

**Square Root (unary)**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "sqrt", "operand_a": 144}'
# Response: {"operation":"sqrt","operand_a":144,"result":12,"timestamp":"..."}
```

**Percentage (binary)**
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "percentage", "operand_a": 200, "operand_b": 15}'
# Response: {"operation":"percentage","operand_a":200,"operand_b":15,"result":30,"timestamp":"..."}
```

---

### GET /healthz

Health check endpoint.

**Response (200 OK)**
```json
{"status": "ok"}
```

---

### GET /metrics

Prometheus metrics in text exposition format.

**Response (200 OK)**
```
# HELP calculator_requests_total Total number of calculator requests
# TYPE calculator_requests_total counter
calculator_requests_total{operation="/api/calculate",status="200"} 42
...
```

---

### GET /swagger/*

Swagger UI for interactive API exploration. Open `http://localhost:8080/swagger/index.html` in a browser.

---

## Operations

| Name | Type | Category | Formula | Description |
|---|---|---|---|---|
| `add` | Binary | Basic Operations | `a + b` | Addition of two operands |
| `subtract` | Binary | Basic Operations | `a - b` | Subtraction of two operands |
| `multiply` | Binary | Basic Operations | `a * b` | Multiplication of two operands |
| `divide` | Binary | Basic Operations | `a / b` | Division of two operands |
| `power` | Binary | Advanced Operations | `a ^ b` | Exponentiation (a raised to the power of b) |
| `sqrt` | Unary | Advanced Operations | `sqrt(a)` | Square root of operand_a |
| `percentage` | Binary | Advanced Operations | `a * (b / 100)` | b percent of a |

## Error Codes

| Code | HTTP Status | Description |
|---|---|---|
| `INVALID_INPUT` | 422 Unprocessable Entity | Malformed JSON or validation failure (e.g., missing required fields) |
| `INVALID_OPERATION` | 400 Bad Request | The requested operation name does not exist in the registry |
| `MISSING_OPERAND` | 422 Unprocessable Entity | Binary operation called without `operand_b` |
| `UNEXPECTED_OPERAND` | 422 Unprocessable Entity | Unary operation called with `operand_b` present |
| `DIVISION_BY_ZERO` | 422 Unprocessable Entity | Division with `operand_b` equal to zero |
| `NEGATIVE_SQRT` | 422 Unprocessable Entity | Square root called with a negative `operand_a` |

All error responses follow the same envelope format:

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable description",
    "details": { }
  }
}
```

The `details` object includes the `operation`, `operand_a`, and `operand_b` (when applicable) that caused the error. It is omitted for `INVALID_INPUT` errors where the payload could not be parsed.
