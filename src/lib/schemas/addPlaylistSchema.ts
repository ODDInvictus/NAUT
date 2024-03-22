import { z } from 'zod'

export const addPlaylistSchema = z
  .object({
    spotify: z.string().min(2),
  })
