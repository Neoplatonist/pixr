<script>
  import { onMount, onDestroy } from 'svelte'
  import Gallery from '../components/Gallery.svelte'
  import { store } from '../store'

  let images = [],
      lock = false,
      unsubscribe

  onMount(() => {
    unsubscribe = store.subscribe(data => {
      images = data.images
    })

    getImages(0)
  })

  const getImages = async(id) => {
    lock = true
    try {
      const resp = await fetch(`/images?id=${id}&order=newest`)
      if (!resp.ok) {
        throw Error(`Could not retrieve images. Try refreshing the page.`)
      }

      let result = await resp.json()
      if (result.length == 0) {
        // nothing to see
        return
      }

      console.log(result)

      store.update(data => ({
        ...data,
        images: [...images, ...result]
      }))
    } catch (error) {
      store.update(data => ({
        ...data,
        error
      }))
    }
    lock = false
  }

  onDestroy(() => unsubscribe())
</script>

<style>
  /* your styles go here */
</style>

<section>
  {#if images.length > 0}
    <Gallery images={images} preview={false} />
  {:else}
    <h3>No images found</h3>
  {/if}
</section>
