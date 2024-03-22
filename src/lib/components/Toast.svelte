<!-- https://github.com/mzohaibqc/svelte-toasts/blob/main/src/FlatToast.svelte -->
<script lang="ts">
  import { onMount } from 'svelte'
  import { linear } from 'svelte/easing'
  import { tweened } from 'svelte/motion'

  export let data: import('svelte-toasts/types/common').ToastProps = {}

  const progress = tweened(1, {
    duration: data.duration,
    easing: linear
  })

  onMount(() => {
    progress.set(0, { duration: data.duration })
  })

  const onRemove = e => {
    e.stopPropagation()
    data?.remove()
    if (typeof data.onRemove === 'function') data.onRemove()
  }
  const onClick = () => {
    if (typeof data.onClick === 'function') data.onClick()
  }
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
<div
  class="st-toast flat {data.theme} {data.type || 'info'}"
  on:click={onClick}
  role="alert"
  aria-live="assertive"
  aria-atomic="true">
  <div class="st-toast-details">
    {#if data.title}
      <h3 class="st-toast-title">{data.title}</h3>
    {/if}

    <p class="st-toast-description">{data.description}</p>
    <div class="st-toast-extra">
      <slot name="extra" />
    </div>
  </div>
  <button class="st-toast-close-btn" type="button" aria-label="close" on:click={onRemove}>
    <slot name="close-icon">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="bx--toast-notification__close-icon"
        width="20"
        height="20"
        viewBox="0 0 32 32"
        aria-hidden="true">
        <path d="M24 9.4L22.6 8 16 14.6 9.4 8 8 9.4 14.6 16 8 22.6 9.4 24 16 17.4 22.6 24 24 22.6 17.4 16 24 9.4z" />
      </svg>
    </slot>
  </button>
  {#if data.showProgress}
    <progress style="height: {data.duration > 0 ? '4px' : 0}" value={$progress} />
  {/if}
</div>

<style lang="sass">
  .st-toast
    display: flex
    pointer-events: auto
    width: 320px
    height: auto
    padding-left: var(--pico-spacing)
    color: var(--pico-h1-color)
    background-color: var(--pico-primary-background)
    box-shadow: 0 2px 6px 0 rgba(0, 0, 0, 0.2)
    position: relative
    cursor: pointer
    border-radius: var(--pico-border-radius)


  .st-toast-details
    margin-top: var(--pico-spacing)
    margin-right: var(--pico-spacing)
    text-align: left
    align-self: flex-start

  .st-toast-details .st-toast-title
    font-size: 1.5rem
    font-weight: 600
    line-height: var(--pico-line-height)
    letter-spacing: 0.16px
    font-weight: 600
    word-break: break-word
    margin: 0
    outline: none
  
  .st-toast-close-btn
    outline: 2px solid transparent
    outline-offset: -2px
    display: flex
    align-items: center
    justify-content: center
    background-color: transparent
    border: none
    cursor: pointer
    margin-left: auto
    padding: 0
    height: 3rem
    width: 3rem
    min-height: 3rem
    min-width: 3rem
    // transition:
    //   outline 110ms
    //   background-color 110ms
</style>
