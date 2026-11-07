import { describe, it, expect, beforeAll, afterAll, afterEach } from 'vitest'
import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { http, HttpResponse } from 'msw'
import { setupServer } from 'msw/node'
import { Calculator } from './Calculator'

const server = setupServer(
  http.post('/api/calculate', async ({ request }) => {
    const body = (await request.json()) as Record<string, unknown>
    if (body.operation === 'add') {
      return HttpResponse.json({
        operation: 'add', operand_a: body.operand_a, operand_b: body.operand_b,
        result: (body.operand_a as number) + (body.operand_b as number),
        timestamp: '2026-11-07T15:04:05Z',
      })
    }
    if (body.operation === 'divide' && body.operand_b === 0) {
      return HttpResponse.json(
        { error: { code: 'DIVISION_BY_ZERO', message: 'Cannot divide by zero' } },
        { status: 422 }
      )
    }
    return HttpResponse.json({ operation: body.operation, operand_a: body.operand_a, result: 0, timestamp: '' })
  })
)

beforeAll(() => server.listen())
afterEach(() => server.resetHandlers())
afterAll(() => server.close())

describe('Calculator', () => {
  it('renders all sections', () => {
    render(<Calculator />)
    expect(screen.getByLabelText('Operand A')).toBeInTheDocument()
    expect(screen.getByLabelText('Operand B')).toBeInTheDocument()
    expect(screen.getByText('Calculate')).toBeInTheDocument()
  })

  it('performs a calculation and shows result', async () => {
    render(<Calculator />)
    const user = userEvent.setup()
    await user.type(screen.getByLabelText('Operand A'), '5')
    await user.type(screen.getByLabelText('Operand B'), '3')
    await user.click(screen.getByText('Calculate'))
    await waitFor(() => { expect(screen.getByText('8')).toBeInTheDocument() })
  })

  it('shows error for division by zero', async () => {
    render(<Calculator />)
    const user = userEvent.setup()
    await user.click(screen.getByLabelText('Divide'))
    await user.type(screen.getByLabelText('Operand A'), '10')
    await user.type(screen.getByLabelText('Operand B'), '0')
    await user.click(screen.getByText('Calculate'))
    await waitFor(() => { expect(screen.getByText('Cannot divide by zero')).toBeInTheDocument() })
  })

  it('disables operand B for unary operations', async () => {
    render(<Calculator />)
    const user = userEvent.setup()
    await user.click(screen.getByLabelText('Square Root'))
    expect(screen.getByLabelText('Operand B')).toBeDisabled()
  })

  it('C button clears inputs and result but preserves operation and history', async () => {
    render(<Calculator />)
    const user = userEvent.setup()
    await user.click(screen.getByLabelText('Divide'))
    await user.type(screen.getByLabelText('Operand A'), '5')
    await user.type(screen.getByLabelText('Operand B'), '3')
    await user.click(screen.getByText('Calculate'))
    await waitFor(() => { expect(screen.getByText('History')).toBeInTheDocument() })

    await user.click(screen.getByText('C'))

    expect(screen.getByLabelText('Operand A')).toHaveValue('')
    expect(screen.getByLabelText('Operand B')).toHaveValue('')
    expect(screen.getByText('0')).toBeInTheDocument()
    // History should be preserved
    expect(screen.getByText('History')).toBeInTheDocument()
    // Operation should be preserved (Divide, not reset to Add)
    expect(screen.getByLabelText('Divide')).toHaveAttribute('aria-pressed', 'true')
  })

  it('AC button clears everything including history and resets operation to add', async () => {
    render(<Calculator />)
    const user = userEvent.setup()
    await user.click(screen.getByLabelText('Divide'))
    await user.type(screen.getByLabelText('Operand A'), '10')
    await user.type(screen.getByLabelText('Operand B'), '2')
    await user.click(screen.getByText('Calculate'))
    await waitFor(() => { expect(screen.getByText('History')).toBeInTheDocument() })

    await user.click(screen.getByText('AC'))

    expect(screen.getByLabelText('Operand A')).toHaveValue('')
    expect(screen.getByLabelText('Operand B')).toHaveValue('')
    expect(screen.getByText('0')).toBeInTheDocument()
    // Operation should reset to add
    expect(screen.getByLabelText('Add')).toHaveAttribute('aria-pressed', 'true')
    // History should be cleared (component returns null when empty)
    expect(screen.queryByText('History')).not.toBeInTheDocument()
  })

  it('carries result forward as operand A after successful calculation', async () => {
    render(<Calculator />)
    const user = userEvent.setup()
    await user.type(screen.getByLabelText('Operand A'), '5')
    await user.type(screen.getByLabelText('Operand B'), '3')
    await user.click(screen.getByText('Calculate'))
    await waitFor(() => { expect(screen.getByText('8')).toBeInTheDocument() })

    // After calculation, operand A should have the result, operand B should be empty
    expect(screen.getByLabelText('Operand A')).toHaveValue('8')
    expect(screen.getByLabelText('Operand B')).toHaveValue('')
  })

  it('auto-focuses operand B after successful binary calculation', async () => {
    render(<Calculator />)
    const user = userEvent.setup()
    await user.type(screen.getByLabelText('Operand A'), '5')
    await user.type(screen.getByLabelText('Operand B'), '3')
    await user.click(screen.getByText('Calculate'))
    await waitFor(() => { expect(screen.getByText('8')).toBeInTheDocument() })
    await waitFor(() => { expect(screen.getByLabelText('Operand B')).toHaveFocus() })
  })

  it('does not carry forward on error', async () => {
    render(<Calculator />)
    const user = userEvent.setup()
    await user.click(screen.getByLabelText('Divide'))
    await user.type(screen.getByLabelText('Operand A'), '10')
    await user.type(screen.getByLabelText('Operand B'), '0')
    await user.click(screen.getByText('Calculate'))
    await waitFor(() => { expect(screen.getByText('Cannot divide by zero')).toBeInTheDocument() })

    // Inputs should stay as they were
    expect(screen.getByLabelText('Operand A')).toHaveValue('10')
    expect(screen.getByLabelText('Operand B')).toHaveValue('0')
  })
})
