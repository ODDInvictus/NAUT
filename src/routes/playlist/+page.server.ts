import { addPlaylistSchema } from '$lib/schemas/addPlaylistSchema';
import { db } from '$lib/server/db';
import { getSetting } from '$lib/server/settings';
import { getSpotify } from '$lib/server/spotify';
import type { Playlist, Track } from '@spotify/web-api-ts-sdk';
import { fail, type Actions } from '@sveltejs/kit';
import { redirect } from 'sveltekit-flash-message/server';
import { superValidate } from 'sveltekit-superforms/server';
import type { PageServerLoad } from './$types';

export const load = (async () => {
  const form = await superValidate(addPlaylistSchema)

  return {
    form,
    playlists: await db.playlist.findMany({
      orderBy: {
        lastPlayed: 'desc'
      }
    }),
    standardPlaylist: await getSetting('STANDARD_PLAYLIST')
  };
}) satisfies PageServerLoad;

export const actions = {
  default: async (event) => {
    const form = await superValidate(event.request, addPlaylistSchema)

    if (!form.valid) {
      return fail(400, { form })
    }

    const playlistId = form.data.spotify

    if (playlistId) {
      let id = String(playlistId)
      if (id.includes('https://')) {
        id = id.split('/').pop() || ''
        if (id.includes('?')) {
          id = id.split('?')[0]
        }
      }

      // First query the playlist from the spotify API
      let playlist: Playlist
      let songs: Track[] = []
      try {
        const spotify = await getSpotify()
        playlist = await spotify.playlists.getPlaylist(id)

        // Now clal getPlaylistItems to get all songs
        // Do this in batches of 50
        let offset = 0

        while (offset < playlist.tracks.total) {
          const items = await spotify.playlists.getPlaylistItems(id, 'NL', undefined, 50, offset)
          songs.push(...items.items.map(item => item.track))
          offset += 50
        }
      } catch (err) {
        console.log(err)
        return redirect({ type: 'error', message: 'Playlist niet gevonden, is hij publiek?' }, event)
      }

      // Filter songs
      songs = songs.filter(song => {
        return !!song.id
      })

      // Now create the playlist
      try {
        // TODO in array vorm, ipv functie
        // await db.$transaction([
        // db.playlist.create...
        // db.song.createMany...
        // db.songInPlaylist.createMany...
        //])
        await db.$transaction(async tx => {
          await tx.playlist.create({
            data: {
              name: playlist.name,
              id: playlist.id,
            }
          })

          await tx.song.createMany({
            data: songs.map(item => {
              return {
                id: item.id,
                title: item.name,
                artists: item.artists.map(a => a.name).join(', '),
                album: item.album.name,
                cover: item.album.images[0]?.url,
                duration: item.duration_ms
              }
            }),
            skipDuplicates: true
          })

          await tx.songInPlaylist.createMany({
            data: songs.map(song => {
              return {
                songId: song.id,
                playlistId: id
              }
            }),
            skipDuplicates: true
          })
        })
      } catch (err) {
        console.error(err)
        return redirect('/playlist', { type: 'error', message: 'Er is iets fout gegaan' }, event)
      }
      return redirect(`/playlist/${id}`, { type: 'success', message: 'Playlist aangemaakt' }, event)
    }
    return { form }
  }
} satisfies Actions;