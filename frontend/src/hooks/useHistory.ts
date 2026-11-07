import { useState, useCallback } from 'react'
import type { CalculateResponse } from '../api/types'

export function useHistory() {
  const [entries, setEntries] = useState<CalculateResponse[]>([])

  const addEntry = useCallback((entry: CalculateResponse) => {
    setEntries((prev) => [entry, ...prev])
  }, [])

  const clearHistory = useCallback(() => {
    setEntries([])
  }, [])

  return { entries, addEntry, clearHistory }
}
