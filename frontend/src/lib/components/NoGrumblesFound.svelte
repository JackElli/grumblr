<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	function create() {
		dispatch('newGrumble');
	}
</script>

<h1 class="mt-20 text-center text-4xl font-bold">No grumbles found here 😢</h1>
<div class="flex justify-center">
	<button on:click={create} class="mt-4 text-[#574658] underline text-xl hover:text-[#8b748e]"
		>Be the first to add one?</button
	>
</div>
