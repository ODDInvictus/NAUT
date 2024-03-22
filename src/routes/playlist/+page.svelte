<script lang="ts">
  import Title from '$lib/components/Title.svelte'
  import { superForm } from 'sveltekit-superforms/client'
  import PhInfo from '~icons/ph/info'
  import PhPlus from '~icons/ph/plus'
  import PhSpinner from '~icons/ph/spinner-gap'
  import type { PageData } from './$types'

  export let data: PageData
  let playlists = data.playlists.filter(playlist => playlist.id !== data.standardPlaylist)

  const { form, errors, constraints, submitting, enhance } = superForm(data.form)

  let dialog: HTMLDialogElement

  function openDialog() {
    dialog?.showModal()
  }

  function closeDialog() {
    dialog?.close()
  }
</script>

<div class="container">
  <Title title="Afspeellijsten" />

  <dialog bind:this={dialog}>
    <article>
      <header>
        <button aria-label="Close" on:click={closeDialog} rel="prev"></button>
        <p>
          <strong>Voeg afspeellijst toe</strong>
        </p>
      </header>
      <p>Plak hieronder een spotify link om een afspeellijst toe te voegen.</p>
      <form method="POST" use:enhance>
        <input
          type="text"
          name="spotify"
          aria-invalid={$errors.spotify ? 'true' : undefined}
          bind:value={$form.spotify}
          {...$constraints.spotify}
          placeholder="Spotify link" />
        {#if $errors.spotify}
          <span class="errors">{$errors.spotify}</span>
        {/if}
        <div class="container-fluid">
          {#if $submitting}
            <button disabled type="submit">
              <span class="spinner">
                <PhSpinner />
              </span>
            </button>
          {:else}
            <button type="submit">Voeg toe</button>
          {/if}
        </div>
      </form>
    </article>
  </dialog>

  <div class="add-button">
    <button on:click={() => openDialog()}> <PhPlus /> </button>
  </div>

  <table>
    <thead>
      <tr>
        <th scope="col">Naam</th>
        <th scope="col">Details</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>Standaard</td>
        <td class="info"><a href={`/playlist/${data.standardPlaylist}`}><PhInfo /></a></td>
      </tr>
      {#each playlists as playlist}
        <tr>
          <td>{playlist.name}</td>
          <td class="info"><a href={`/playlist/${playlist.id}`}><PhInfo /></a></td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>

<style lang="sass">
  .container
    display: flex
    width: 100vw
    flex-direction: column
    justify-content: center

  .info
    a
      display: flex
      justify-content: center

  .spinner :global(svg)
    animation: spin 1.5s linear infinite

  @keyframes spin
    from
      transform: rotate(0deg)
    to
      transform: rotate(360deg)

  @media (max-width: 800px)
    table
      max-width: 85vw

  .add-button
    position: absolute
    right: 1.5rem
    bottom: 1.5rem
    z-index: 100

    button
      background-color: var(--pico-color-purple-650)
      border: 2px solid var(--pico-color-purple-650)
      width: 3.5rem
      height: 3.5rem
      padding: 0
      margin: 0
      border-radius: 100%
</style>
