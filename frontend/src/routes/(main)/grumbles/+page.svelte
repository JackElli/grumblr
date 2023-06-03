<script lang="ts">
	import Card from '$lib/components/Card.svelte';
	import PageTitle from '$lib/components/PageTitle.svelte';
	import { onMount } from 'svelte';
	import type { Grumble } from './grumbles';

	let grumbles: Grumble[];

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
			<Card>{grumble.message}</Card>
		{/each}
	</div>
{/if}
