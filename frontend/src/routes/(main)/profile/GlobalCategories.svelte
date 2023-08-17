<script lang="ts">
	import GrumbleService from '$lib/services/GrumbleService';
	import type { _Category } from '../grumbles/grumbles';

	let categories: _Category[];

	$: getCategories();
	async function getCategories() {
		try {
			categories = await GrumbleService.categories('global');
		} catch (err) {
			console.log(err);
		}
	}
</script>

<div class="flex gap-2 mt-4">
	{#if categories}
		{#each categories as category}
			<p class="px-2 py-1 rounded-lg bg-zinc-100 border border-gray-300">{category.name}</p>
		{/each}
		<p
			class="px-2 py-1 rounded-lg bg-zinc-100 border border-gray-300 hover:bg-gray-200 cursor-pointer"
		>
			+
		</p>
	{/if}
</div>
