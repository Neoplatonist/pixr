<script>
  import Image from './Image.svelte'

  export let images, preview
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
