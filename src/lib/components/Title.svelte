<script lang="ts">
  export let title: string

  export let showTitle: boolean = true
</script>

<svelte:head>
  {#if showTitle}
    <title>{title} - Invictus Radio</title>
  {:else}
    <title>Invictus Radio</title>
  {/if}
</svelte:head>

<div class="container">
  <h1>{title}</h1>

  <hr />
</div>

<style lang="sass">
  .container
    margin-top: 1rem

  hr
    margin-right: 1rem
</style>
