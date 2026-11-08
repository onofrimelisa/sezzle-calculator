import { describe, it, expect } from 'vitest'
import { render, screen } from '@testing-library/react'
import { Display } from './Display'

describe('Display', () => {
  it('shows 0 when no result', () => {
    render(<Display result={null} error={null} isLoading={false} />)
    expect(screen.getByText('0')).toBeInTheDocument()
  })
  it('shows result', () => {
    render(<Display result={42} error={null} isLoading={false} />)
    expect(screen.getByText('42')).toBeInTheDocument()
  })
  it('shows error below display', () => {
    render(<Display result={null} error="Cannot divide by zero" isLoading={false} />)
    expect(screen.getByText('Cannot divide by zero')).toBeInTheDocument()
    expect(screen.getByText('0')).toBeInTheDocument()
  })
  it('shows loading', () => {
    render(<Display result={null} error={null} isLoading={true} />)
    expect(screen.getByText('...')).toBeInTheDocument()
  })
})
