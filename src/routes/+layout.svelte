<script lang="ts">
  import { goto } from '$app/navigation'
  import type { LayoutData } from './$types'

  export let data: LayoutData

  async function home() {
    await goto('/')
  }
</script>

<div class="root">
  <nav class="navigation">
    <section class="title">
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <hgroup tabindex="0" role="link" on:click={home}>
        <h1>N.A.U.T</h1>
        <small>De DevOps™ ervaring</small>
      </hgroup>
    </section>

    <section class="items">
      <details>
        <summary>Builds</summary>
        <div><a href="/">Huts</a></div>
        <div><a href="/">Huts</a></div>
        <div><a href="/">Huts</a></div>
      </details>
      <details>
        <summary>Database</summary>
        <div><a href="/database/backup">Backups</a></div>
        <div><a href="/database/development">Development</a></div>
      </details>
      <details>
        <summary>Jobs</summary>
        <div><a href="/jobs/">Overzicht</a></div>
        <div><a href="/jobs/history">Geschiedenis</a></div>
      </details>
      <details>
        <summary>Status</summary>
        <div><a href="/status">Services</a></div>
        {#each data.applications as app}
          <div><a href="/status/{app.slug}">{app.name}</a></div>
        {/each}
      </details>
    </section>

    <section class="info">
      <a href="/about">Versie {data.version?.number ?? 'onbekend'}</a>
    </section>
  </nav>

  <main>
    <slot />
  </main>
</div>

<style lang="sass">
  $nav-width: 250px

  .root
    height: 100vh
    width: 100vw
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

    small
      color: var(--pico-primary)
</style>
