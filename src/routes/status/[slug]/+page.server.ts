import { db } from '$lib/server/db';
import type { PageServerLoad } from './$types';

export const load = (async ({ params }) => {
  const slug = params.slug

  return {
    app: await db.service.findFirstOrThrow({ where: { slug } }),
    uptime: await db.uptime.findMany({ where: { serviceSlug: slug } })
  };
}) satisfies PageServerLoad;