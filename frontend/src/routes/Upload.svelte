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
  const handleDragEnter = e => {
    e.preventDefault()
    e.stopPropagation()
    active = true
  }

  const handleDragOver = e => {
    e.preventDefault()
    e.stopPropagation()
    active = true
  }

  const handleDrageLeave = e => {
    e.preventDefault()
    e.stopPropagation()
    active = false
  }

  const handleDrop = e => {
    e.preventDefault()
    e.stopPropagation()
    active = false
  }

  // Form Button Selection
  const handleInput = e => {
    e.preventDefault()
    e.stopPropagation()
    active = false
  }

  onDestroy(() => unsubscribe())
</script>

<style>
  section {
    min-width: 400px;
    width: 80%;
    /* min-height: 52vh; */
    margin: 50px auto;
    padding: 20px;
    border: 2px dashed var(--font-color);
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    align-self: center;
    align-items: center;
    justify-content: center;
  }

  .active {
    border-color: var(--accent-color) !important;
  }


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
