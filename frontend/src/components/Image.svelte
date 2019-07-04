<script>
  import { onMount } from 'svelte'
  import { store } from '../store'

  export let image, preview

  let img
  onMount(() => {
    if (preview) {
      let reader = new FileReader()
      reader.readAsDataURL(image)
      reader.onloadend = () => {
        img = reader.result
      }
    }
  })

  const handleModal = e => {
    store.update(data => ({
      ...data,
      modal: {
        isActive: true,
        image: image
      }
    }))
  }
</script>

<style>
  img {
    width: 15rem;
    margin: 1rem;
    vertical-align: middle;
  }
</style>

{#if preview && img != undefined}
  <img
    id="preview"
    src={img}
    alt={image.name} />
{:else if !preview}
  <img
    on:click|preventDefault|stopPropagation={handleModal}
    src={image.file_location}
    alt={image.name} />
{/if}
