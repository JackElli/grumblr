<script>
	import ActionButton from '$lib/components/ActionButton.svelte';
	import PageTitle from '$lib/components/PageTitle.svelte';

	export let data;
</script>

<div class="flex items-center justify-between">
	<PageTitle>Profile</PageTitle>
	<ActionButton>Profile settings</ActionButton>
</div>
<div class="mt-4 flex gap-4 items-center">
	<div class="w-20 h-20 rounded-md border border-gray-500" />
	<p class="text-2xl">{data.user.username}</p>
</div>
<p class="mt-2">(You will be able to upload a profile pic I promise)</p>
