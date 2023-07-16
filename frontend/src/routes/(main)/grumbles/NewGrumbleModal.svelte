<script lang="ts">
	import ActionButton from '$lib/components/ActionButton.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import { createEventDispatcher } from 'svelte';
	import type { _Category } from './grumbles';

	export let newGrumbleModalVisible = false;
	export let categories: _Category[] | undefined;
	export let grumbleText = '';

	let grumbleTextbox: HTMLTextAreaElement;
	let selectedCategory = '';

	const dispatch = createEventDispatcher();

	function newGrumble() {
		dispatch('newGrumble', {
			grumbleText: grumbleText,
			category: selectedCategory
		});
	}

	$: grumbleTextbox && grumbleTextbox.focus();
</script>

<Modal
	title="New grumble to your friends"
	subtitle="Only your friends can see this grumble."
	bind:visible={newGrumbleModalVisible}
	class="w-96 pb-5"
>
	<select bind:value={selectedCategory}>
		{#if categories}
			{#each categories as category}
				<option value={category.name}>{category.name}</option>
			{/each}
		{/if}
	</select>
	<p>Add your grumble text, what are you angry about?</p>
	<textarea
		bind:this={grumbleTextbox}
		bind:value={grumbleText}
		class="mt-4 p-2 bg-gray-100 border border-black w-full h-40 resize-none outline-none rounded-md"
		placeholder="Prompt: This website needs some work..."
	/>
	<ActionButton class="mt-4" on:click={newGrumble}>Create</ActionButton>
</Modal>
