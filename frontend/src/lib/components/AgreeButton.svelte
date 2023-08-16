<script lang="ts">
	import { userStore } from '$lib/stores/userStore';
	import { createEventDispatcher } from 'svelte';

	export let agrees: Record<string, boolean>;
	export let textSize: string = 'text-xs';

	const dispatch = createEventDispatcher();
</script>

<button
	class="inline {textSize} text-gray-500 {$$props.class} {$userStore?.id in agrees
		? 'text-green-800 font-bold'
		: 'hover:text-green-700'} cursor-pointer"
	on:click={() => dispatch('agree')}>{Object.keys(agrees).length} agrees</button
>
