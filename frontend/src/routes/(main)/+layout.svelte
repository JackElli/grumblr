<script>
	import { version } from '../../global';
	import '../../app.css';
	import Header from '../../lib/components/Header.svelte';
</script>

<div class="bg-[#e9ede8] min-h-screen pb-10">
	<Header />
	<div class="block w-[90%] mx-auto pt-4">
		<slot />
	</div>
</div>

<h1 class="fixed bottom-2 right-2">v{version}</h1>
