<script>
  import { onMount, onDestroy, afterUpdate } from 'svelte'
  import { store } from '../store'
  import Gallery from '../components/Gallery.svelte'
  import Loader from '../components/Loader.svelte'
  import UploadBtn from '../components/UploadBtn.svelte'

  let active,
  btnGrey = false,
      dropzone,
      images = [],
      inputFiles,
      selectButton,
      unsubscribe,
      upload

  onMount(() => {
    store.update(data => ({
      ...data,
      error: ''
    }))

    unsubscribe = store.subscribe(data => {
      images = data.imagesToUpload
      upload = data.upload
    })
  })

  afterUpdate(() => {
    if (images.length == 0) {
      dropzone.style.justifyContent = 'center'
      btnGrey = false
    } else {
      dropzone.style.justifyContent = 'normal'
      btnGrey = true
    }
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
  #error {
    text-align: center;
    margin: 2rem 0 0 0;
    color: var(--error-color);
  }

  #dropzone {
    min-width: 400px;
    width: 80%;
    min-height: 52vh;
    margin: 2.5rem auto;
    padding: 2rem;
    border: 2px dashed var(--font-color);
    border-radius: 1rem;
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
    margin-bottom: 1rem;
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
    padding: 1rem;
    background: var(--accent-color);
    cursor: pointer;
    border-radius: 0.5rem;
    border: 1px solid var(--font-color);
    color: var(--background-color);
  }

  form section label:hover {
    background: var(--font-color);
    border: 1px solid var(--accent-color);
  }

  .btnGrey {
    background: var(--font-color) !important;
    border: 1px solid var(--accent-color) !important;
  }

  .btnGrey:hover {
    background: var(--accent-color) !important;
    border: 1px solid var(--font-color) !important;
    color: var(--background-color) !important;
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
        class:btnGrey
        bind:this={selectButton}
        for="fileElem">Select Files</label>

      {#if images.length > 0}
        <UploadBtn />
      {/if}
    </section>
  </form>

  {#if upload}
     <Loader />
  {/if}

  {#if images.length > 0}
     <Gallery images={images} preview={true} />
  {/if}
</section>
