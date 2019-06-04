<script>
  import { Router, links, Route } from 'svelte-routing'
  import Home from './routes/Home.svelte'
  import Upload from './routes/Upload.svelte'
	import About from './routes/About.svelte'

  export let url = ""
  const year = new Date().getFullYear()
</script>

<style></style>

<Router url='{url}'>
  <header use:links>
    <a href="/"><h1>Pixr</h1></a>

		<nav>
			<a href="/">Home</a>
			<a href="/upload">Upload</a>
			<a href="/about">About</a>
		</nav>
  </header>

  <main>
    <Route path="/about" component={About} />
    <Route path="/upload" component={Upload} />
    <Route path="/" component={Home} />
  </main>

  <footer>
    <p>All Rights Reserved. { year } &copy; made with &hearts; by <a href="https://joshfjohnston.com">Josh Johnston</a>.</p>
  </footer>
</Router>
