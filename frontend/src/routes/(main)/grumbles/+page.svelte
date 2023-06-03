<script lang="ts">
	import PageTitle from '$lib/components/PageTitle.svelte';
	import { onMount } from 'svelte';
	import type { _Grumble } from './grumbles';
	import Grumble from '$lib/components/Grumble.svelte';

	let grumbles: _Grumble[];

	onMount(async () => {
		const resp = await fetch('http://localhost:3200/grumbles', {
			method: 'GET',
			credentials: 'include'
		});
		grumbles = await resp.json();
	});
</script>

<PageTitle>Friends grumbles</PageTitle>
{#if grumbles}
	<div class="mt-4">
		{#each grumbles as grumble}
			<Grumble {grumble} />
		{/each}
	</div>
{/if}
