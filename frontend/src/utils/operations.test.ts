import { describe, it, expect } from 'vitest'
import { OPERATIONS, BASIC_OPERATIONS, ADVANCED_OPERATIONS, getOperation, isUnary, formatCalculation } from './operations'

describe('operations', () => {
  it('has 7 operations total', () => { expect(OPERATIONS).toHaveLength(7) })
  it('has 4 basic and 3 advanced', () => {
    expect(BASIC_OPERATIONS).toHaveLength(4)
    expect(ADVANCED_OPERATIONS).toHaveLength(3)
  })
  it('getOperation returns correct operation', () => {
    const add = getOperation('add')
    expect(add?.symbol).toBe('+')
    expect(add?.type).toBe('binary')
  })
  it('getOperation returns undefined for unknown', () => { expect(getOperation('unknown')).toBeUndefined() })
  it('isUnary returns true for sqrt', () => { expect(isUnary('sqrt')).toBe(true) })
  it('isUnary returns false for add', () => { expect(isUnary('add')).toBe(false) })
  it('formatCalculation formats binary operation', () => { expect(formatCalculation('add', 5, 3, 8)).toBe('5 + 3 = 8') })
  it('formatCalculation formats unary operation', () => { expect(formatCalculation('sqrt', 9, undefined, 3)).toBe('√9 = 3') })
})
