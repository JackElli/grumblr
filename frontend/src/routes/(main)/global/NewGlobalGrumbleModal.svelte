<script>
	import ActionButton from '$lib/components/ActionButton.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import { createEventDispatcher } from 'svelte';

	export let newGrumbleModalVisible = false;
	export let grumbleText = '';

	const dispatch = createEventDispatcher();

	function newGrumble() {
		dispatch('newGrumble', {
			grumbleText: grumbleText
		});
	}
</script>

<Modal
	title="New global grumble"
	subtitle="This grumble can be seen be everyone on the platform, write at your own risk."
	bind:visible={newGrumbleModalVisible}
	class="w-96 h-96"
>
	<p>Add your grumble text, what are you angry about?</p>
	<textarea
		bind:value={grumbleText}
		class="mt-4 p-2 bg-gray-100 border border-black w-full h-40 resize-none outline-none rounded-md"
		placeholder="Prompt: This website needs some work..."
	/>
	<ActionButton colour="bg-green-700" class="mt-2" on:click={newGrumble}>Save</ActionButton>
</Modal>
