<script lang="ts">
	import { userStore } from '$lib/stores/userStore';
	import { createEventDispatcher } from 'svelte';

	export let disagrees: Record<string, boolean>;
	export let forModal = false;

	const dispatch = createEventDispatcher();
</script>

{#if !forModal}
	<button
		class="inline text-xs text-gray-500 {$$props.class} {$userStore?.id in disagrees
			? 'text-red-800 font-bold'
			: 'hover:text-red-700 cursor-pointer'}"
		on:click={() => dispatch('disagree')}>{Object.keys(disagrees).length} disagrees</button
	>
{:else}
	<button
		class="inline rounded-lg px-3 bg-gray-200 border border-gray-300 text-sm {$$props.class} {$userStore?.id in
		disagrees
			? 'text-red-800 border-red-700 font-bold'
			: 'hover:text-red-700 text-gray-800 border-gray-300 cursor-pointer'}"
		on:click={() => dispatch('disagree')}>{Object.keys(disagrees).length} disagrees</button
	>
{/if}
