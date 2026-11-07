import { describe, it, expect, beforeAll, afterAll, afterEach } from 'vitest'
import { http, HttpResponse } from 'msw'
import { setupServer } from 'msw/node'
import { calculate } from './calculator'
import { CalculatorApiError } from './types'

const server = setupServer()

beforeAll(() => server.listen())
afterEach(() => server.resetHandlers())
afterAll(() => server.close())

describe('calculate', () => {
  it('returns result for successful calculation', async () => {
    server.use(
      http.post('/api/calculate', () => {
        return HttpResponse.json({
          operation: 'add', operand_a: 5, operand_b: 3, result: 8,
          timestamp: '2026-11-07T15:04:05Z',
        })
      })
    )
    const result = await calculate({ operation: 'add', operand_a: 5, operand_b: 3 })
    expect(result.result).toBe(8)
    expect(result.operation).toBe('add')
  })

  it('throws CalculatorApiError on domain error', async () => {
    server.use(
      http.post('/api/calculate', () => {
        return HttpResponse.json(
          { error: { code: 'DIVISION_BY_ZERO', message: 'Cannot divide by zero', details: { operation: 'divide' } } },
          { status: 422 }
        )
      })
    )
    await expect(calculate({ operation: 'divide', operand_a: 10, operand_b: 0 })).rejects.toThrow(CalculatorApiError)
  })
})
