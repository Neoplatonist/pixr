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
    let formData = new FormData()
    images.forEach(image => {
      formData.append("file", image)
    });

    let err, status = ""
    try {
      const resp = await fetch('/images', {
        method: 'POST',
        body: formData
      })

      status = resp.status
    } catch(error) {
      err = error
      console.log(err)
    }

    if (status == 200 && err != "") {
      store.update(data => {
        return {
          ...data,
          imagesToUpload: [],
          upload: false
        }
      })
    } else {
      store.update(data => {
        return {
          ...data,
          error: err
        }
      })
    }
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
