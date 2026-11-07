import { describe, it, expect, vi } from 'vitest'
import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { History } from './History'
import type { CalculateResponse } from '../api/types'

const entries: CalculateResponse[] = [
  { operation: 'add', operand_a: 5, operand_b: 3, result: 8, timestamp: '' },
  { operation: 'sqrt', operand_a: 9, result: 3, timestamp: '' },
]

describe('History', () => {
  it('renders nothing when empty', () => {
    const { container } = render(<History entries={[]} onClear={() => {}} />)
    expect(container.firstChild).toBeNull()
  })
  it('renders entries', () => {
    render(<History entries={entries} onClear={() => {}} />)
    expect(screen.getByText('5 + 3 = 8')).toBeInTheDocument()
    expect(screen.getByText('√9 = 3')).toBeInTheDocument()
  })
  it('calls onClear when clear clicked', async () => {
    const onClear = vi.fn()
    render(<History entries={entries} onClear={onClear} />)
    await userEvent.click(screen.getByText('Clear'))
    expect(onClear).toHaveBeenCalled()
  })
})
