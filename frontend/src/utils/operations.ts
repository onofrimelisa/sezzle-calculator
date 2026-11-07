export interface OperationInfo {
  name: string
  label: string
  symbol: string
  category: 'basic' | 'advanced'
  type: 'binary' | 'unary'
}

export const OPERATIONS: OperationInfo[] = [
  { name: 'add', label: 'Add', symbol: '+', category: 'basic', type: 'binary' },
  { name: 'subtract', label: 'Subtract', symbol: '-', category: 'basic', type: 'binary' },
  { name: 'multiply', label: 'Multiply', symbol: '×', category: 'basic', type: 'binary' },
  { name: 'divide', label: 'Divide', symbol: '÷', category: 'basic', type: 'binary' },
  { name: 'power', label: 'Power', symbol: '^', category: 'advanced', type: 'binary' },
  { name: 'sqrt', label: 'Square Root', symbol: '√', category: 'advanced', type: 'unary' },
  { name: 'percentage', label: 'Percentage', symbol: '%', category: 'advanced', type: 'binary' },
]

export const BASIC_OPERATIONS = OPERATIONS.filter((op) => op.category === 'basic')
export const ADVANCED_OPERATIONS = OPERATIONS.filter((op) => op.category === 'advanced')

export function getOperation(name: string): OperationInfo | undefined {
  return OPERATIONS.find((op) => op.name === name)
}

export function isUnary(operationName: string): boolean {
  return getOperation(operationName)?.type === 'unary'
}

export function formatCalculation(operation: string, operandA: number, operandB: number | undefined, result: number): string {
  const op = getOperation(operation)
  if (!op) return `${operandA} = ${result}`
  if (op.type === 'unary') return `${op.symbol}${operandA} = ${result}`
  return `${operandA} ${op.symbol} ${operandB} = ${result}`
}
