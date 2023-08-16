<script lang="ts">
	import { userStore } from '$lib/stores/userStore';
	import { createEventDispatcher } from 'svelte';

	export let agrees: Record<string, boolean>;
	export let forModal = false;

	const dispatch = createEventDispatcher();
</script>

{#if !forModal}
	<button
		class="inline text-xs text-gray-500 {$$props.class} {$userStore?.id in agrees
			? 'text-green-800 font-bold'
			: 'hover:text-green-700'} cursor-pointer"
		on:click={() => dispatch('agree')}>{Object.keys(agrees).length} agrees</button
	>
{:else}
	<button
		class="inline rounded-lg px-3 py-1 bg-gray-200 border text-sm {$$props.class} {$userStore?.id in
		agrees
			? 'text-green-800 font-bold border-green-800'
			: 'hover:text-green-700 text-gray-800 border-gray-300 cursor-pointer'} cursor-pointer"
		on:click={() => dispatch('agree')}>{Object.keys(agrees).length} agrees</button
	>
{/if}
