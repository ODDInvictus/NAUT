<script lang="ts">
  import Title from '$lib/components/Title.svelte'
  import type { PageData } from './$types'

  export let data: PageData
</script>

<div class="container">
  <Title title="Over Invictus Radio" />

  <div class="inner">
    <p>Invictus Radio is <e>de</e> herrie zender van Enschede!</p>

    <h2>Huidige instantie</h2>

    {#if data.version}
      <div>Versie: {data.version} ({data.gitSha})</div>
      <div>Gebouwd op: {data.buildDate?.toLocaleString()}</div>
    {:else}
      <div>Versie data missend</div>
    {/if}

    <h2>Huidige gebruiker</h2>
    <div>ID: {data.user.username}</div>
    <div>Email: {data.user.email}</div>
    <div>Admin-rechten: {data.user.isAdmin}</div>
    <div>Groepen: {data.user.groups?.join(', ')}</div>

    <h2>Links</h2>

    <a href={data.githubLink}>GitHub</a>
  </div>
</div>

<style lang="sass">
  .inner
    padding-left: 1rem
    padding-right: 1rem
    height: 100%
    width: 100%
</style>
