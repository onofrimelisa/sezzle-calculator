import { describe, it, expect, vi } from 'vitest'
import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { OperationButtons } from './OperationButtons'

describe('OperationButtons', () => {
  it('renders all operation buttons', () => {
    render(<OperationButtons selectedOperation="add" onSelect={() => {}} />)
    expect(screen.getByLabelText('Add')).toBeInTheDocument()
    expect(screen.getByLabelText('Subtract')).toBeInTheDocument()
    expect(screen.getByLabelText('Square Root')).toBeInTheDocument()
  })
  it('calls onSelect when button clicked', async () => {
    const onSelect = vi.fn()
    render(<OperationButtons selectedOperation="add" onSelect={onSelect} />)
    await userEvent.click(screen.getByLabelText('Subtract'))
    expect(onSelect).toHaveBeenCalledWith('subtract')
  })
  it('highlights selected operation', () => {
    render(<OperationButtons selectedOperation="add" onSelect={() => {}} />)
    const addButton = screen.getByLabelText('Add')
    expect(addButton.className).toContain('bg-blue-600')
  })
})
