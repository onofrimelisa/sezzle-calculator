import { describe, it, expect } from 'vitest'
import { CalculatorApiError } from './types'

describe('CalculatorApiError', () => {
  it('creates error with code and message', () => {
    const err = new CalculatorApiError('DIVISION_BY_ZERO', 'Cannot divide by zero')
    expect(err.code).toBe('DIVISION_BY_ZERO')
    expect(err.message).toBe('Cannot divide by zero')
    expect(err.name).toBe('CalculatorApiError')
  })

  it('creates error with details', () => {
    const details = { operation: 'divide', operand_a: 10, operand_b: 0 }
    const err = new CalculatorApiError('DIVISION_BY_ZERO', 'Cannot divide by zero', details)
    expect(err.details).toEqual(details)
  })
})
