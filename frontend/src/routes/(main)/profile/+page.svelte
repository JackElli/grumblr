<script>
	import ActionButton from '$lib/components/ActionButton.svelte';
	import PageTitle from '$lib/components/PageTitle.svelte';
	import { userStore } from '$lib/stores/userStore';
	import { userIconText } from '../../../global';
</script>

<div class="flex items-center justify-between">
	<PageTitle>Profile</PageTitle>
	<ActionButton>Profile settings</ActionButton>
</div>
<div class="mt-4 flex gap-4 items-center">
	<p
		class="flex text-4xl items-center justify-center w-24 h-24 border border-black rounded-md bg-gray-300 hover:bg-gray-100 cursor-pointer"
	>
		{userIconText($userStore.username)}
	</p>
	<p class="text-2xl">{$userStore.username}</p>
</div>
<p class="mt-2">(You will be able to upload a profile pic I promise)</p>
