<script lang="ts">
  import Title from '$lib/components/Title.svelte'
  import type { PageData } from './$types'

  export let data: PageData

  function greeting() {
    const date = new Date()
    const hours = date.getHours()
    if (hours < 6) {
      return 'Goedenacht'
    } else if (hours < 12) {
      return 'Goedemorgen'
    } else if (hours < 18) {
      return 'Goedemiddag'
    } else {
      return 'Goedenavond'
    }
  }
</script>

<div class="container">
  <div class="desktop">
    <Title title="{greeting()}, {data.user.name}" showTitle={false} />
  </div>

  Welkom bij Invictus Radio!
</div>

<style lang="sass">
  .container
    margin-top: 1rem

  @media (max-width: 800px)
    .container
      padding: 0 1rem


</style>
