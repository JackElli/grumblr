<script lang="ts">
	import PageTitle from '$lib/components/PageTitle.svelte';
	import Grumble from '$lib/components/Grumble.svelte';
	import ActionButton from '$lib/components/ActionButton.svelte';
	import Categories from '$lib/components/Categories.svelte';
	import Loading from '$lib/components/Loading.svelte';
	import NetworkError from '$lib/components/NetworkError.svelte';
	import GrumbleService from '$lib/services/GrumbleService';
	import NewGlobalGrumbleModal from '../NewGlobalGrumbleModal.svelte';
	import NoGrumblesFound from '$lib/components/NoGrumblesFound.svelte';

	export let data;

	let newGrumbleModalVisible = false;
	let loading = false;

	$: grumbles = data.grumbles;
	$: categories = data.categories;
	$: error = data.error;

	async function newGrumble(e: CustomEvent) {
		loading = true;
		const grumbleText = e.detail.grumbleText;
		const category = e.detail.category;
		const dataType = e.detail.dataType;
		if (grumbleText == '') {
			return;
		}
		try {
			const grumble = await GrumbleService.new(grumbleText, dataType, category, 'global');
			newGrumbleModalVisible = false;

			grumbles?.splice(0, 0, grumble);
			grumbles = grumbles;
		} catch (e) {
			error = (e as Error).message;
		}
		loading = false;
	}
</script>

<svelte:head>
	<title>Grumblr Global | {data.currentCategory}</title>
</svelte:head>

<div class="w-full">
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
				{#if grumbles.length > 0}
					{#each grumbles as grumble}
						<Grumble {grumble} />
					{/each}
				{:else}
					<NoGrumblesFound on:newGrumble={() => (newGrumbleModalVisible = true)} />
				{/if}
			</div>
		{/if}
	</Loading>
</div>
