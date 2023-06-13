<script lang="ts">
	import PageTitle from '$lib/components/PageTitle.svelte';
	import Grumble from '$lib/components/Grumble.svelte';
	import ActionButton from '$lib/components/ActionButton.svelte';
	import Modal from '$lib/components/Modal.svelte';
	import Loading from '$lib/components/Loading.svelte';
	import type { _Grumble } from './grumbles';
	import { onMount } from 'svelte';
	import NewGrumbleModal from './NewGrumbleModal.svelte';
	import StartMessage from '$lib/components/StartMessage.svelte';

	let grumbles: _Grumble[];
	let newGrumbleModalVisible = false;
	let welcome: false;
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

	async function getUser() {
		const userId = '1f21823a-8682-4900-b627-d6bd39e1b95b';

		const resp = await fetch(`http://localhost:3200/user/${userId}`, {
			method: 'GET',
			credentials: 'include'
		});

		const user = await resp.json();
		welcome = user.welcome;
	}

	async function newGrumble(e: CustomEvent) {
		const grumbleText = e.detail.grumbleText;
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
		await getUser();
		await getGrumbles();
	});
</script>

<div class="flex items-center justify-between">
	<PageTitle>Friends grumbles</PageTitle>
	<ActionButton on:click={() => (newGrumbleModalVisible = true)}>New grumble</ActionButton>
	<NewGrumbleModal bind:newGrumbleModalVisible on:newGrumble={newGrumble} />
</div>
<Loading {loading}>
	<div class="mt-4">
		{#if welcome}
			<StartMessage />
		{/if}
		{#if grumbles.length > 0}
			{#each grumbles as grumble}
				<Grumble {grumble} />
			{/each}
		{:else}
			<h1 class="mt-2">No grumbles found here.</h1>
		{/if}
	</div>
</Loading>
