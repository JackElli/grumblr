<script lang="ts">
	import PageTitle from '$lib/components/PageTitle.svelte';
	import Grumble from '$lib/components/Grumble.svelte';
	import ActionButton from '$lib/components/ActionButton.svelte';
	import type { _Category, _Grumble } from '../grumbles';
	import NewGrumbleModal from '../NewGrumbleModal.svelte';
	import StartMessage from '$lib/components/StartMessage.svelte';
	import { userStore } from '$lib/stores/userStore';
	import Categories from '$lib/components/Categories.svelte';
	import Loading from '$lib/components/Loading.svelte';
	import NetworkError from '$lib/components/NetworkError.svelte';

	export let data;

	let newGrumbleModalVisible = false;

	$: grumbles = data.grumbles;
	$: categories = data.categories;
	$: error = data.error;
	$: welcome = $userStore?.welcome ?? true;

	async function newGrumble(e: CustomEvent) {
		const grumbleText = e.detail.grumbleText;
		if (grumbleText == '') {
			return;
		}

		const newGrumble: _Grumble = {
			createdBy: $userStore.id,
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

		grumbles?.splice(0, 0, newGrumble);
		grumbles = grumbles;
	}
</script>

<div class="flex items-center justify-between">
	<PageTitle>Friends grumbles</PageTitle>
	<ActionButton on:click={() => (newGrumbleModalVisible = true)}>New grumble</ActionButton>
	<NewGrumbleModal bind:newGrumbleModalVisible on:newGrumble={newGrumble} />
</div>
<Loading loading={grumbles == undefined && error == undefined}>
	{#if error}
		<NetworkError {error} />
	{/if}
	{#if grumbles && categories}
		<div class="mt-4">
			{#if welcome}
				<StartMessage />
			{/if}
			<Categories {categories} class="mt-4" />
			{#if grumbles.length > 0}
				{#each grumbles as grumble}
					<Grumble {grumble} />
				{/each}
			{:else}
				<h1 class="mt-2">No grumbles found here.</h1>
			{/if}
		</div>
	{/if}
</Loading>
