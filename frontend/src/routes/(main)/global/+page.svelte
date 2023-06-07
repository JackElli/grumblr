<script lang="ts">
	import ActionButton from '$lib/components/ActionButton.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import PageTitle from '$lib/components/PageTitle.svelte';
	import { onMount } from 'svelte';
	import Grumble from '$lib/components/Grumble.svelte';
	import type { _Grumble } from '../grumbles/grumbles';
	import Loading from '$lib/components/Loading.svelte';

	let grumbles: _Grumble[];
	let newGrumbleModalVisible = false;
	let grumbleText: string;
	let loading = true;

	async function getGrumbles() {
		loading = true;
		const resp = await fetch('http://localhost:3200/global', {
			method: 'GET',
			credentials: 'include'
		});
		grumbles = await resp.json();
		loading = false;
	}

	async function newGrumble() {
		if (grumbleText == '') {
			return;
		}

		const newGrumble: _Grumble = {
			createdBy: 'user:1',
			message: grumbleText,
			dateCreated: new Date().toISOString(),
			type: 'global'
		};

		try {
			await fetch('http://localhost:3200/grumble', {
				method: 'POST',
				credentials: 'include',
				body: JSON.stringify(newGrumble)
			});
		} catch (e) {
			console.log(e);
		}

		newGrumbleModalVisible = false;

		grumbles.splice(0, 0, newGrumble);
		grumbles = grumbles;
	}

	onMount(async () => {
		await getGrumbles();
	});
</script>

<div class="flex items-center justify-between">
	<PageTitle>Global grumbles</PageTitle>
	<ActionButton on:click={() => (newGrumbleModalVisible = true)}>New grumble</ActionButton>
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
</div>

<Loading {loading}>
	{#if grumbles.length > 0}
		<div class="mt-4">
			{#each grumbles as grumble}
				<Grumble {grumble} />
			{/each}
		</div>
	{:else}
		<h1>No grumbles found here.</h1>
	{/if}
</Loading>
