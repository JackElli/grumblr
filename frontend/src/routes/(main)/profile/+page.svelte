<script>
	import ActionButton from '$lib/components/ActionButton.svelte';
	import PageTitle from '$lib/components/PageTitle.svelte';
	import { userStore } from '$lib/stores/userStore';
	import { userIconText } from '../../../global';
	import GlobalCategories from './GlobalCategories.svelte';
</script>

<svelte:head>
	<title>Grumblr | Profile</title>
</svelte:head>

<div class="w-full">
	<div class="flex items-center justify-between">
		<PageTitle>Profile</PageTitle>
		<ActionButton disabled>Profile settings</ActionButton>
	</div>
	<div class="mt-4 flex flex-col gap-y-10">
		<div class="flex gap-4 items-center pb-3 border-b border-b-gray-200">
			<p
				class="flex text-4xl items-center justify-center w-24 h-24 border border-black rounded-md bg-gray-200 hover:bg-gray-300 cursor-pointer"
			>
				{userIconText($userStore.username)}
			</p>
			<div>
				<p class="text-2xl">{$userStore.username}</p>
				<div class="flex gap-2">
					<p class="font-semibold">
						Friends: <span class="font-normal">{$userStore.friends.length}</span>
					</p>
					<p class="font-semibold">
						Total engagement: <span class="font-normal">{$userStore.friends.length}</span>
					</p>
				</div>
			</div>
		</div>

		<div class="pb-3 border-b border-b-gray-200">
			<h1 class="text-xl font-semibold">Global categories</h1>
			<p class="text-gray-700 text-xs">These are the categories we'll show you on the global tab</p>

			<GlobalCategories />
		</div>

		<div class="pb-3 border-b border-b-gray-200">
			<h1 class="text-xl font-semibold">Your grumbles</h1>
			<p class="text-gray-700 text-xs">All of your grumbles you've ever posted</p>
			<p class="mt-4">(Under construction) 🔨</p>
		</div>
	</div>
</div>
