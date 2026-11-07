export interface CalculateRequest {
  operation: string
  operand_a: number
  operand_b?: number
}

export interface CalculateResponse {
  operation: string
  operand_a: number
  operand_b?: number
  result: number
  timestamp: string
}

export interface ApiErrorDetail {
  code: string
  message: string
  details?: Record<string, unknown>
}

export interface ApiErrorResponse {
  error: ApiErrorDetail
}

export class CalculatorApiError extends Error {
  constructor(
    public readonly code: string,
    message: string,
    public readonly details?: Record<string, unknown>
  ) {
    super(message)
    this.name = 'CalculatorApiError'
  }
}
