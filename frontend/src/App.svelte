<script>
  import { Router, links, Route } from 'svelte-routing'
  import Home from './routes/Home.svelte'
  import Upload from './routes/Upload.svelte'
  import About from './routes/About.svelte'
  import TOS from './routes/TOS.svelte'
	import PrivacyPolicy from './routes/PrivacyPolicy.svelte'

  export let url = ''
  const year = new Date().getFullYear()
</script>

<style>
  :global(:root) {
    --background-color: #222629;
    --secondary-background-color: #333;
    --primary-color: #61892F;
    /* --secondary-color: #222629; */
    --accent-color: #86C232;
    --font-color: #6B6E70;
    --font-secondary-color: #61892F;
  }

  :global(html) {
		font-size: 62.5%;
	}

	:global(body) {
		min-height: 100vh;
		width: 100%;
		display: flex;
		flex-direction: column;
		align-items: center;
		width: 100%;
		margin: 0;
		padding: 0;
		font-size: 1.6rem;
		background: var(--background-color);
		color: var(--font-color);
	}

	header {
		height: 70px;
		width: 100%;
		border-bottom: 1px var(--accent-color) solid;
		color: var(--font-secondary-color);
		display: flex;
		align-items: center;
		justify-content:space-between;
	}

	header > a:first-child {
		margin-left: 15%;
		color: var(--font-secondary-color);
	}

	header > a:first-child:hover {
		text-decoration: none;
	}

	header > nav {
		margin-right: 15%;
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	header > nav > a {
		padding: 5px 10px;
		color: var(--accent-color);
	}

	main {
		width: 70%;
		padding: 50px 0;
		display: flex;
		flex-direction: column;
		flex: 1;
	}

	footer {
		width: 100%;
    padding: 1.1rem; /* maintains the proper height */
		background: var(--secondary-background-color);
		color: var(--font-secondary-color);
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		font-size: 1.3rem;
	}

	footer p {
		margin: 0.2rem;
		padding: 0;
	}

	footer a {
		color: var(--accent-color);
	}

  footer section {
		font-size: 1rem;
	}
</style>

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
    <Route path="tos" component="{TOS}" />
    <Route path="privacypolicy" component="{PrivacyPolicy}" />
    <Route path="/about" component={About} />
    <Route path="/upload" component={Upload} />
    <Route path="/" component={Home} />
  </main>

  <footer>
    <p>All Rights Reserved. { year } &copy; made with &hearts; by <a href="https://joshfjohnston.com">Josh Johnston</a>.</p>

    <section use:links>
			<a href="/tos">Terms of Service</a> | <a href="/privacypolicy">Privacy Policy</a>
		</section>
  </footer>
</Router>
