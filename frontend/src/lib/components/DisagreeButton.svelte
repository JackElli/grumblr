<script lang="ts">
	import { userStore } from '$lib/stores/userStore';
	import { createEventDispatcher } from 'svelte';

	export let disagrees: Record<string, boolean>;
	export let textSize: string = 'text-xs';

	const dispatch = createEventDispatcher();
</script>

<button
	class="inline {textSize} text-gray-500 {$$props.class} {$userStore?.id in disagrees
		? 'text-red-800 font-bold'
		: 'hover:text-red-700'}  cursor-pointer"
	on:click={() => dispatch('disagree')}>{Object.keys(disagrees).length} disagrees</button
>
