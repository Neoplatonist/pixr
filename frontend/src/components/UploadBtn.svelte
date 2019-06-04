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
  /* your styles go here */
</style>

<button on:click={handleClick}>Upload Files</button>
