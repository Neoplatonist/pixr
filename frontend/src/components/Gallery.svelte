<script>
  import Image from './Image.svelte'

  export let images, preview

  console.log('Gallery says hello!')
</script>

<style>
  section {
    display: block;
    padding: 2rem;
  }
</style>

<section>
  {#each images as image}
    <Image image={image} preview={preview} />
  {/each}
</section>
