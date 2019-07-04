<script>
  import { onMount, onDestroy } from 'svelte'
  import { store } from '../store'

  export let image, preview

  let img,
      images,
      unsubscribe

  onMount(() => {
    unsubscribe = store.subscribe(data => {
      images = data.imagesToUpload
    })

    if (preview) {
      let reader = new FileReader()
      reader.readAsDataURL(image)
      reader.onloadend = () => {
        img = reader.result
      }
    }
  })

  const handleDelete = e => {
    const filtered = images.filter(data => data != image)
    store.update(data => ({
      ...data,
      imagesToUpload: filtered
    }))
  }

  const handleModal = e => {
    store.update(data => ({
      ...data,
      modal: {
        isActive: true,
        image: image
      }
    }))
  }

  onDestroy(() => unsubscribe())
</script>

<style>
  img {
    width: 15rem;
    margin: 1rem;
    vertical-align: middle;
  }

  img:hover {
    cursor: pointer;
  }

  #preview:hover {
    outline: 5px red solid;
  }
</style>

{#if preview && img != undefined}
  <img
    id="preview"
    on:click|preventDefault|stopPropagation={handleDelete}
    src={img}
    alt={image.name} />
{:else if !preview}
  <img
    on:click|preventDefault|stopPropagation={handleModal}
    src={image.file_location}
    alt={image.name} />
{/if}
