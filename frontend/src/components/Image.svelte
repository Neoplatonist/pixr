<script>
  import { onMount } from 'svelte'

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
</script>

<style>
  img {
    width: 15rem;
    margin: 1rem;
    vertical-align: middle;
  }
</style>

{#if preview}
   <img src={img} alt={image.name}>
{:else}
   <img src={image.file_location} alt={image.name}>
{/if}
