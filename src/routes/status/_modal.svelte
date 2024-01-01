<script lang="ts">
  import { toSlug } from '$lib/utils/slug'
  import { superForm } from 'sveltekit-superforms/client'
  import type { PageData } from './$types'

  export let modalOpen: boolean
  export let close: () => void
  export let data: PageData

  const superform = superForm(data.form, {
    onResult: ({ result }) => {
      if (result.type === 'success') {
        window.location.href = '/status'
      }
    }
  })
  const { form, enhance, errors } = superform

  // let slug = ''
  async function onModalInput(event: Event & { currentTarget: EventTarget & HTMLInputElement }) {
    $form.slug = toSlug(event.currentTarget.value)
  }
</script>

<dialog open={modalOpen}>
  <article>
    <header>
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <div role="link" tabindex="0" on:click={close} aria-label="Close" class="close" />
      <p>
        <strong>Nieuwe service toevoegen</strong>
      </p>
    </header>

    <form method="post" use:enhance>
      <fieldset>
        <label>
          Service naam
          <input
            on:input={onModalInput}
            bind:value={$form.name}
            aria-invalid={$errors.name ? 'true' : undefined}
            name="name"
            placeholder="Naam" />
          {#if $errors.name}
            <p class="errors">{$errors.name}</p>
          {/if}
        </label>

        <label>
          Slug
          <input
            bind:value={$form.slug}
            aria-invalid={$errors.name ? 'true' : undefined}
            name="slug"
            placeholder="Slug" />
        </label>
        {#if $errors.slug}
          <p class="errors">{$errors.slug}</p>
        {/if}
      </fieldset>

      <fieldset class="type">
        <legend>Service type</legend>
        <input type="radio" id="http" name="type" bind:group={$form.type} value="http" />
        <label for="http">HTTP</label>
        <input type="radio" id="docker" name="type" bind:group={$form.type} value="docker" />
        <label for="docker">Docker</label>
      </fieldset>

      <fieldset id="health">
        <label>
          Healthcheck URL / container naam
          <input name="url" placeholder="URL" aria-invalid={$errors.name ? 'true' : undefined} bind:value={$form.url} />
        </label>
        {#if $errors.url}
          <p class="errors">{$errors.url}</p>
        {/if}
      </fieldset>

      <button>Versturen</button>
    </form>
  </article>
</dialog>

<style lang="sass">
  header
    margin-bottom: 2rem

  fieldset
    margin-bottom: 0rem

  #health
    margin-top: 1rem
    margin-bottom: 1rem

  input
    margin-bottom: 0.25rem
</style>
