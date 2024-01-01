import { DEV } from '$env/static/private';
import { db } from '$lib/server/db';
import { getAllApplications } from '$lib/server/status';
import type { LayoutServerLoad } from './$types';

function parseHeaders(headers: Headers) {
  const groupsString = headers.get('x-authentik-groups')

  const groups = groupsString?.split('|')

  const isAdmin = DEV === 'true'

  return {
    username: headers.get('x-authentik-username'),
    name: headers.get('x-authentik-name'),
    email: headers.get('x-authentik-email'),
    groups,
    isDev: groups?.includes('bakkentrekkers'),
    isAdmin: isAdmin || groups?.includes('ibs-admins')
  }
}

export const load = (async (event) => {
  return {
    applications: await getAllApplications(),
    user: parseHeaders(event.request.headers),
    version: await db.version.findFirst({ where: { id: 1 } })
  };
}) satisfies LayoutServerLoad;