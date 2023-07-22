<script lang="ts">
	import LoadingSpinner from './LoadingSpinner.svelte';

	export let loading = false;
</script>

<button
	on:click
	disabled={loading}
	class="flex gap-2 items-center {loading
		? 'opacity-50 cursor-not-allowed'
		: ''} justify-center min-w-[120px] px-3 py-1 shadow-[0_2px_2px_1px_rgba(0,0,0,0.2)] rounded-md hover:shadow-[0_2px_4px_2px_rgba(0,0,0,0.3)] bg-[#dfd7d8] {$$props.class} "
>
	{#if loading}
		<LoadingSpinner class="w-4" />
	{/if}
	<slot />
</button>
