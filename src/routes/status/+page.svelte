<script lang="ts">
  import Title from '$lib/components/Title.svelte'
  import Icon from '@iconify/svelte'
  import type { PageData } from './$types'
  import Modal from './_modal.svelte'

  export let data: PageData

  let modalOpen = false
  const closeModal = () => (modalOpen = false)
</script>

<div class={modalOpen ? 'modal-is-opening' : 'modal-is-closing'}>
  <Modal {modalOpen} close={closeModal} {data} />

  <div class="container">
    <Title title="Services" />

    <table>
      <thead>
        <tr>
          <th>Naam</th>
          <th>Status</th>
          <th>Laatst gezien</th>
          <th>Meer informatie</th>
        </tr>
      </thead>
      <tbody>
        {#each data.applications as app}
          <tr class={app.status ? 'online' : 'offline'}>
            <td>{app.name}</td>
            <td class="status">{app.status ? 'Online' : 'Offline'}</td>
            <td>{app.lastSeen.toLocaleString('nl')}</td>
            <td><a href="/status/{app.slug}"><Icon icon="tabler:info-circle" /></a></td>
          </tr>
        {/each}
      </tbody>
    </table>

    {#if data.user.isAdmin}
      <button on:click={() => (modalOpen = true)}> Voeg service toe </button>
    {/if}
  </div>
</div>

<style lang="sass">
  tbody > .offline > .status
    color: #D93526

  tbody > .online > .status
    color: #00895A

</style>
