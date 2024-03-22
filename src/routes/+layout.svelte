<script lang="ts">
  import { goto } from '$app/navigation'
  import { page } from '$app/stores'
  import Toast from '$lib/components/Toast.svelte'
  import { ToastContainer, toasts } from 'svelte-toasts'
  import { getFlash } from 'sveltekit-flash-message'
  import '../app.sass'
  import type { LayoutData } from './$types'

  export let data: LayoutData
  const flash = getFlash(page)

  if ($flash) {
    toasts.add({
      title: $flash?.type === 'success' ? 'Gelukt' : 'Oei',
      type: $flash?.type,
      description: $flash?.message
    })
  }

  async function home() {
    await goto('/')
  }
</script>

<div class="root">
  <ToastContainer placement="top-right" let:data>
    <Toast {data} />
  </ToastContainer>

  <!-- Desktop menu -->
  <nav class="navigation desktop">
    <section class="title">
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <hgroup tabindex="0" role="link" on:click={home}>
        <h1>Invictus Radio</h1>
        <small>Herrie in de ether</small>
      </hgroup>
    </section>

    <section class="items">
      <details>
        <summary>Afspeellijsten</summary>
        <div><a href="/playlist">Overzicht</a></div>
        {#each data.playlists as playlist}
          <div><a href={`/playlist/${playlist.id}`}>{playlist.name}</a></div>
        {/each}
      </details>

      <section>
        <a href="/history">Geschiedenis</a>
      </section>

      <section>
        <a href="/queue">Wachtrij</a>
      </section>

      <section>
        <a href="/settings">Instellingen</a>
      </section>
      <section>
        <a href="/player">Speler</a>
      </section>
    </section>

    <section class="info">
      <a href="/about">Versie {data.version ?? 'onbekend'}</a>
    </section>
  </nav>

  <!-- Mobile menu -->

  <nav class="navigation mobile">
    <ul>
      <section class="title">
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <hgroup tabindex="0" role="link" on:click={home}>
          <h1>Invictus Radio</h1>
        </hgroup>
      </section>
    </ul>
    <ul class="button">
      <li>
        <details class="dropdown">
          <summary> Menu </summary>
          <ul dir="rtl">
            <li><a href="/playlist">Afspeelijsten</a></li>
            <li><a href="/history">Geschiedenis</a></li>
            <li><a href="/queue">Wachtrij</a></li>
            <li><a href="/player">Speler</a></li>
            <li><a href="/settings">Instellingen</a></li>
            <li><a href="/about">Over</a></li>
          </ul>
        </details>
      </li>
    </ul>
  </nav>

  <main>
    <slot />
  </main>
</div>

<style lang="sass">
  $nav-width: 250px
  $nav-height: 60px

  .root
    height: 100vh
    width: 100vw
    max-width: 100vw
    display: grid
    grid-template-columns: calc($nav-width + 1rem) auto
    gap: 1rem

  nav
    height: 100%
    margin-left: 1rem
    width: $nav-width
    display: flex
    flex-direction: column
    justify-content: flex-start
    
    & a
      text-decoration: none

    .title
      margin-top: 1rem
      margin-bottom: 0rem

    .info
      margin-top: auto
      margin-bottom: 1rem

      a
        text-decoration: none

    h1
      margin: 0
      font-size: 1.5rem

    small
      color: var(--pico-primary)

  .mobile
    display: none
  
  @media (max-width: 800px)
    .desktop
      display: none
    .mobile
      display: grid

    .root 
      grid-template-columns: auto
      grid-template-rows: $nav-height auto

    nav
      width: 100vw
      display: grid
      grid-template-columns: 1fr auto
      
      .title
        margin-left: 1rem

      h1
        color: var(--pico-primary) !important


      .button
        margin-top: 0.5rem
        margin-right: 1.5rem

    


</style>
