<script>
  import { onMount, onDestroy } from 'svelte'
  import Gallery from '../components/Gallery.svelte'
  import Modal from '../components/Modal.svelte'
  import { store } from '../store'

  let error = '',
      images = [],
      lock = false,
      modal = {
        isActive: false,
        image: ''
      },
      selectedOrder = 'oldest',
      unsubscribe

  onMount(() => {
    unsubscribe = store.subscribe(data => {
      error = data.error
      images = data.images
      modal = data.modal
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

  const handleWheel = e => {
    const nearBottom = window.innerHeight + window.scrollY >= document.body.offsetHeight - 100

    if (nearBottom && !lock) {
      getImages(images[images.length-1].id)
    }
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
#selector {
    float: right;
  }

  #selector label {
    clear: both;
    float: left;
    margin-top: 0.7rem;
  }

  select {
    background-color: var(--background-color);
    color: var(--font-secondary-color);
    border: none;
  }

  select:hover {
    cursor: pointer;
  }

  select:focus {
    outline: none;
  }
</style>

<!-- Infinite scroll window handler -->
<svelte:window on:wheel={handleWheel}/>

<section>
  {#if error != ''}
    <p id="error">{ error }</p>
  {/if}

  {#if images.length > 0}
    <div id="selector">
      <label class="label" for="orderSelector">Order: </label>
      <select
        id="orderSelector"
        name="order"
        bind:value={selectedOrder}
        on:change={handleSelect}>

        <option value="oldest">Oldest</option>
        <option value="newest">Newest</option>
      </select>
    </div>

    <Gallery images={images} preview={false} />
  {:else}
    <h3>No images found</h3>
  {/if}

  {#if modal.isActive}
    <Modal />
  {/if}
</section>
