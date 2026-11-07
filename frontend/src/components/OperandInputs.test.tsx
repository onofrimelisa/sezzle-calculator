import { describe, it, expect } from 'vitest'
import { render, screen } from '@testing-library/react'
import { useForm } from 'react-hook-form'
import { OperandInputs } from './OperandInputs'
import type { CalculatorFormData } from '../schemas/calculator'

function Wrapper({ isUnary }: { isUnary: boolean }) {
  const { register, formState: { errors } } = useForm<CalculatorFormData>()
  return <OperandInputs register={register} errors={errors} isUnary={isUnary} />
}

describe('OperandInputs', () => {
  it('renders both inputs', () => {
    render(<Wrapper isUnary={false} />)
    expect(screen.getByLabelText('Operand A')).toBeInTheDocument()
    expect(screen.getByLabelText('Operand B')).toBeInTheDocument()
  })
  it('disables operand B when unary', () => {
    render(<Wrapper isUnary={true} />)
    expect(screen.getByLabelText('Operand B')).toBeDisabled()
  })
  it('enables operand B when binary', () => {
    render(<Wrapper isUnary={false} />)
    expect(screen.getByLabelText('Operand B')).not.toBeDisabled()
  })
})
