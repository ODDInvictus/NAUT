<script lang="ts">
  import type { LayoutData } from './$types';
    
  export let data: LayoutData;
</script>

<div class="root">
  <nav class="navigation">
    <section class="title">
      <hgroup>
        <h1>N.A.U.T</h1>
        <small>De DevOps™ ervaring</small>
      </hgroup>
      <hr />
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
        <div><a href="/database/backup">Backups</div>
        <div><a href="/database/development">Development</a></div>
      </details>
      <details>
        <summary>Status</summary>
        <div><a href="/status">Status pagina</a></div>
        <div><a href="/status/ibs3">ibs3</a></div>
      </details>
    </section>

    <section class="info">
      <a href={data.version.link}>Versie {data.version.number}</a>
    </section>
  </nav>
  
  <main>
    <slot />
  </main>
</div>

<style lang="sass">
  .root
    height: 100vh
    width: 100vw
    display: grid
    grid-template-columns: calc(200px + 1rem) auto
    gap: 1rem

  hr
    height: 1px
    width: 100%
    color: white

  nav
    height: 100%
    margin-left: 1rem
    width: 200px
    display: flex
    flex-direction: column
    justify-content: flex-start

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