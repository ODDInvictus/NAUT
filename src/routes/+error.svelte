<script lang="ts">
  import { page } from '$app/stores'
</script>

<svelte:head>
  <title>{$page.status} - Invictus Radio</title>
</svelte:head>

<main class="container">
  <h1>{$page.status}</h1>
  <p>{$page.error?.message}</p>
</main>

<style lang="sass">
  main
    height: 100%
    width: 100%
    display: flex
    flex-direction: column
    justify-content: center
    align-items: center

  @media (max-width: 800px)
    main
      height: calc(100vh - 140px)
</style>
