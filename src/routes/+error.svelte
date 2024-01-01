<script lang="ts">
  import { page } from '$app/stores';
</script>

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

</style>