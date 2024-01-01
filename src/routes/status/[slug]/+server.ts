import { pingService } from '$lib/server/status'
import type { RequestHandler } from './$types'

export const POST: RequestHandler = async ({ params }) => {
  return new Response(JSON.stringify(await pingService(params.slug)), {
    status: 200
  })
}
