<script lang="ts">
	import PageTitle from '$lib/components/PageTitle.svelte';
	import Grumble from '$lib/components/Grumble.svelte';
	import ActionButton from '$lib/components/ActionButton.svelte';
	import NewGrumbleModal from '../NewGrumbleModal.svelte';
	import StartMessage from '$lib/components/StartMessage.svelte';
	import { userStore } from '$lib/stores/userStore';
	import Categories from '$lib/components/Categories.svelte';
	import Loading from '$lib/components/Loading.svelte';
	import NetworkError from '$lib/components/NetworkError.svelte';
	import GrumbleService from '$lib/services/GrumbleService';
	import WelcomeBackMessage from '$lib/components/WelcomeBackMessage.svelte';
	import NoFriends from '../NoFriends.svelte';

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
			const grumble = await GrumbleService.new(grumbleText, category, 'friends');
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
	<title>Grumblr | {data.currentCategory}</title>
</svelte:head>
<div class="w-full">
	<div class="flex items-center justify-between">
		<div>
			<PageTitle>Friends grumbles</PageTitle>
			{#if categories}
				<Categories type="friends" {categories} class="mt-4" />
			{/if}
		</div>
		{#if data.friends != 0}
			<ActionButton on:click={() => (newGrumbleModalVisible = true)}>New Grumble</ActionButton>
			<NewGrumbleModal
				{loading}
				{categories}
				bind:newGrumbleModalVisible
				on:newGrumble={newGrumble}
			/>
		{/if}
	</div>
	<Loading loading={grumbles == undefined && error == undefined && data.friends == undefined}>
		{#if error}
			<NetworkError {error} />
		{/if}

		{#if welcome}
			<StartMessage class="mt-3" />
		{/if}

		{#if data.friends == 0}
			<NoFriends />
		{:else if grumbles}
			<div class="mt-4">
				{#if data.currentCategory == 'recents'}
					{#if !welcome}
						<WelcomeBackMessage />
					{/if}
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
</div>
