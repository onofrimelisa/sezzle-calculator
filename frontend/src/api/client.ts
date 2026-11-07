import axios from 'axios'
import type { ApiErrorResponse } from './types'
import { CalculatorApiError } from './types'

const client = axios.create({
  baseURL: '/api',
  headers: { 'Content-Type': 'application/json' },
  timeout: 10000,
})

client.interceptors.response.use(
  (response) => response,
  (error) => {
    if (axios.isAxiosError(error) && error.response?.data) {
      const data = error.response.data as ApiErrorResponse
      if (data.error) {
        throw new CalculatorApiError(data.error.code, data.error.message, data.error.details)
      }
    }
    throw new CalculatorApiError('NETWORK_ERROR', 'Unable to connect to the server. Please try again.')
  }
)

export default client
