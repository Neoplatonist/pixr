<script>
  import { onMount, onDestroy } from 'svelte'
  import { store } from '../store'

  let images = [],
      index = 0,
      modal = {
        isActive: true,
        image: ''
      },
      mymodal,
      unsubscribe

  onMount(() => {
    unsubscribe = store.subscribe(data => {
      images = data.images
      modal = data.modal
    })

    index = images.indexOf(modal.image)
  })

  const handleClose = e => {
    store.update(data => ({
      ...data,
      modal: {
        isActive: false,
        image: ''
      }
    }))
  }

  const handlePrevImage = e => {
    if (index > 0) {
      index--
      const prev = images[index]
      store.update(data => ({
        ...data,
        modal: {
          isActive: true,
          image: prev
        }
      }))
    }
  }

  const handleNextImage = e => {
    if (index < images.length - 1) {
      index++
      const next = images[index]
      store.update(data => ({
        ...data,
        modal: {
          isActive: true,
          image: next
        }
      }))
    }
  }

  const handleOutsideClose = e => {
    let target = e.target

    do {
      if (target == mymodal) {
        return
      }

      target = target.parentNode
    } while(target)

    handleClose()
  }

  onDestroy(() => unsubscribe())
</script>

<style>
  article {
    background: var(--secondary-background-color);
    position: absolute;
    z-index: 10;
    top: 50%;
    left: 50%;
    -ms-transform:translate(-50%,-50%);
    -moz-transform:translate(-50%,-50%);
    -webkit-transform:translate(-50%,-50%);
    transform:translate(-50%, -50%);
    min-width: 300px;
    border-radius: 1rem;
    border: 2px var(--primary-color) solid;
  }

  article > section {
    display: flex;
    align-items: center;
    justify-content: center;
  }

  article > section > img {
    margin: 2rem 0 2rem 0;
    max-width: 50rem;
    max-height: 80vh;
  }

  article > section > button {
    margin: 2rem;
    padding: 5rem 1rem;
  }

  article > .modalBtn {
    display: block;
    text-align: center;
    margin: 0 auto 2rem auto;
    padding: 1rem 2rem;
  }

  .modalBtn {
    background: var(--primary-color);
    border-color: var(--background-color);
    border-radius: 1rem;
    font-weight: bold;
    color: var(--background-color);
    text-transform: uppercase;
  }

  .modalBtn:hover {
    background: var(--accent-color);
    cursor: pointer;
  }
</style>

<svelte:body on:click={handleOutsideClose} />

<article bind:this={mymodal}>
  <section>
    <button
      class="modalBtn"
      on:click|preventDefault|stopPropagation={handlePrevImage}>&lt;</button>

    <img src={modal.image.file_location} alt={modal.image.name}>

    <button
      class="modalBtn"
      on:click|preventDefault|stopPropagation={handleNextImage}>&gt;</button>
  </section>

  <button
    class="modalBtn"
    on:click|preventDefault|stopPropagation={handleClose}>Close</button>
</article>
