import { describe, it, expect } from 'vitest'
import { operandSchema } from './calculator'

describe('operandSchema', () => {
  it('accepts valid number string', () => {
    const result = operandSchema.safeParse('5')
    expect(result.success).toBe(true)
    if (result.success) expect(result.data).toBe(5)
  })
  it('accepts negative numbers', () => {
    const result = operandSchema.safeParse('-3.5')
    expect(result.success).toBe(true)
    if (result.success) expect(result.data).toBe(-3.5)
  })
  it('accepts zero', () => {
    const result = operandSchema.safeParse('0')
    expect(result.success).toBe(true)
  })
  it('rejects empty string', () => {
    const result = operandSchema.safeParse('')
    expect(result.success).toBe(false)
  })
  it('rejects non-numeric string', () => {
    const result = operandSchema.safeParse('abc')
    expect(result.success).toBe(false)
  })
})
