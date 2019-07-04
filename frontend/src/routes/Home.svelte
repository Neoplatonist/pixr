<script>
  import { onMount, onDestroy } from 'svelte'
  import Gallery from '../components/Gallery.svelte'
  import { store } from '../store'

  let error = '',
      images = [],
      lock = false,
      selectedOrder = 'oldest',
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
      const resp = await fetch(`/images?id=${id}&order=${selectedOrder}`)
      if (!resp.ok) {
        throw Error(`Could not retrieve images. Try refreshing the page.`)
      }

      let result = await resp.json()
      if (result.length == 0) {
        // nothing to see
        return
      }

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

  const handleSelect = e => {
    store.update(data => ({
      ...data,
      images: []
    }))

    getImages(0)
  }

  onDestroy(() => unsubscribe())
</script>

<style>
  #error {
    color: var(--error-color);
  }

  select {
    background-color: var(--background-color);
    color: var(--font-secondary-color);
    border: none;
    float: right;
  }

  select:hover {
    cursor: pointer;
  }

  select:focus {
    outline: none;
  }
</style>

<section>
  {#if error != ''}
    <p id="error">{ error }</p>
  {/if}

  {#if images.length > 0}
    <label for="orderSelector"></label>
    <select
      id="orderSelector"
      name="order"
      bind:value={selectedOrder}
      on:change={handleSelect}>

      <option value="oldest">Oldest</option>
      <option value="newest">Newest</option>
    </select>

    <Gallery images={images} preview={false} />
  {:else}
    <h3>No images found</h3>
  {/if}
</section>
