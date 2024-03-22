<script lang="ts">
  import Title from '$lib/components/Title.svelte'
  import type { PageData } from './$types'

  export let data: PageData
</script>

<div class="container">
  <Title title="Luisteren" />

  <div class="inner">
    <audio controls src={data.radioLink} />

    <br />
    <br />

    <small> Werkt de speler niet? Probeer dan <a href={data.radioLink}>icecast</a> </small>
    {#if data.user?.groups?.includes('colosseum')}
      <small>of <a href={data.radioLocalLink}>lokaal</a></small>
    {/if}
    <small></small>
  </div>
</div>

<style lang="sass">
  .inner
    height: 100%
    width: 100%

</style>
