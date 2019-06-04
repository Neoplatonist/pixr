<script>
  import { onMount, onDestroy } from 'svelte'
  import { store } from '../store'
  import Gallery from '../components/Gallery.svelte'
  import UploadBtn from '../components/UploadBtn.svelte'

  let active,
      dropzone,
      images = [],
      inputFiles,
      selectButton,
      unsubscribe

  onMount(() => {
    unsubscribe = store.subscribe(data => {
      images = data.imagesToUpload
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

    dropzone.style.justifyContent = 'normal'
    selectButton.style.background = 'var(--font-color)'

    const files = e.dataTransfer.files
    handleFiles(files)
  }

  // Form Button Selection
  const handleInput = e => {
    e.preventDefault()
    e.stopPropagation()
    active = false

    handleFiles(e.files)
  }

  const handleFiles = files => {
    // keep previously add files
    const arrOfAll = [
      ...images,
      ...files
    ]

    // filters out duplicates
    const filtered = Array.from(new Set(arrOfAll.map(v => v.name)))
      .map(name => arrOfAll.find(v => v.name === name))

    // saves the filtered list
    store.update(data => ({
      ...data,
      imagesToUpload: filtered
    }))
  }

  onDestroy(() => unsubscribe())
</script>

<style>
  #dropzone {
    min-width: 400px;
    width: 80%;
    min-height: 52vh;
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

  form {
    margin-bottom: 10px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }

  #fileElem {
    display: none;
  }

  form section label {
    display: inline-block;
    padding: 10px;
    background: var(--accent-color);
    cursor: pointer;
    border-radius: 5px;
    border: 1px solid var(--font-color);
    color: var(--background-color);
  }

  form section label:hover {
    background: var(--primary-color);
    border: 1px solid var(--accent-color);
  }
</style>

<section
  id="dropzone"
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

      {#if images.length > 0}
        <UploadBtn />
      {/if}
    </section>
  </form>

  <!-- loader -->

  {#if images.length > 0}
     <Gallery images={images} preview={true} />
  {/if}
</section>
