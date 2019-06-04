<script>
  import { Router, links, Route } from 'svelte-routing'
  import Home from './routes/Home.svelte'

  export let url = ""
</script>

<style></style>

<Router url='{url}'>
  <header use:links>
    <a href="/">Home</a>
  </header>

  <main>
    <Route path="/" component={Home} />
  </main>

  <footer>
    <p>All Rights Reserved. 2019 &copy; made with &hearts; by <a href="https://joshfjohnston.com">Josh Johnston</a>.</p>
  </footer>
</Router>
