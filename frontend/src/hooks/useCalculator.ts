import { useState, useCallback } from 'react'
import { calculate } from '../api/calculator'
import type { CalculateResponse } from '../api/types'
import { CalculatorApiError } from '../api/types'

export function useCalculator() {
  const [selectedOperation, setSelectedOperation] = useState('add')
  const [result, setResult] = useState<CalculateResponse | null>(null)
  const [error, setError] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState(false)

  const selectOperation = useCallback((operation: string) => {
    setSelectedOperation(operation)
  }, [])

  const performCalculation = useCallback(
    async (operandA: number, operandB?: number) => {
      setIsLoading(true)
      setError(null)
      setResult(null)
      try {
        const response = await calculate({
          operation: selectedOperation,
          operand_a: operandA,
          operand_b: operandB,
        })
        setResult(response)
        return response
      } catch (err) {
        if (err instanceof CalculatorApiError) {
          setError(err.message)
        } else {
          setError('An unexpected error occurred')
        }
        return null
      } finally {
        setIsLoading(false)
      }
    },
    [selectedOperation]
  )

  return { selectedOperation, selectOperation, result, error, isLoading, performCalculation }
}
