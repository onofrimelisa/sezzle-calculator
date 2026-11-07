interface DisplayProps {
  result: number | null
  error: string | null
  isLoading: boolean
}

export function Display({ result, error, isLoading }: DisplayProps) {
  return (
    <div className="w-full rounded-lg bg-gray-900 p-4 text-right font-mono text-2xl text-white sm:text-3xl">
      {isLoading && <span className="animate-pulse text-gray-400">Calculating...</span>}
      {error && <span className="text-red-400">{error}</span>}
      {result !== null && !error && !isLoading && <span>{result}</span>}
      {result === null && !error && !isLoading && <span className="text-gray-500">0</span>}
    </div>
  )
}
