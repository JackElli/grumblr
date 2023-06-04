<script lang="ts">
	import PageTitle from '$lib/components/PageTitle.svelte';
	import { onMount } from 'svelte';
	import type { _Grumble } from './grumbles';
	import Grumble from '$lib/components/Grumble.svelte';
	import ActionButton from '$lib/components/ActionButton.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import Loading from '$lib/components/Loading.svelte';

	let grumbles: _Grumble[];
	let newGrumbleModalVisible = false;
	let grumbleText: string;
	let loading = true;

	async function getGrumbles() {
		loading = true;
		const resp = await fetch('http://localhost:3200/grumbles', {
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
			type: 'friends'
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
	<PageTitle>Friends grumbles</PageTitle>
	<ActionButton on:click={() => (newGrumbleModalVisible = true)}>New grumble</ActionButton>
	<Modal
		title="New grumble to your friends"
		bind:visible={newGrumbleModalVisible}
		class="w-96 h-96"
	>
		<p>Add your grumble text, what are you angry about?</p>
		<textarea
			bind:value={grumbleText}
			class="mt-4 p-2 border border-black w-full h-40 resize-none outline-none rounded-md"
			placeholder="Prompt: This website needs some work..."
		/>
		<ActionButton colour="bg-green-700 hover:bg-green-600" class="mt-4" on:click={newGrumble}
			>Save</ActionButton
		>
	</Modal>
</div>
<Loading {loading}>
	{#if grumbles}
		<div class="mt-4">
			{#each grumbles as grumble}
				<Grumble {grumble} />
			{/each}
		</div>
	{/if}
</Loading>
