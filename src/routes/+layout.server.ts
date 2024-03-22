
import { db } from '$lib/server/db';
import { getSetting } from '$lib/server/settings';
import { createSession, parseHeaders } from '$lib/server/user';
import { loadFlash } from 'sveltekit-flash-message/server';
import type { LayoutServerLoad } from './$types';

export const load = loadFlash(async (event) => {
  const headers = parseHeaders(event.request.headers);

  event.locals = headers

  // get cookie
  const session = createSession(event, headers);


  return {
    user: headers,
    version: await getSetting('VERSION'),
    radioLink: await getSetting('RADIO_LINK'),
    radioLocalLink: await getSetting('RADIO_LOCAL_LINK'),
    playlists: await db.playlist.findMany({
      orderBy: {
        lastPlayed: 'desc'
      },
      // Pick the first 5
      take: 5
    })
  };
}) satisfies LayoutServerLoad;