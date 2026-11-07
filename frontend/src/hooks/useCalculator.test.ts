import { describe, it, expect, vi, beforeEach } from 'vitest'
import { renderHook, act } from '@testing-library/react'
import { useCalculator } from './useCalculator'

vi.mock('../api/calculator', () => ({
  calculate: vi.fn(),
}))

import { calculate } from '../api/calculator'
import { CalculatorApiError } from '../api/types'

const mockCalculate = vi.mocked(calculate)

describe('useCalculator', () => {
  beforeEach(() => { vi.clearAllMocks() })

  it('starts with default state', () => {
    const { result } = renderHook(() => useCalculator())
    expect(result.current.selectedOperation).toBe('add')
    expect(result.current.result).toBeNull()
    expect(result.current.error).toBeNull()
    expect(result.current.isLoading).toBe(false)
  })

  it('sets selected operation', () => {
    const { result } = renderHook(() => useCalculator())
    act(() => result.current.selectOperation('sqrt'))
    expect(result.current.selectedOperation).toBe('sqrt')
  })

  it('performs calculation successfully', async () => {
    const mockResponse = {
      operation: 'add', operand_a: 5, operand_b: 3, result: 8, timestamp: '2026-11-07T15:04:05Z',
    }
    mockCalculate.mockResolvedValueOnce(mockResponse)
    const { result } = renderHook(() => useCalculator())
    await act(async () => { await result.current.performCalculation(5, 3) })
    expect(result.current.result).toEqual(mockResponse)
    expect(result.current.error).toBeNull()
  })

  it('handles API error', async () => {
    mockCalculate.mockRejectedValueOnce(new CalculatorApiError('DIVISION_BY_ZERO', 'Cannot divide by zero'))
    const { result } = renderHook(() => useCalculator())
    act(() => result.current.selectOperation('divide'))
    await act(async () => { await result.current.performCalculation(10, 0) })
    expect(result.current.error).toBe('Cannot divide by zero')
    expect(result.current.result).toBeNull()
  })

  it('clears error on new calculation', async () => {
    mockCalculate.mockRejectedValueOnce(new CalculatorApiError('DIVISION_BY_ZERO', 'Cannot divide by zero'))
    const { result } = renderHook(() => useCalculator())
    await act(async () => { await result.current.performCalculation(10, 0) })
    expect(result.current.error).toBe('Cannot divide by zero')

    mockCalculate.mockResolvedValueOnce({
      operation: 'add', operand_a: 1, operand_b: 1, result: 2, timestamp: '2026-11-07T15:04:05Z',
    })
    await act(async () => { await result.current.performCalculation(1, 1) })
    expect(result.current.error).toBeNull()
  })

  it('clear() resets result and error but preserves operation', async () => {
    mockCalculate.mockResolvedValueOnce({
      operation: 'multiply', operand_a: 2, operand_b: 3, result: 6, timestamp: '2026-11-07T15:04:05Z',
    })
    const { result } = renderHook(() => useCalculator())
    act(() => result.current.selectOperation('multiply'))
    await act(async () => { await result.current.performCalculation(2, 3) })
    expect(result.current.result).not.toBeNull()

    act(() => result.current.clear())
    expect(result.current.result).toBeNull()
    expect(result.current.error).toBeNull()
    expect(result.current.selectedOperation).toBe('multiply')
  })

  it('reset() resets everything including operation to add', async () => {
    mockCalculate.mockResolvedValueOnce({
      operation: 'multiply', operand_a: 2, operand_b: 3, result: 6, timestamp: '2026-11-07T15:04:05Z',
    })
    const { result } = renderHook(() => useCalculator())
    act(() => result.current.selectOperation('multiply'))
    await act(async () => { await result.current.performCalculation(2, 3) })

    act(() => result.current.reset())
    expect(result.current.result).toBeNull()
    expect(result.current.error).toBeNull()
    expect(result.current.selectedOperation).toBe('add')
  })
})
