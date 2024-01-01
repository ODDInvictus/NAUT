import { z } from 'zod'

export const addServiceSchema = z
  .object({
    name: z.string().min(2),
    slug: z.string().min(2),
    type: z.enum(['docker', 'http']).default('http'),
    url: z.string().min(2)
  })
