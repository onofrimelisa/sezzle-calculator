import type { UseFormRegister, FieldErrors } from 'react-hook-form'
import type { CalculatorFormData } from '../schemas/calculator'
import type { KeyboardEvent } from 'react'

interface OperandInputsProps {
  register: UseFormRegister<CalculatorFormData>
  errors: FieldErrors<CalculatorFormData>
  isUnary: boolean
}

const ALLOWED_KEYS = new Set([
  'Backspace', 'Delete', 'Tab', 'Escape', 'Enter',
  'ArrowLeft', 'ArrowRight', 'ArrowUp', 'ArrowDown',
  'Home', 'End',
])

function handleNumericKeyDown(e: KeyboardEvent<HTMLInputElement>) {
  if (ALLOWED_KEYS.has(e.key)) return
  if ((e.ctrlKey || e.metaKey) && ['a', 'c', 'v', 'x', 'z'].includes(e.key)) return
  if (/^[0-9]$/.test(e.key)) return
  if (e.key === '.' && !e.currentTarget.value.includes('.')) return
  if (e.key === '-' && e.currentTarget.selectionStart === 0 && !e.currentTarget.value.includes('-')) return
  e.preventDefault()
}

export function OperandInputs({ register, errors, isUnary }: OperandInputsProps) {
  return (
    <div className="grid grid-cols-1 gap-3 sm:grid-cols-2">
      <div>
        <label htmlFor="operandA" className="mb-1 block text-sm font-medium text-gray-700">Operand A</label>
        <input
          id="operandA"
          type="text"
          inputMode="decimal"
          {...register('operandA')}
          onKeyDown={handleNumericKeyDown}
          className={`w-full rounded-lg border p-3 text-lg ${errors.operandA ? 'border-red-500' : 'border-gray-300'}`}
          placeholder="Enter number"
        />
        {errors.operandA && <p className="mt-1 text-sm text-red-500">{errors.operandA.message}</p>}
      </div>
      <div>
        <label htmlFor="operandB" className="mb-1 block text-sm font-medium text-gray-700">Operand B</label>
        <input
          id="operandB"
          type="text"
          inputMode="decimal"
          {...register('operandB')}
          onKeyDown={handleNumericKeyDown}
          disabled={isUnary}
          className={`w-full rounded-lg border p-3 text-lg ${isUnary ? 'cursor-not-allowed bg-gray-100 text-gray-400' : ''} ${errors.operandB ? 'border-red-500' : 'border-gray-300'}`}
          placeholder={isUnary ? 'N/A' : 'Enter number'}
        />
        {errors.operandB && <p className="mt-1 text-sm text-red-500">{errors.operandB.message}</p>}
      </div>
    </div>
  )
}
