import { SpotifyApi } from '@spotify/web-api-ts-sdk';
import { getSetting } from './settings';

let spotify: SpotifyApi

export async function getSpotify() {
  if (!spotify) {
    await initSpotify()
  }

  return spotify
}

async function initSpotify() {
  const clientId = await getSetting('SPOTIFY_CLIENT_ID')
  const clientSecret = await getSetting('SPOTIFY_CLIENT_SECRET')

  spotify = SpotifyApi.withClientCredentials(clientId, clientSecret)
}