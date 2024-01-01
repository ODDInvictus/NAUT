<script lang="ts">
  import Title from '$lib/components/Title.svelte'
  import type { PageData } from './$types'

  export let data: PageData
</script>

<div class="container">
  <Title title="Over N.A.U.T" />

  <div class="inner">
    <p>Niels' Automagische Update Tool is gebeund om IBS eindelijk eens te voorzien van fatsoenlijke CI/CD</p>

    <h2>Huidige instantie</h2>

    {#if data.version}
      <div>Versie: {data.version?.number}</div>
      <div>Gebouwd op: {data.version?.buildDate.toLocaleString('nl')}</div>
    {:else}
      <div>Versie data missend</div>
    {/if}

    <h2>Huidige gebruiker</h2>
    <div>ID: {data.user.username}</div>
    <div>Email: {data.user.email}</div>
    <div>Admin-rechten? {data.user.isAdmin}</div>
    <div>Developer-rechten? {data.user.isDev}</div>
    <div>Groepen: {data.user.groups?.join(', ')}</div>

    <h2>Links</h2>

    <a href={data.version?.link}>GitHub</a>
    <br />
    <a href={data.version?.releaseLink}>Huidige release</a>
  </div>
</div>

<style lang="sass">
  .inner
    height: 100%
    width: 100%
</style>
