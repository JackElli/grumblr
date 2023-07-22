<script lang="ts">
	import PageTitle from '$lib/components/PageTitle.svelte';
	import Grumble from '$lib/components/Grumble.svelte';
	import ActionButton from '$lib/components/ActionButton.svelte';
	import StartMessage from '$lib/components/StartMessage.svelte';
	import { userStore } from '$lib/stores/userStore';
	import Categories from '$lib/components/Categories.svelte';
	import Loading from '$lib/components/Loading.svelte';
	import NetworkError from '$lib/components/NetworkError.svelte';
	import GrumbleService from '$lib/services/GrumbleService';
	import NewGlobalGrumbleModal from '../NewGlobalGrumbleModal.svelte';

	export let data;

	let newGrumbleModalVisible = false;
	let loading = false;

	$: grumbles = data.grumbles;
	$: categories = data.categories;
	$: error = data.error;
	$: welcome = $userStore?.welcome ?? true;

	async function newGrumble(e: CustomEvent) {
		loading = true;
		const grumbleText = e.detail.grumbleText;
		const category = e.detail.category;
		if (grumbleText == '') {
			return;
		}
		try {
			const grumble = await GrumbleService.new(grumbleText, category, 'global');
			newGrumbleModalVisible = false;

			grumbles?.splice(0, 0, grumble);
			grumbles = grumbles;
		} catch (e) {
			error = (e as Error).message;
		}
		loading = false;
	}
</script>

<div class="flex items-center justify-between">
	<div>
		<PageTitle>Global grumbles</PageTitle>
		{#if categories}
			<Categories type="global" {categories} class="mt-4" />
		{/if}
	</div>
	<ActionButton on:click={() => (newGrumbleModalVisible = true)}>New Grumble</ActionButton>
	<NewGlobalGrumbleModal {categories} bind:newGrumbleModalVisible on:newGrumble={newGrumble} />
</div>
<Loading loading={grumbles == undefined && error == undefined}>
	{#if error}
		<NetworkError {error} />
	{/if}
	{#if grumbles}
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
	{/if}
</Loading>
