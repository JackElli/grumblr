<script lang="ts">
	import ActionButton from '$lib/components/ActionButton.svelte';
	import PageTitle from '$lib/components/PageTitle.svelte';
	import { onMount } from 'svelte';
	import Grumble from '$lib/components/Grumble.svelte';
	import type { _Grumble } from '../grumbles/grumbles';
	import Loading from '$lib/components/Loading.svelte';
	import NewGlobalGrumbleModal from './NewGlobalGrumbleModal.svelte';
	import { Auth } from '$lib/services/AuthService';
	import { userStore } from '$lib/stores/userStore';

	let grumbles: _Grumble[];
	let newGrumbleModalVisible = false;
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

	async function newGrumble(e: CustomEvent) {
		const grumbleText = e.detail.grumbleText;
		if (grumbleText == '') {
			return;
		}

		const newGrumble: _Grumble = {
			createdBy: $userStore.id,
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
		await Auth();
		await getGrumbles();
	});
</script>

<div class="flex items-center justify-between">
	<PageTitle>Global grumbles</PageTitle>
	<ActionButton on:click={() => (newGrumbleModalVisible = true)}>New grumble</ActionButton>
	<NewGlobalGrumbleModal bind:newGrumbleModalVisible on:newGrumble={newGrumble} />
</div>

<Loading {loading}>
	{#if grumbles.length > 0}
		<div class="mt-4">
			{#each grumbles as grumble}
				<Grumble {grumble} />
			{/each}
		</div>
	{:else}
		<h1 class="mt-2">No grumbles found here.</h1>
	{/if}
</Loading>
