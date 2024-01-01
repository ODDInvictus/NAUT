<script lang="ts">
  import Title from '$lib/components/Title.svelte'
  import type { PageData } from './$types'

  export let data: PageData
</script>

<div class="container">
  <Title title="Goedemiddag, {data.user.name}" showTitle={false} />
</div>

<style lang="sass">
  .container
    margin-top: 1rem
</style>
