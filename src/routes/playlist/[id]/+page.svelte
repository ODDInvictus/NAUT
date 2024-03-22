<script lang="ts">
  import { goto } from '$app/navigation'
  import Title from '$lib/components/Title.svelte'
  import { date } from '$lib/utils/date'
  import { toasts } from 'svelte-toasts'
  import PhCaretLeft from '~icons/ph/caret-left'
  import PhCaretRight from '~icons/ph/caret-right'
  import PhPlay from '~icons/ph/play'
  import PhDelete from '~icons/ph/trash'
  import type { PageData } from './$types'

  export let data: PageData

  // get page number from query string
  const url = new URL(location.href)
  let page = Number(url.searchParams.get('page')) || 0

  function nav(forward: boolean) {
    if (page == 0 && !forward) return

    const newPage = forward ? page + 1 : page - 1
    page = newPage

    const url = new URL(location.href)
    url.searchParams.set('page', newPage.toString())
    goto(url.toString(), { replaceState: true })
  }

  async function f(method: string, body: object) {
    let data = null

    try {
      const res = await fetch('', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          method,
          data: body
        })
      })

      if (!res.ok) {
        toasts.add({
          title: `Fout bij ${method}`,
          description: 'Er is iets fout gegaan bij het uitvoeren van de actie.',
          type: 'error'
        })
        return
      }

      data = await res.json()

      if (!data.success) {
        toasts.add({
          title: `Fout bij ${method}`,
          description: data.message,
          type: 'error'
        })
        return
      }
    } catch (err) {
      console.error(err)
      toasts.add({
        title: `Fout bij ${method}`,
        description: 'Er is iets fout gegaan bij het uitvoeren van de actie.',
        type: 'error'
      })
    }

    return data
  }

  async function deletePlaylist() {
    await f('delete', { id: data.playlist?.id })
  }

  async function playPlaylist() {
    await f('play', { id: data.playlist?.id })
  }

  async function queueSong(songId: string) {
    await f('queue', { songId })
  }
</script>

<div class="container">
  <Title title={data.playlist?.name ?? 'Onbekend'} />

  <div class="info">
    <div>Toegevoegd op: {date(data.playlist.lastPlayed)}</div>
    <div>Laast gespeeld op: {date(data.playlist.lastPlayed)}</div>
    <div>ID: {data.playlist.id}</div>
    <div>Aantal nummers: {data.amountOfSongs}</div>
  </div>

  <div class="buttons">
    <button on:click={() => playPlaylist()}><PhPlay />Speel</button>
    <button on:click={() => deletePlaylist()}><PhDelete />Verwijder</button>
  </div>

  <hr />

  <div class="playlist">
    {#each data.songs as song}
      <article class="song">
        <div class="image">
          <img src={song.cover} alt={song.title} />
        </div>
        <div class="song-info">
          <div>{song.title}</div>
          <div>{song.artists}</div>
        </div>
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <div class="icon" on:click={() => queueSong(song.id)}><PhPlay /></div>
      </article>
    {/each}

    {#if data.songs.length === 0}
      <p>Geen nummers gevonden</p>
    {/if}
  </div>
  <div class="pagination">
    <button class="outline" on:click={() => nav(false)}><PhCaretLeft /></button>
    <div>{page}</div>
    <button class="outline" on:click={() => nav(true)}><PhCaretRight /></button>
  </div>
</div>

<style lang="sass">
  .playlist
    overflow-y: scroll
    max-height: 50vh

    @media (max-width: 800px)
      max-height: 40vh

  .info
    margin-bottom: 1rem

  .buttons
    display: flex
    justify-content: center
    gap: 1rem
    margin-bottom: 1rem

    button :global(svg)
      margin-right: 0.5rem

  .pagination
    display: flex
    gap: 1rem
    justify-content: center
    align-items: center
    margin-top: 1rem

      

  article
    max-width: 90vw
    display: grid
    grid-template-columns: 64px auto 64px
    grid-template-rows: 1fr 1fr 1fr
    gap: 1rem

    padding-top: 0.70rem

    height: 90px

    .icon
      display: flex
      align-items: center
      justify-content: center

    .icon:hover
      cursor: pointer
      color: var(--pico-primary)

    .song-info
      max-width: calc(75vw - 128px)
      display: flex
      flex-direction: column
      justify-content: center
      
      div
        margin: 0
        text-overflow: ellipsis
        white-space: nowrap
        overflow: hidden

    .image
      grid-row: span 3
      height: 64px
      width: 64px

    img
      height: 64px
      width: 64px
</style>
