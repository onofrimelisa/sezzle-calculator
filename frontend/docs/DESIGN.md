# Frontend Design Document

## Component Architecture

The frontend follows the **Container/Presentational** pattern:

- **Container components** (e.g., `Calculator`) manage state and side effects through custom hooks. They wire together presentational components and pass data/callbacks as props.
- **Presentational components** (e.g., `Display`, `OperandInputs`, `OperationButtons`, `History`) are stateless, receive data via props, and focus purely on rendering UI.

This separation keeps components small, testable, and reusable. Presentational components can be tested in isolation with simple prop fixtures, while container components are tested via integration tests.

## Custom Hooks

### useCalculator
Encapsulates all calculator-related state and API interaction:
- `selectedOperation` — the currently selected math operation
- `result` — the latest `CalculateResponse` from the API
- `error` — a user-facing error message (or null)
- `isLoading` — whether a request is in flight
- `performCalculation(operandA, operandB?)` — calls the API and updates state
- `selectOperation(op)` — changes the selected operation

By isolating API logic in a hook, components never deal with HTTP directly.

### useHistory
Manages an in-memory list of past calculations:
- `entries` — array of `CalculateResponse` objects, newest first
- `addEntry(entry)` — prepends a new entry
- `clearHistory()` — resets the list

History is kept client-side in React state. It resets on page reload.

## API Adapter Pattern

The API layer follows the **Adapter Pattern**:

1. **`client.ts`** — configures an Axios instance with base URL and error interceptors. Interceptors transform HTTP error responses into `CalculatorApiError` instances.
2. **`calculator.ts`** — exposes a `calculate()` function that accepts a `CalculateRequest` and returns a `CalculateResponse`. Components import this function; they never interact with Axios directly.
3. **`types.ts`** — defines TypeScript interfaces for requests, responses, and error types.

This pattern means swapping the HTTP library (e.g., replacing Axios with fetch) only requires changes in the adapter layer, not in components or hooks.

## Validation Strategy

Validation uses **Zod** schemas combined with **React Hook Form**:

- `operandSchema` validates that string inputs are non-empty and parseable as numbers, then transforms them to `number`.
- `calculatorSchema` defines the full form shape with `operandA` (required) and `operandB` (optional).
- The `Calculator` container manually calls `operandSchema.safeParse()` on submit to convert and validate before passing to the API hook.

React Hook Form handles form state, registration, and error display. Zod handles the actual validation rules and type coercion.

## Responsive Approach

The UI is **mobile-first** using Tailwind CSS:

- Default styles target small screens (single column layout, smaller text).
- `sm:` breakpoint adjusts layout for wider screens (two-column input grid, larger display text).
- The calculator container is constrained to `max-w-md` and centered with `mx-auto` for readability on large screens.
- Operation button grids use `grid-cols-4` (basic) and `grid-cols-3` (advanced) for compact, touch-friendly sizing.

## Testing Strategy

### Unit Tests
- **Hooks** (`useHistory`, `useCalculator`): tested with `renderHook` and `act` from `@testing-library/react`. API calls are mocked with `vi.mock()`.
- **Presentational components** (`Display`, `OperandInputs`, `OperationButtons`, `History`): tested with `render` and assertions on DOM output. User interactions tested with `@testing-library/user-event`.

### Integration Tests
- **Calculator container**: uses **MSW** (Mock Service Worker) to intercept HTTP requests at the network level. Tests exercise the full flow from user input through API call to result display, without mocking internal modules.

### Tools
- **Vitest** as the test runner (Vite-native, fast).
- **jsdom** as the DOM environment.
- **@testing-library/jest-dom** for extended DOM matchers (`toBeInTheDocument`, `toBeDisabled`, etc.).

## Error Handling

### React Errors — ErrorBoundary
A class-based `ErrorBoundary` component wraps the entire app. If any child component throws during rendering, the boundary catches it and displays a fallback UI with the error message and a "Try again" button that resets the boundary state.

### API Errors — CalculatorApiError
The Axios error interceptor in `client.ts` converts 4xx/5xx responses with a structured error body into `CalculatorApiError` instances. The `useCalculator` hook catches these and exposes the error message via its `error` state. The `Display` component renders the message in red. Non-API errors (network failures, unexpected exceptions) are caught with a generic "An unexpected error occurred" fallback.
