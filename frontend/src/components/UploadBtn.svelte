<script>
  import { onMount, onDestroy } from 'svelte'
  import { store } from '../store'

  let images,
      unsubscribe,
      upload

  onMount(() => unsubscribe = store.subscribe(data => {
    images = data.imagesToUpload
    upload = data.upload
  }))

  const handleClick = e => {
    e.preventDefault()
    e.stopPropagation()

    store.update(data => ({
      ...data,
      upload: true
    }))

    uploadFiles()
  }

  const uploadFiles = async() => {
    console.info('TODO: upload files')

    store.update(data => ({
      ...data,
      upload: false
    }))
  }


  onDestroy(() => unsubscribe())
</script>

<style>
  button {
    padding: 10px;
    background: var(--accent-color);
    cursor: pointer;
    border-radius: 5px;
    border: 1px solid var(--font-color);
  }

  button:hover {
    background: var(--font-color);
    border: 1px solid var(--accent-color);
  }
</style>

<button on:click={handleClick}>Upload Files</button>
