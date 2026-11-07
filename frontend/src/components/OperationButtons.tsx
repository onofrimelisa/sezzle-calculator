import { BASIC_OPERATIONS, ADVANCED_OPERATIONS } from '../utils/operations'

interface OperationButtonsProps {
  selectedOperation: string
  onSelect: (operation: string) => void
}

export function OperationButtons({ selectedOperation, onSelect }: OperationButtonsProps) {
  return (
    <div className="space-y-3">
      <div>
        <p className="mb-1 text-xs font-semibold uppercase tracking-wide text-gray-500">Basic</p>
        <div className="grid grid-cols-4 gap-2">
          {BASIC_OPERATIONS.map((op) => (
            <button
              key={op.name}
              onClick={() => onSelect(op.name)}
              className={`rounded-lg p-3 text-lg font-bold transition-colors ${
                selectedOperation === op.name
                  ? 'bg-blue-600 text-white'
                  : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
              }`}
              aria-label={op.label}
              aria-pressed={selectedOperation === op.name}
            >
              {op.symbol}
            </button>
          ))}
        </div>
      </div>
      <div>
        <p className="mb-1 text-xs font-semibold uppercase tracking-wide text-gray-500">Advanced</p>
        <div className="grid grid-cols-3 gap-2">
          {ADVANCED_OPERATIONS.map((op) => (
            <button
              key={op.name}
              onClick={() => onSelect(op.name)}
              className={`rounded-lg p-3 text-lg font-bold transition-colors ${
                selectedOperation === op.name
                  ? 'bg-purple-600 text-white'
                  : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
              }`}
              aria-label={op.label}
              aria-pressed={selectedOperation === op.name}
            >
              {op.symbol}
            </button>
          ))}
        </div>
      </div>
    </div>
  )
}
