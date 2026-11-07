import type { CalculateResponse } from '../api/types'
import { formatCalculation } from '../utils/operations'

interface HistoryProps {
  entries: CalculateResponse[]
  onClear: () => void
}

export function History({ entries, onClear }: HistoryProps) {
  if (entries.length === 0) return null
  return (
    <div className="rounded-lg border border-gray-200 bg-white p-4">
      <div className="mb-2 flex items-center justify-between">
        <h2 className="text-sm font-semibold uppercase tracking-wide text-gray-500">History</h2>
        <button onClick={onClear} className="text-sm text-red-500 hover:text-red-700">Clear</button>
      </div>
      <ul className="space-y-1">
        {entries.map((entry, index) => (
          <li key={index} className="font-mono text-sm text-gray-700">
            {formatCalculation(entry.operation, entry.operand_a, entry.operand_b, entry.result)}
          </li>
        ))}
      </ul>
    </div>
  )
}
