<script lang="ts">
	import LoadingSpinner from './LoadingSpinner.svelte';

	export let loading = false;
	export let disabled = false;
</script>

<button
	on:click
	disabled={loading || disabled}
	class="flex gap-2 items-center {disabled
		? 'opacity-50 cursor-not-allowed'
		: ''} justify-center min-w-[120px] px-3 py-1 shadow-[0_2px_2px_1px_rgba(0,0,0,0.2)] rounded-md hover:shadow-[0_2px_4px_2px_rgba(0,0,0,0.3)] bg-[#574658] text-white {$$props.class} "
>
	{#if loading}
		<LoadingSpinner class="w-4 text-white fill-current" />
	{/if}
	<slot />
</button>
