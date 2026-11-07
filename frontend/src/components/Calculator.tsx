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
  const { selectedOperation, selectOperation, result, error, isLoading, performCalculation } = useCalculator()
  const { entries, addEntry, clearHistory } = useHistory()

  const { register, handleSubmit, setValue, formState: { errors } } = useForm<CalculatorFormData>({
    defaultValues: { operandA: '', operandB: '' },
  })

  const unary = isUnary(selectedOperation)

  useEffect(() => {
    if (unary) { setValue('operandB', '') }
  }, [unary, setValue])

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
    if (response) { addEntry(response) }
  }

  return (
    <div className="mx-auto w-full max-w-md space-y-4 p-4">
      <Display result={result?.result ?? null} error={error} isLoading={isLoading} />
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
