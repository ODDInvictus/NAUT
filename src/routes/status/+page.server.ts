import { addServiceSchema } from '$lib/schemas/addService'
import { db } from '$lib/server/db'
import { fail } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms/server'
import type { Actions, PageServerLoad } from './$types'

export const load = (async () => {
  const form = await superValidate(addServiceSchema)
  return { form }
}) satisfies PageServerLoad

export const actions: Actions = {
  default: async ({ request }) => {
    const form = await superValidate(request, addServiceSchema)

    console.log('POST', form)

    if (!form.valid) return fail(400, { form })

    // Now save the service
    await db.service.create({
      data: {
        ...form.data,
        lastSeen: new Date()
      }
    })

    return
  }
}
