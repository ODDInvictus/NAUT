import { db } from '$lib/server/db';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async ({ params, url }) => {
  const playlist = await db.playlist.findFirst({
    where: {
      id: params.id
    }
  })

  if (!playlist) {
    error(404, 'Playlist niet gevonden')
  }

  // Get the offset from the query string
  const page = url.searchParams.get('page') || '0';

  const songs = await db.song.findMany({
    where: {
      SongInPlaylist: {
        some: {
          playlistId: playlist.id
        }
      }
    },
    // select the first 20
    take: 20,
    skip: parseInt(page) * 20
  })

  const amountOfSongs = await db.song.count({
    where: {
      SongInPlaylist: {
        some: {
          playlistId: playlist.id
        }
      }
    }
  })

  return { playlist, songs, amountOfSongs };
}) satisfies PageServerLoad;