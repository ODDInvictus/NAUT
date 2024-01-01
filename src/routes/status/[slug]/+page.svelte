<script lang="ts">
  import Title from '$lib/components/Title.svelte'
  import type { PageData } from './$types'

  export let data: PageData

  let canPing = true

  $: data.uptime.sort((a, b) => b.createdAt.getTime() - a.createdAt.getTime())

  async function ping() {
    canPing = false
    setTimeout(() => (canPing = true), 60 * 1000)

    const res = await fetch('', { method: 'POST' })
    const d = await res.json()
    data.uptime.push({
      ...d,
      createdAt: new Date(d.createdAt)
    })
    data = data
  }
</script>

<div class="container">
  <Title title={'Service: ' + data.app.name} />

  <div>
    Status:
    <span class={data.app.status ? 'online' : 'offline'}>
      {data.app.status ? 'Online!' : 'Offline :('}
    </span>
  </div>
  <div>Laatst gezien op: {data.app.lastSeen.toLocaleString('nl')}</div>
  <div>Service type: {data.app.type}</div>
  {#if data.app.type === 'http'}
    <div>Healthcheck URL: {data.app.url}</div>
  {:else if data.app.type === 'docker'}
    <div>Container naam: {data.app.url}</div>
  {/if}

  <br />

  <button disabled={!canPing} on:click={ping}>Ping deze service</button>
  {#if !canPing}
    <span>Je kan over 60 seconden weer pingen</span>
  {/if}

  <h2>Historische data</h2>

  <div class="uptime-graph">
    {#each data.uptime as time}
      <div class="uptime uptime-{time.status}" data-tooltip={time.createdAt.toLocaleString('nl')} />
    {/each}
  </div>
</div>

<style lang="sass">
  $online: #00895A
  $offline: #D93526
  $block-size: 32px

  .container
    width: 90%

  .uptime-graph
    display: grid
    margin-right: 1rem
    gap: 2px
    grid-template-columns: repeat( auto-fit, $block-size )

  .uptime
    border-bottom: none
    width: $block-size
    height: $block-size

  .uptime-true
    background-color: $online

  .uptime-false
    background-color: $offline

  .offline
    color: $offline

  .online
    color: $online

  button
    margin-bottom: 1rem
</style>
