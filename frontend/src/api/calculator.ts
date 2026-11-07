import client from './client'
import type { CalculateRequest, CalculateResponse } from './types'

export async function calculate(request: CalculateRequest): Promise<CalculateResponse> {
  const response = await client.post<CalculateResponse>('/calculate', request)
  return response.data
}
