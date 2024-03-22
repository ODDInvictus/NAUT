<script lang="ts">
  import PhCheckCircleDuotone from '~icons/ph/check-circle-duotone'
  import PhXCircleDuotone from '~icons/ph/x-circle-duotone'
  export let value: boolean
</script>

<div class="container">
  {#if value}
    <span><PhCheckCircleDuotone color={'#00895A'} /></span>
  {:else}
    <span><PhXCircleDuotone color={'#D93526'} /></span>
  {/if}
</div>

<style lang="sass">
  .container
    display: flex
    align-items: center
    justify-content: center
</style>
