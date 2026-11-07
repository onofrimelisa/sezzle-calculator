import type { UseFormRegister, FieldErrors } from 'react-hook-form'
import type { CalculatorFormData } from '../schemas/calculator'

interface OperandInputsProps {
  register: UseFormRegister<CalculatorFormData>
  errors: FieldErrors<CalculatorFormData>
  isUnary: boolean
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
          disabled={isUnary}
          className={`w-full rounded-lg border p-3 text-lg ${isUnary ? 'cursor-not-allowed bg-gray-100 text-gray-400' : ''} ${errors.operandB ? 'border-red-500' : 'border-gray-300'}`}
          placeholder={isUnary ? 'N/A' : 'Enter number'}
        />
        {errors.operandB && <p className="mt-1 text-sm text-red-500">{errors.operandB.message}</p>}
      </div>
    </div>
  )
}
