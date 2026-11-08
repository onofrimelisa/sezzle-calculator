interface DisplayProps {
  result: number | null
  error: string | null
  isLoading: boolean
}

export function Display({ result, error, isLoading }: DisplayProps) {
  return (
    <div>
      <div className="w-full rounded-lg bg-gray-900 p-4 text-right font-mono text-2xl text-white sm:text-3xl">
        {isLoading && <span className="animate-pulse text-gray-400">...</span>}
        {!isLoading && <span>{result ?? 0}</span>}
      </div>
      {error && (
        <p className="mt-2 rounded-md bg-red-50 px-3 py-2 text-sm text-red-600">
          {error}
        </p>
      )}
    </div>
  )
}
