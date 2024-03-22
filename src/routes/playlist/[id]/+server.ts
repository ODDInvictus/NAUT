import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ({ request }) => {
  // Get the playlist id from the request body
  const { data, method } = await request.json()

  const success = false
  let responseMessage = ""

  switch (method) {
    case 'play':
      responseMessage = "Deze functie is nog niet geïmplementeerd"
      break
    case 'delete':
      responseMessage = "Deze functie is nog niet geïmplementeerd"
      break
    case 'queue':
      responseMessage = "Deze functie is nog niet geïmplementeerd"
      break
    default:
      responseMessage = "Method not found"
      break
  }

  return new Response(JSON.stringify({
    message: responseMessage,
    success,
  }))
}