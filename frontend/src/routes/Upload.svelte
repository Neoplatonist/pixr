<script>
  import { onMount, onDestroy } from 'svelte'
  import {store } from '../store'

  let active,
      dropzone,
      inputFiles,
      selectButton,
      unsubscribe

  onMount(() => {
    unsubscribe = store.subscribe(data => {

    })
  })

  // Drag Events
  const handleDragEnter = e => {}
  const handleDragOver = e => {}
  const handleDrageLeave = e => {}
  const handleDrop = e => {}

  // Form Button Selection
  const handleInput = e => {}

  onDestroy(() => unsubscribe())
</script>

<style>
  /* your styles go here */
</style>

<section
  class:active
  bind:this={dropzone}
  on:dragenter={handleDragEnter}
  on:dragover={handleDragOver}
  on:dragleave={handleDrageLeave}
  on:drop={handleDrop}
  >
  <form>
    <p>Click and drag files or use the button.</p>

    <input
      id="fileElem"
      type="file"
      accept="image/*"
      bind:this={inputFiles}
      on:change={handleInput}
      multiple>

    <section>
      <label
        bind:this={selectButton}
        for="fileElem">Select Files</label>

      <!-- upload button -->
    </section>
  </form>

  <!-- loader -->

  <!-- gallery -->
</section>
