<script lang="ts">
	import ActionButton from '$lib/components/ActionButton.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import { createEventDispatcher } from 'svelte';
	import type { _Category } from '../grumbles/grumbles';

	export let newGrumbleModalVisible = false;
	export let loading = false;
	export let categories: _Category[] | undefined;
	export let grumbleText = '';

	let grumbleTextbox: HTMLTextAreaElement;
	$: selectedCategory = categories && categories.length ? categories[0].name : '';

	const dispatch = createEventDispatcher();

	function newGrumble() {
		if (grumbleText != '') {
			dispatch('newGrumble', {
				grumbleText: grumbleText,
				category: selectedCategory
			});
		}
	}

	$: grumbleTextbox && grumbleTextbox.focus();
</script>

<Modal
	title="New global grumble"
	subtitle="Beware! The public can see this grumble"
	bind:visible={newGrumbleModalVisible}
	class="w-96 pb-5"
>
	<p>Select a category for your grumble</p>
	<select
		class="mt-1 bg-zinc-50 focus:bg-white border border-black px-2 rounded-sm cursor-pointer"
		bind:value={selectedCategory}
	>
		{#if categories}
			{#each categories as category}
				<option value={category.name}>{category.name}</option>
			{/each}
		{/if}
	</select>
	<p class="mt-3">Add your grumble text, what are you angry about?</p>
	<textarea
		bind:this={grumbleTextbox}
		bind:value={grumbleText}
		class="mt-1 p-2 bg-gray-100 border border-black w-full h-40 resize-none outline-none rounded-md focus:bg-white"
		placeholder="Prompt: This website needs some work..."
	/>
	<ActionButton {loading} class="mt-4" on:click={newGrumble}>Create</ActionButton>
</Modal>
