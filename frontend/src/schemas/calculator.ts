import { z } from 'zod'

export const operandSchema = z
  .string()
  .min(1, 'Required')
  .refine((val) => !isNaN(Number(val)), { message: 'Must be a valid number' })
  .transform(Number)

export const calculatorSchema = z.object({
  operandA: operandSchema,
  operandB: z.string().optional(),
})

export type CalculatorFormData = z.input<typeof calculatorSchema>
