import { describe, it, expect } from 'vitest'
import { renderHook, act } from '@testing-library/react'
import { useHistory } from './useHistory'
import type { CalculateResponse } from '../api/types'

const mockEntry: CalculateResponse = {
  operation: 'add', operand_a: 5, operand_b: 3, result: 8, timestamp: '2026-11-07T15:04:05Z',
}

describe('useHistory', () => {
  it('starts with empty history', () => {
    const { result } = renderHook(() => useHistory())
    expect(result.current.entries).toEqual([])
  })

  it('adds entries', () => {
    const { result } = renderHook(() => useHistory())
    act(() => result.current.addEntry(mockEntry))
    expect(result.current.entries).toHaveLength(1)
    expect(result.current.entries[0].result).toBe(8)
  })

  it('adds newest entries at the top', () => {
    const { result } = renderHook(() => useHistory())
    const entry2: CalculateResponse = { ...mockEntry, result: 15 }
    act(() => result.current.addEntry(mockEntry))
    act(() => result.current.addEntry(entry2))
    expect(result.current.entries[0].result).toBe(15)
    expect(result.current.entries[1].result).toBe(8)
  })

  it('clears history', () => {
    const { result } = renderHook(() => useHistory())
    act(() => result.current.addEntry(mockEntry))
    act(() => result.current.clearHistory())
    expect(result.current.entries).toEqual([])
  })
})
