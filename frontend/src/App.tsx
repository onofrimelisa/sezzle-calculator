import { ErrorBoundary } from './components/ErrorBoundary'
import { Calculator } from './components/Calculator'

function App() {
  return (
    <ErrorBoundary>
      <div className="flex min-h-screen flex-col items-center bg-gray-50 pt-8 sm:pt-16">
        <h1 className="mb-6 text-3xl font-bold text-gray-800 sm:text-4xl">Calculator</h1>
        <Calculator />
      </div>
    </ErrorBoundary>
  )
}

export default App
