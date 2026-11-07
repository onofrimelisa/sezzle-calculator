import { useForm } from 'react-hook-form'
import { useCalculator } from '../hooks/useCalculator'
import { useHistory } from '../hooks/useHistory'
import { isUnary } from '../utils/operations'
import { operandSchema } from '../schemas/calculator'
import type { CalculatorFormData } from '../schemas/calculator'
import { Display } from './Display'
import { OperandInputs } from './OperandInputs'
import { OperationButtons } from './OperationButtons'
import { History } from './History'
import { useEffect } from 'react'

export function Calculator() {
  const { selectedOperation, selectOperation, result, error, isLoading, performCalculation, clear, reset } = useCalculator()
  const { entries, addEntry, clearHistory } = useHistory()

  const { register, handleSubmit, setValue, reset: resetForm, formState: { errors } } = useForm<CalculatorFormData>({
    defaultValues: { operandA: '', operandB: '' },
  })

  const unary = isUnary(selectedOperation)

  useEffect(() => {
    if (unary) { setValue('operandB', '') }
  }, [unary, setValue])

  const handleClear = () => {
    clear()
    resetForm({ operandA: '', operandB: '' })
  }

  const handleAllClear = () => {
    reset()
    resetForm({ operandA: '', operandB: '' })
    clearHistory()
  }

  const onSubmit = async (data: CalculatorFormData) => {
    const parsedA = operandSchema.safeParse(data.operandA)
    if (!parsedA.success) return

    let operandB: number | undefined
    if (!unary && data.operandB) {
      const parsedB = operandSchema.safeParse(data.operandB)
      if (!parsedB.success) return
      operandB = parsedB.data
    }

    const response = await performCalculation(parsedA.data, operandB)
    if (response) {
      addEntry(response)
      // Carry-forward: load result into operand A, clear operand B
      setValue('operandA', String(response.result))
      setValue('operandB', '')
      // Auto-focus operand B for binary operations
      if (!unary) {
        setTimeout(() => {
          const operandBInput = document.getElementById('operandB') as HTMLInputElement | null
          operandBInput?.focus()
        }, 0)
      }
    }
  }

  return (
    <div className="mx-auto w-full max-w-md space-y-4 p-4">
      <Display result={result?.result ?? null} error={error} isLoading={isLoading} />
      <div className="grid grid-cols-2 gap-2">
        <button
          type="button"
          onClick={handleClear}
          aria-label="Clear"
          className="rounded-lg bg-gray-200 p-2 text-sm font-semibold text-gray-700 transition-colors hover:bg-gray-300"
        >
          C
        </button>
        <button
          type="button"
          onClick={handleAllClear}
          aria-label="All Clear"
          className="rounded-lg bg-red-100 p-2 text-sm font-semibold text-red-700 transition-colors hover:bg-red-200"
        >
          AC
        </button>
      </div>
      <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
        <OperandInputs register={register} errors={errors} isUnary={unary} />
        <OperationButtons selectedOperation={selectedOperation} onSelect={selectOperation} />
        <button
          type="submit"
          disabled={isLoading}
          className="w-full rounded-lg bg-green-600 p-3 text-lg font-bold text-white transition-colors hover:bg-green-700 disabled:opacity-50"
        >
          {isLoading ? 'Calculating...' : 'Calculate'}
        </button>
      </form>
      <History entries={entries} onClear={clearHistory} />
    </div>
  )
}
